#!/usr/bin/env python3
"""
Parse SMBIOS structure definition tables from AI-generated markdown.

Handles:
  - Normalizing en-dash to hyphen
  - Merging consecutive <table> blocks that were split across pages
  - Filtering for tables whose headers contain at minimum:
      Offset, Name, Length, Value, Description
  - Preserving all columns present in the header (e.g. Spec. Version)
  - Capturing the 5 lines of context text above each table group
  - Building a parent-child hierarchy of main type tables and sub-tables

Usage as a module:
    from parse_smbios import SMBIOSParser
    parser = SMBIOSParser("DSP0134_3_3_0.md")
    result = parser.parse()

Usage from CLI:
    python parse_smbios.py DSP0134_3_3_0.md -o output.json
"""

import re
import sys
import json
import argparse
from pathlib import Path
from dataclasses import dataclass, field, asdict
from html.parser import HTMLParser


# ── HTML Table Parser ────────────────────────────────────────────────────────

class _TableHTMLParser(HTMLParser):
    """Extract rows from a single <table> block."""

    def __init__(self):
        super().__init__()
        self.rows: list[list[str]] = []
        self._current_row: list[str] | None = None
        self._current_cell: list[str] | None = None
        self._in_cell = False

    def handle_starttag(self, tag, attrs):
        if tag == "tr":
            self._current_row = []
        elif tag in ("td", "th"):
            self._current_cell = []
            self._in_cell = True
        elif tag == "br" and self._in_cell:
            self._current_cell.append("\n")

    def handle_endtag(self, tag):
        if tag in ("td", "th") and self._current_cell is not None:
            self._current_row.append("".join(self._current_cell).strip())
            self._current_cell = None
            self._in_cell = False
        elif tag == "tr" and self._current_row is not None:
            self.rows.append(self._current_row)
            self._current_row = None

    def handle_data(self, data):
        if self._in_cell and self._current_cell is not None:
            self._current_cell.append(data)


# ── Data Structures ──────────────────────────────────────────────────────────

@dataclass
class ParsedTable:
    """A single parsed table (main or child) with its context and position."""
    context: str
    table_name: str
    headers: list[str] = field(default_factory=list)
    rows: list[dict[str, str]] = field(default_factory=list)
    line: int = 0


# ── Parser Class ─────────────────────────────────────────────────────────────

