#!/usr/bin/env python3
"""
Parse SMBIOS structure definition tables from AI-generated markdown.

Handles:
  - Normalizing en-dash (–) to hyphen (-)
  - Merging consecutive <table> blocks that were split across pages
  - Filtering for tables whose headers contain at minimum:
      Offset, Name, Length, Value, Description
  - Preserving all columns present in the header (e.g. Spec. Version)
  - Capturing the 5 lines of context text above each table group
"""

import re
import json
import argparse
from pathlib import Path
from dataclasses import dataclass, field, asdict
from html.parser import HTMLParser


class TableHTMLParser(HTMLParser):
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



@dataclass
class ParsedTable:
    """A single parsed table (main or child) with its context and position."""
    context: str                        # up to 5 lines of text above the table
    table_name: str
    headers: list[str] = field(default_factory=list)
    rows: list[dict[str, str]] = field(default_factory=list)
    line: int = 0                       # starting line in the source file


# Minimum required headers for a "main" SMBIOS type table
REQUIRED_HEADERS = {"Offset", "Name", "Length", "Value", "Description"}

def normalize_dashes(text: str) -> str:
    """Replace en-dash (–) with standard hyphen (-)."""
    return text.replace("\u2013", "-")


def strip_html_tags(text: str) -> str:
    """Remove inline HTML tags like <i>, <b>, etc. from cell text."""
    return re.sub(r"<[^>]+>", "", text)


def extract_table_blocks(lines: list[str]) -> list[dict]:
    """
    Find all <table>...</table> blocks and return them with their
    line indices and the 5 preceding context lines.

    Returns list of dicts:
        { 'start': int, 'end': int, 'html': str, 'context': str }
    """
    blocks = []
    i = 0
    while i < len(lines):
        if lines[i].strip() == "<table>":
            start = i
            # find closing </table>
            j = i + 1
            while j < len(lines) and lines[j].strip() != "</table>":
                j += 1
            end = j  # line index of </table>

            # grab 5 non-empty, non-HTML context lines above the table
            # stop at the previous </table> to avoid bleeding context from
            # an earlier section's text
            context_lines = []
            k = start - 1
            while k >= 0 and len(context_lines) < 5:
                stripped = lines[k].strip()
                # stop if we hit a prior table boundary
                if stripped == "</table>":
                    break
                # skip blank lines and residual HTML rows
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

def are_consecutive(block_a: dict, block_b: dict, lines: list[str]) -> bool:
    """
    Two table blocks are consecutive if there's only whitespace
    between the </table> of A and the <table> of B.
    """
    between = lines[block_a["end"] + 1 : block_b["start"]]
    return all(line.strip() == "" for line in between)


def merge_consecutive_blocks(blocks: list[dict], lines: list[str]) -> list[list[dict]]:
    """Group consecutive table blocks into merged groups."""
    if not blocks:
        return []

    groups: list[list[dict]] = [[blocks[0]]]
    for block in blocks[1:]:
        if are_consecutive(groups[-1][-1], block, lines):
            groups[-1].append(block)
        else:
            groups.append([block])
    return groups


def parse_table_html(html: str) -> list[list[str]]:
    """Parse an HTML table string into a list of rows (list of cell strings)."""
    parser = TableHTMLParser()
    parser.feed(html)
    return parser.rows


def has_required_headers(row: list[str]) -> bool:
    """Check if a row contains at least all REQUIRED_HEADERS."""
    return REQUIRED_HEADERS.issubset({cell.strip() for cell in row})


def parse_merged_group(group: list[dict]) -> ParsedTable:
    """
    Parse a group of consecutive table blocks into a single ParsedTable.
    Handles any table structure — merges continuation blocks and
    deduplicates headers.
    """
    all_rows: list[list[str]] = []
    header_row: list[str] | None = None

    for i, block in enumerate(group):
        rows = parse_table_html(block["html"])
        if not rows:
            continue

        if i == 0:
            header_row = [cell.strip() for cell in rows[0]]
            all_rows.extend(rows)
        else:
            # Continuation table — skip its header row
            all_rows.extend(rows[1:])

    if header_row is None:
        return None

    # Normalize header keys to snake_case for JSON output
    def to_key(h: str) -> str:
        return h.strip().lower().replace(".", "").replace(" ", "_")

    keys = [to_key(h) for h in header_row]

    context = group[0]["context"]
    line = group[0]["start"]
    match = re.search(r"Table \d+\s?-\s?.*Type \d+\)",context)
    if match == None:
        name = ""
    else:
        name = match.group(0)
    table = ParsedTable(table_name=name,context=context, headers=header_row, line=line)

    for row in all_rows[1:]:  # skip header
        while len(row) < len(keys):
            row.append("")
        entry = {
            keys[j]: strip_html_tags(row[j]).strip()
            for j in range(len(keys))
        }
        table.rows.append(entry)

    return table