class SMBIOSParser:
    """
    Parses SMBIOS structure definition tables from AI-generated markdown.

    Attributes:
        filepath:   Path to the source markdown file.
        stats:      Dict of parsing statistics (populated after parse).
        tables:     List of main tables with nested children (populated after parse).
        table_map:  TOC lookup mapping 'Table N' -> table title string.
    """

    REQUIRED_HEADERS = {"Offset", "Name", "Length", "Value", "Description"}

    def __init__(self, filepath: str):
        self.filepath = Path(filepath)
        self.table_map: dict[str, str] = {}
        self.stats: dict = {}
        self.tables: list[dict] = []
        self._lines: list[str] = []

    # ── Public API ───────────────────────────────────────────────────────

    def parse(self) -> dict:
        """
        Run the full parsing pipeline and return a dict with 'stats' and 'tables'.

        The returned structure:
            {
                "stats": { ... },
                "tables": [
                    {
                        "context": "...",
                        "table_name": "Table 6 BIOS Information (Type 0)",
                        "headers": [...],
                        "rows": [{...}, ...],
                        "line": 1325,
                        "children": [ { same shape minus children }, ... ]
                    },
                    ...
                ]
            }
        """
        raw = self.filepath.read_text(encoding="utf-8")
        normalized = self._normalize_dashes(raw)
        self._load_table_map(normalized)
        self._lines = normalized.splitlines()

        blocks = self._extract_table_blocks()
        groups = self._merge_consecutive_blocks(blocks)

        all_tables = []
        for group in groups:
            table = self._parse_merged_group(group)
            if table is not None:
                all_tables.append(asdict(table))

        self.tables = self._build_hierarchy(all_tables)

        main_count = len(self.tables)
        child_count = sum(len(t["children"]) for t in self.tables)
        total_rows = sum(
            len(t["rows"]) + sum(len(c["rows"]) for c in t["children"])
            for t in self.tables
        )

        self.stats = {
            "total_html_table_blocks": len(blocks),
            "merged_table_groups": len(groups),
            "total_parsed_tables": len(all_tables),
            "main_tables": main_count,
            "child_tables": child_count,
            "total_data_rows": total_rows,
        }

        return {"stats": self.stats, "tables": self.tables}

    def to_json(self, indent: int = 2) -> str:
        """Return the parsed result as a JSON string."""
        return json.dumps(
            {"stats": self.stats, "tables": self.tables},
            indent=indent,
            ensure_ascii=False,
        )

    def print_stats(self, file=sys.stderr):
        """Print parsing statistics to the given stream."""
        s = self.stats
        print(f"── Stats ──────────────────────────────", file=file)
        print(f"  HTML <table> blocks found:  {s['total_html_table_blocks']}", file=file)
        print(f"  Merged table groups:        {s['merged_table_groups']}", file=file)
        print(f"  Total parsed tables:        {s['total_parsed_tables']}", file=file)
        print(f"  Main type tables:           {s['main_tables']}", file=file)
        print(f"  Child tables:               {s['child_tables']}", file=file)
        print(f"  Total data rows:            {s['total_data_rows']}", file=file)
        print(f"───────────────────────────────────────", file=file)

    # ── Normalization ────────────────────────────────────────────────────

    @staticmethod
    def _normalize_dashes(text: str) -> str:
        """Replace en-dash with standard hyphen."""
        return text.replace("\u2013", "-")

    @staticmethod
    def _strip_html_tags(text: str) -> str:
        """Remove inline HTML tags like <i>, <b>, etc."""
        return re.sub(r"<[^>]+>", "", text)

    # ── Table-of-contents map ────────────────────────────────────────────

    def _load_table_map(self, markdown: str):
        """
        Build a lookup from 'Table N' -> table title by parsing the
        document's table-of-contents section.
        """
        self.table_map = {}
        toc_text = markdown[:60000]
        toc_match = re.search(r"Tables(.*)Foreword", toc_text, re.DOTALL)
        if toc_match is None:
            print("WARNING: Could not find table-of-contents section", file=sys.stderr)
            return

        for line in toc_match.group(1).strip().splitlines():
            m = re.search(r"(Table\s+\d+)\s*-\s*(.*?)\s*\.{2,}", line)
            if m:
                self.table_map[m.group(1)] = m.group(2)

    def _find_table_name(self, context: str) -> str:
        """
        Extract a table name from the context text.
        Tries an inline 'Table N - ... (Type N)' pattern first,
        then falls back to the TOC map.
        """
        match = re.search(r"Table \d+\s?-\s?.*Type \d+\)", context)
        if match is not None:
            return match.group(0)

        refs = re.findall(r"Table\s\d+", context)
        if refs:
            table_num = refs[-1]
            return (table_num + " " + self.table_map.get(table_num, "")).strip()

        return ""

    # ── Block extraction ─────────────────────────────────────────────────

    def _extract_table_blocks(self) -> list[dict]:
        """
        Find all <table>...</table> blocks and return them with their
        line indices and up to 5 preceding context lines.
        """
        lines = self._lines
        blocks = []
        i = 0
        while i < len(lines):
            if lines[i].strip() == "<table>":
                start = i
                j = i + 1
                while j < len(lines) and lines[j].strip() != "</table>":
                    j += 1
                end = j

                context_lines = []
                k = start - 1
                while k >= 0 and len(context_lines) < 5:
                    stripped = lines[k].strip()
                    if stripped == "</table>":
                        break
                    if (stripped
                        and not stripped.startswith("<")
                        and not stripped.startswith("</")
                    ):
                        context_lines.insert(0, stripped)
                    k -= 1

                blocks.append({
                    "start": start,
                    "end": end,
                    "html": "\n".join(lines[start:end + 1]),
                    "context": "\n".join(context_lines),
                })
                i = end + 1
            else:
                i += 1
        return blocks

    # ── Block merging ────────────────────────────────────────────────────

    def _are_consecutive(self, block_a: dict, block_b: dict) -> bool:
        """True if only whitespace separates two table blocks."""
        between = self._lines[block_a["end"] + 1 : block_b["start"]]
        return all(line.strip() == "" for line in between)

    def _merge_consecutive_blocks(self, blocks: list[dict]) -> list[list[dict]]:
        """Group consecutive table blocks into merged groups."""
        if not blocks:
            return []

        groups: list[list[dict]] = [[blocks[0]]]
        for block in blocks[1:]:
            if self._are_consecutive(groups[-1][-1], block):
                groups[-1].append(block)
            else:
                groups.append([block])
        return groups

    # ── Table parsing ────────────────────────────────────────────────────

    @staticmethod
    def _parse_table_html(html: str) -> list[list[str]]:
        """Parse an HTML table string into a list of rows."""
        parser = _TableHTMLParser()
        parser.feed(html)
        return parser.rows

    @staticmethod
    def _to_key(header: str) -> str:
        """Convert a header string to a snake_case dict key."""
        return header.strip().lower().replace(".", "").replace(" ", "_")

    def _parse_merged_group(self, group: list[dict]) -> ParsedTable:
        """
        Parse a group of consecutive table blocks into a single ParsedTable.
        Merges continuation blocks and deduplicates headers.
        """
        all_rows: list[list[str]] = []
        header_row: list[str] | None = None

        for i, block in enumerate(group):
            rows = self._parse_table_html(block["html"])
            if not rows:
                continue

            if i == 0:
                header_row = [cell.strip() for cell in rows[0]]
                all_rows.extend(rows)
            else:
                all_rows.extend(rows[1:])

        if header_row is None:
            return None

        keys = [self._to_key(h) for h in header_row]
        context = group[0]["context"]
        line = group[0]["start"]
        name = self._find_table_name(context)

        table = ParsedTable(
            table_name=name,
            context=context,
            headers=header_row,
            line=line,
        )

        for row in all_rows[1:]:
            while len(row) < len(keys):
                row.append("")
            table.rows.append({
                keys[j]: self._strip_html_tags(row[j]).strip()
                for j in range(len(keys))
            })

        return table

    # ── Classification & hierarchy ───────────────────────────────────────

    def _is_main_type_table(self, table: dict) -> bool:
        """
        A main SMBIOS type table has the required headers AND its first
        row is offset 00h with a field named 'Type'.
        """
        headers_set = {h.strip() for h in table.get("headers", [])}
        if not self.REQUIRED_HEADERS.issubset(headers_set):
            return False
        if not table.get("rows"):
            return False
        first = table["rows"][0]
        return first.get("offset", "") == "00h" and first.get("name", "") == "Type"

    def _build_hierarchy(self, all_tables: list[dict]) -> list[dict]:
        """
        Arrange tables into a parent-child hierarchy.
        Every table between two consecutive main tables becomes a child
        of the preceding main table.
        """
        main_indices = [
            i for i, t in enumerate(all_tables) if self._is_main_type_table(t)
        ]

        if not main_indices:
            return []

        result = []
        for pos, main_idx in enumerate(main_indices):
            next_main_idx = (
                main_indices[pos + 1]
                if pos + 1 < len(main_indices)
                else len(all_tables)
            )

            main_table = dict(all_tables[main_idx])
            main_table["children"] = all_tables[main_idx + 1 : next_main_idx]
            result.append(main_table)

        return result


# ── CLI ──────────────────────────────────────────────────────────────────────

def main():
    argp = argparse.ArgumentParser(
        description="Parse SMBIOS structure tables from AI-generated markdown"
    )
    argp.add_argument("input", help="Path to the markdown file")
    argp.add_argument(
        "-o", "--output",
        help="Output JSON file (default: stdout)",
        default=None,
    )
    args = argp.parse_args()

    parser = SMBIOSParser(args.input)
    parser.parse()
    parser.print_stats()

    output = parser.to_json()
    if args.output:
        Path(args.output).write_text(output, encoding="utf-8")
        print(f"Wrote to {args.output}", file=sys.stderr)
    else:
        print(output)


if __name__ == "__main__":
    main()