def is_main_type_table(table: dict) -> bool:
    """
    A main SMBIOS type table has the required headers AND its first
    row is offset 00h with a field named 'Type'.
    """
    headers_set = {h.strip() for h in table.get("headers", [])}
    if not REQUIRED_HEADERS.issubset(headers_set):
        return False
    if not table.get("rows"):
        return False
    first = table["rows"][0]
    return first.get("offset", "") == "00h" and first.get("name", "") == "Type"


def build_hierarchy(all_tables: list[dict]) -> list[dict]:
    """
    Arrange tables into a parent-child hierarchy.

    Every table between two consecutive main tables becomes a child
    of the first main table. Tables before the first main table are
    also collected as children of it (preamble tables).
    """
    # Find indices of main tables
    main_indices = [i for i, t in enumerate(all_tables) if is_main_type_table(t)]

    if not main_indices:
        return []

    result = []
    for pos, main_idx in enumerate(main_indices):
        # Determine the range of child tables: everything after this
        # main table up to (but not including) the next main table
        next_main_idx = (
            main_indices[pos + 1] if pos + 1 < len(main_indices) else len(all_tables)
        )

        main_table = dict(all_tables[main_idx])  # shallow copy
        children = all_tables[main_idx + 1 : next_main_idx]

        main_table["children"] = children
        result.append(main_table)

    return result


def parse_smbios_markdown(filepath: str) -> dict:
    """
    Main entry point. Reads the markdown file, normalizes dashes,
    extracts and merges tables, builds a parent-child hierarchy,
    and returns a dict with stats and parsed tables.
    """
    raw = Path(filepath).read_text(encoding="utf-8")
    normalized = normalize_dashes(raw)
    lines = normalized.splitlines()

    blocks = extract_table_blocks(lines)
    groups = merge_consecutive_blocks(blocks, lines)

    # Parse every table group (not just ones with required headers)
    all_tables = []
    for group in groups:
        table = parse_merged_group(group)
        if table is not None:
            all_tables.append(asdict(table))

    hierarchy = build_hierarchy(all_tables)

    main_count = len(hierarchy)
    child_count = sum(len(t["children"]) for t in hierarchy)
    total_rows = sum(
        len(t["rows"]) + sum(len(c["rows"]) for c in t["children"])
        for t in hierarchy
    )

    return {
        "stats": {
            "total_html_table_blocks": len(blocks),
            "merged_table_groups": len(groups),
            "total_parsed_tables": len(all_tables),
            "main_tables": main_count,
            "child_tables": child_count,
            "total_data_rows": total_rows,
        },
        "tables": hierarchy,
    }


def main():
    parser = argparse.ArgumentParser(
        description="Parse SMBIOS structure tables from AI-generated markdown"
    )
    parser.add_argument("input", help="Path to the markdown file")
    parser.add_argument(
        "-o", "--output",
        help="Output JSON file (default: stdout)",
        default=None,
    )
    args = parser.parse_args()

    result = parse_smbios_markdown(args.input)
    stats = result["stats"]

    # Print stats to stderr so they're visible even when piping JSON
    import sys
    print(f"── Stats ──────────────────────────────", file=sys.stderr)
    print(f"  HTML <table> blocks found:  {stats['total_html_table_blocks']}", file=sys.stderr)
    print(f"  Merged table groups:        {stats['merged_table_groups']}", file=sys.stderr)
    print(f"  Total parsed tables:        {stats['total_parsed_tables']}", file=sys.stderr)
    print(f"  Main type tables:           {stats['main_tables']}", file=sys.stderr)
    print(f"  Child tables:               {stats['child_tables']}", file=sys.stderr)
    print(f"  Total data rows:            {stats['total_data_rows']}", file=sys.stderr)
    print(f"───────────────────────────────────────", file=sys.stderr)

    output = json.dumps(result, indent=2, ensure_ascii=False)
    if args.output:
        Path(args.output).write_text(output, encoding="utf-8")
        print(f"Wrote to {args.output}", file=sys.stderr)
    else:
        print(output)


if __name__ == "__main__":
    main()