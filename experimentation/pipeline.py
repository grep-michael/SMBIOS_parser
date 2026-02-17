#!/usr/bin/env python3
"""
SMBIOS Spec → Go Code Pipeline

Parses AI-generated markdown files of SMBIOS specifications and generates
Go structs with shared interfaces across spec versions.

Usage:
    # Parse markdown files and generate Go code
    python smbios_pipeline.py /path/to/markdown_files/ -o ./output

    # Parse + dump intermediate JSON (for debugging)
    python smbios_pipeline.py /path/to/markdown_files/ -o ./output --json

Stages:
    1. Parse:    .md files → hierarchical table structures (SMBIOSParser)
    2. Generate: table structures → Go structs + interfaces  (CodeGenerator)
"""

import re
import sys
import os
import json
import argparse
from pathlib import Path
from dataclasses import dataclass, field, asdict
from html.parser import HTMLParser


# ╔══════════════════════════════════════════════════════════════════════════════╗
# ║  Stage 1: Markdown Parsing                                                 ║
# ╚══════════════════════════════════════════════════════════════════════════════╝

class _TableHTMLParser(HTMLParser):
    """Extract rows from a single <table> block."""

    def __init__(self):
        super().__init__()
        self.rows: list[list[str]] = []
        self._current_row: list[str] | None = None
        self._current_cell: list[str] | None = None
        self._in_cell = False
        self._skip_tags = {"sup", "sub"}
        self._skip_depth = 0

    def handle_starttag(self, tag, attrs):
        if tag in self._skip_tags:
            self._skip_depth += 1
            return
        if tag == "tr":
            self._current_row = []
        elif tag in ("td", "th"):
            self._current_cell = []
            self._in_cell = True
        elif tag == "br" and self._in_cell:
            self._current_cell.append("\n")

    def handle_endtag(self, tag):
        if tag in self._skip_tags:
            self._skip_depth -= 1
            return
        if tag in ("td", "th") and self._current_cell is not None:
            self._current_row.append("".join(self._current_cell).strip())
            self._current_cell = None
            self._in_cell = False
        elif tag == "tr" and self._current_row is not None:
            self.rows.append(self._current_row)
            self._current_row = None

    def handle_data(self, data):
        if self._skip_depth > 0:
            return
        if self._in_cell and self._current_cell is not None:
            self._current_cell.append(data)


@dataclass
class ParsedTable:
    """A single parsed table (main or child) with its context and position."""
    context: str
    table_name: str
    table_num: str
    type: str
    headers: list[str] = field(default_factory=list)
    rows: list[dict[str, str]] = field(default_factory=list)
    line: int = 0


TABLETYPE_STRUCT = "Struct"
TABLETYPE_ENUM = "Enum"
TABLETYPE_BITFIELD = "BitField"
TABLETYPE_BITRANGE = "BitRange"


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
        self.Name: str = ""
        self.Version: str = ""
        self.tables: list[dict] = []
        self._lines: list[str] = []

    def parse(self) -> dict:
        """Run the full parsing pipeline."""
        raw = self.filepath.read_text(encoding="utf-8")
        normalized = self._normalize_dashes(raw)
        self._load_table_map(normalized)
        self._parse_version(normalized)
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
            {
                "DocumentInfo": {
                    "Version": self.Version,
                    "DocumentName": self.Name,
                    "File": self.filepath.as_posix(),
                },
                "stats": self.stats,
                "tables": self.tables,
                "TableMap": self.table_map,
            },
            indent=indent,
            ensure_ascii=False,
        )

    def to_dict(self) -> dict:
        """Return the parsed result as a dict (for passing directly to CodeGenerator)."""
        return {
            "DocumentInfo": {
                "Version": self.Version,
                "DocumentName": self.Name,
                "File": self.filepath.as_posix(),
            },
            "stats": self.stats,
            "tables": self.tables,
            "TableMap": self.table_map,
        }

    def print_stats(self, file=sys.stderr):
        """Print parsing statistics."""
        s = self.stats
        print(f"── {self.filepath.name} ────────────────────────", file=file)
        print(f"  Version:                    {self.Version}", file=file)
        print(f"  HTML <table> blocks found:  {s['total_html_table_blocks']}", file=file)
        print(f"  Merged table groups:        {s['merged_table_groups']}", file=file)
        print(f"  Total parsed tables:        {s['total_parsed_tables']}", file=file)
        print(f"  Main type tables:           {s['main_tables']}", file=file)
        print(f"  Child tables:               {s['child_tables']}", file=file)
        print(f"  Total data rows:            {s['total_data_rows']}", file=file)

    # ── Version / document info ──────────────────────────────────────────

    def _parse_version(self, markdown: str):
        top_of_file = markdown[:1000]
        document_name = re.search(r"Document\s*.*:\s*(.*)", top_of_file)
        self.Name = document_name.group(1) if document_name is not None else "DocumentNameNotFound"
        version = re.search(r"Version\s*.*:\s*(.*)", top_of_file)
        if version is not None:
            self.Version = version.group(1)
        else:
            # Fallback: extract version from filename (e.g. "3.3.0.md" or "DSP0134_3_3_0.md")
            stem = self.filepath.stem
            m = re.search(r"(\d+[\._]\d+[\._]\d+)", stem)
            self.Version = m.group(1).replace("_", ".") if m is not None else "VersionNotFound"

    def _get_table_type(self, table_rows: list[str]) -> str:
        headers = [h.strip().lower() for h in table_rows]
        if "value" in headers[0] and "meaning" in headers[-1]:
            return TABLETYPE_ENUM
        if "bit position" in headers[0]:
            return TABLETYPE_BITFIELD
        if "bit range" in headers[0]:
            return TABLETYPE_BITRANGE
        if {"offset", "name", "length", "value", "description"}.issubset(headers):
            return TABLETYPE_STRUCT
        return ""

    # ── Normalization ────────────────────────────────────────────────────

    @staticmethod
    def _normalize_dashes(text: str) -> str:
        return text.replace("\u2013", "-")

    @staticmethod
    def _strip_html_tags(text: str) -> str:
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
            # The OCR sometimes renders the table list as actual HTML tables
            line = line.strip()
            line = line.replace("<tr>", "")
            line = line.replace("<td>", "")
            line = line.split("...")[0]
            line = line.split("</td>")[0]
            m = re.search(r"(Table\s+\d+)\s*-\s*(.*)", line)
            if m:
                self.table_map[m.group(1)] = m.group(2)

    def _find_table_identifiers(self, context: str) -> tuple[str, str]:
        """Extract table number and name from context text."""
        match = re.search(r"Table \d+\s?-\s?.*Type \d+\)", context)
        if match is not None:
            ids = match.group(0).split("-")
            return (ids[0].strip(), ids[1].strip())

        refs = re.findall(r"Table\s\d+", context)
        if refs:
            table_num: str = refs[-1]
            return (table_num.strip(), self.table_map.get(table_num, "").strip())

        return ("", "")

    # ── Block extraction ─────────────────────────────────────────────────

    def _extract_table_blocks(self) -> list[dict]:
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
                    if stripped and not stripped.startswith("<") and not stripped.startswith("</"):
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
        between = self._lines[block_a["end"] + 1 : block_b["start"]]
        return all(line.strip() == "" for line in between)

    def _merge_consecutive_blocks(self, blocks: list[dict]) -> list[list[dict]]:
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
        parser = _TableHTMLParser()
        parser.feed(html)
        return parser.rows

    @staticmethod
    def _to_key(header: str) -> str:
        return header.strip().lower().replace(".", "").replace(" ", "_")

    def _parse_merged_group(self, group: list[dict]) -> ParsedTable:
        """
        Parse a group of consecutive table blocks into a single ParsedTable.
        Merges continuation blocks and deduplicates headers.
        """
        all_rows: list[list[str]] = []
        header_row: list[str] = None

        for i, block in enumerate(group):
            rows = self._parse_table_html(block["html"])
            if not rows:
                continue
            if i == 0:
                header_row = [cell.strip() for cell in rows[0]]
                all_rows.extend(rows)
            else:
                if header_row is not None and rows[0] == header_row:
                    all_rows.extend(rows[1:])
                else:
                    all_rows.extend(rows)

        if header_row is None:
            return None

        keys = [self._to_key(h) for h in header_row]
        context = group[0]["context"]
        line = group[0]["start"]
        table_num, table_name = self._find_table_identifiers(context)
        ttype = self._get_table_type(header_row)

        table = ParsedTable(
            table_name=table_name,
            table_num=table_num,
            type=ttype,
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
        headers_set = {h.strip() for h in table.get("headers", [])}
        if not self.REQUIRED_HEADERS.issubset(headers_set):
            return False
        if not table.get("rows"):
            return False
        first = table["rows"][0]
        return first.get("offset", "") == "00h" and first.get("name", "") == "Type"

    def _build_hierarchy(self, all_tables: list[dict]) -> list[dict]:
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


# ╔══════════════════════════════════════════════════════════════════════════════╗
# ║  Stage 2: Go Code Generation                                               ║
# ╚══════════════════════════════════════════════════════════════════════════════╝

TYPE_MAP = {
    "BYTE": "byte",
    "WORD": "uint16",
    "DWORD": "uint32",
    "QWORD": "uint64",
}

DESIRED_TYPES = [
    "0", "1", "3", "4", "9", "16", "17",
    "22", "24", "39", "126", "127",
]

INTERFACE_NAMES = {
    "0":   "BIOS",
    "1":   "System",
    "3":   "Chassis",
    "4":   "Processor",
    "9":   "SystemSlot",
    "16":  "PhysicalMemoryArray",
    "17":  "MemoryDevice",
    "22":  "PortableBattery",
    "24":  "HardwareSecurity",
    "39":  "PowerSupply",
    "126": "Inactive",
    "127": "EndOfTable",
}


def FilterName(name: str) -> str:
    for c in " ()-,./\\":
        name = name.replace(c, "")
    return name


class StructField:
    def __init__(self, row_table: dict):
        self.row: dict = row_table
        self.name: str = FilterName(self.row["name"])
        self.comment: str = self._gen_comment()
        self.type: str = self._gen_type()
        self.is_string: bool = self.row.get("value", "") == "STRING"
        self.is_enum: bool = self.row.get("value", "") == "ENUM"
        self.is_bitfield: bool = self.row.get("value", "") == "Bit Field"

    def _gen_comment(self):
        comment = "//"
        value = self.row["value"]
        value = value if value in ["STRING", "Bit Field", "ENUM"] else ""
        comment += value
        return comment

    def _gen_type(self) -> str:
        length: str = self.row["length"]
        if length not in TYPE_MAP:
            self.comment += " Type:" + length

        type_s = TYPE_MAP.get(length, None)
        if type_s is not None:
            return type_s

        if self.name == "BIOSCharacteristicsExtensionBytes":
            return "[2]byte"
        if self.name == "ContainedElements":
            return "[3]byte"
        match = re.match(r"(\d+) bytes", length.lower())
        if match is not None:
            return f"[{match.group(1)}]byte"

        return "interface{}"

    def gen_field_string(self) -> str:
        return f"{self.name} {self.type} {self.comment}"


class GoStruct:
    def __init__(self, table: dict, smbios_ver: str):
        self.ver: str = smbios_ver.replace(".", "_")
        self.table: dict = table
        self.field_names: dict = {}
        self.StructName: str = ""
        self.StructNumber: str = ""
        self.fields: list[StructField] = []
        self._gen_name()
        self._gen_fields()

    def _gen_name(self):
        name: str = self.table["rows"][0]["description"]
        struct_num: str = self.table["rows"][0]["value"]
        self.StructName = FilterName(name)
        self.StructName = f"SMB{self.ver}_S{struct_num}_{self.StructName.lower()}"
        self.StructNumber = struct_num

    def _gen_fields(self):
        rows: list = self.table["rows"]
        self.field_names = {}
        count = 0
        for row in rows:
            field = StructField(row)
            if field.name.strip() == "":
                continue
            if field.name not in self.field_names:
                self.field_names[field.name] = "present"
            else:
                field.name = field.name + "_" + str(count)
                count += 1
            self.fields.append(field)

    def gen_struct_string(self) -> str:
        definition = f"type {self.StructName} struct {{\n"
        for row in self.fields:
            field = row.gen_field_string()
            if "interface" in field:
                print(self.StructName, field)
            definition += f"\t{field}\n"
        definition += "}"
        return definition


# ── Interface Generation ─────────────────────────────────────────────────────

class InterfaceMethod:
    """Represents a single getter method on a shared interface."""
    def __init__(self, field_name: str, go_type: str, is_string: bool, is_enum: bool):
        self.field_name = field_name
        self.go_type = go_type
        self.is_string = is_string
        self.is_enum = is_enum

    @property
    def method_name(self) -> str:
        return f"Get{self.field_name}"

    def signature(self) -> str:
        if self.is_string or self.is_enum:
            return f"{self.method_name}(strings []string) string"
        return f"{self.method_name}() {self.go_type}"

    def impl(self, struct_name: str) -> str:
        if self.is_string or self.is_enum:
            return (
                f"func (s *{struct_name}) {self.method_name}(strings []string) string {{\n"
                f"\tidx := int(s.{self.field_name})\n"
                f"\tif idx > 0 && idx <= len(strings) {{\n"
                f"\t\treturn strings[idx-1]\n"
                f"\t}}\n"
                f"\treturn \"\"\n"
                f"}}\n"
            )
        return (
            f"func (s *{struct_name}) {self.method_name}() {self.go_type} {{\n"
            f"\treturn s.{self.field_name}\n"
            f"}}\n"
        )


class InterfaceGenerator:
    """Builds an interface from fields common across all versions of a type."""
    def __init__(self, type_num: str, structs: list[GoStruct]):
        self.type_num = type_num
        self.iface_name = INTERFACE_NAMES.get(type_num, f"Type{type_num}") + "Info"
        self.wrapper_name = INTERFACE_NAMES.get(type_num, f"Type{type_num}")
        self.methods: list[InterfaceMethod] = []
        self.struct_names: list[str] = [s.StructName for s in structs]
        self._build(structs)

    def _build(self, structs: list[GoStruct]):
        if not structs:
            return

        field_maps: list[dict[str, StructField]] = []
        for s in structs:
            fm = {}
            for f in s.fields:
                fm[f.name] = f
            field_maps.append(fm)

        common_names = set(field_maps[0].keys())
        for fm in field_maps[1:]:
            common_names &= set(fm.keys())

        skip = {"Type", "Length", "Handle"}
        for f in structs[0].fields:
            if f.name in common_names and f.name not in skip:
                self.methods.append(InterfaceMethod(
                    field_name=f.name,
                    go_type=f.type,
                    is_string=f.is_string,
                    is_enum=f.is_enum,
                ))

    def gen_interface_string(self) -> str:
        lines = [f"type {self.iface_name} interface {{"]
        for m in self.methods:
            lines.append(f"\t{m.signature()}")
        lines.append("}")
        return "\n".join(lines)

    def gen_wrapper_string(self) -> str:
        return (
            f"type {self.wrapper_name} struct {{\n"
            f"\tInfo    {self.iface_name}\n"
            f"\tStrings []string\n"
            f"}}"
        )

    def gen_impl_string(self, struct_name: str) -> str:
        lines = []
        for m in self.methods:
            lines.append(m.impl(struct_name))
        return "\n".join(lines)


# ── Code Generator ───────────────────────────────────────────────────────────

class CodeGenerator:
    """
    Generates Go code from parsed SMBIOS data.

    Can accept data either as pre-parsed dicts (from SMBIOSParser.to_dict())
    or by loading JSON files from disk.
    """
    def __init__(self, dest_dir: str):
        self.dest_dir: Path = Path(dest_dir)
        self.version_structs: dict[str, dict[str, GoStruct]] = {}
        self.version_data: dict[str, dict] = {}

    def add_version(self, data: dict):
        """Add a parsed SMBIOS version from a dict (no file I/O needed)."""
        version = data["DocumentInfo"]["Version"]
        self.version_data[version] = data

        structs: dict[str, GoStruct] = {}
        for table in data["tables"]:
            gs = GoStruct(table, version)
            if gs.StructNumber in DESIRED_TYPES:
                structs[gs.StructNumber] = gs

        self.version_structs[version] = structs

    def load_json(self, filepath: Path):
        """Add a parsed SMBIOS version from a JSON file on disk."""
        with open(filepath, "r") as f:
            data = json.load(f)
        self.add_version(data)

    def generate(self):
        """Write all Go output files."""
        os.makedirs(self.dest_dir, exist_ok=True)

        for ver in self.version_data:
            self._write_struct_file(ver)

        self._write_interfaces_file()

    def _write_struct_file(self, version: str):
        ver_underscore = version.replace(".", "_")
        structs = self.version_structs[version]

        code = "package GeneratedCode\n\n"

        for type_num in DESIRED_TYPES:
            gs = structs.get(type_num)
            if gs is None:
                continue
            code += gs.gen_struct_string() + "\n"

        interfaces = self._build_interfaces()
        for type_num, iface in interfaces.items():
            gs = structs.get(type_num)
            if gs is None:
                continue
            code += "\n// -- " + iface.wrapper_name + " interface methods --\n\n"
            code += iface.gen_impl_string(gs.StructName) + "\n"

        outfile = self.dest_dir.joinpath(f"smbios_{ver_underscore}.go")
        with open(outfile, "w") as f:
            f.write(code)
        print(f"  Wrote {outfile}", file=sys.stderr)

    def _write_interfaces_file(self):
        interfaces = self._build_interfaces()

        code = "package GeneratedCode\n\n"
        for type_num in DESIRED_TYPES:
            iface = interfaces.get(type_num)
            if iface is None or not iface.methods:
                continue
            code += "// " + iface.wrapper_name + "\n"
            code += iface.gen_interface_string() + "\n\n"
            code += iface.gen_wrapper_string() + "\n\n"

        outfile = self.dest_dir.joinpath("smbios_interfaces.go")
        with open(outfile, "w") as f:
            f.write(code)
        print(f"  Wrote {outfile}", file=sys.stderr)

    def _build_interfaces(self) -> dict[str, InterfaceGenerator]:
        interfaces: dict[str, InterfaceGenerator] = {}
        for type_num in DESIRED_TYPES:
            all_structs = []
            for ver_structs in self.version_structs.values():
                gs = ver_structs.get(type_num)
                if gs is not None:
                    all_structs.append(gs)

            if all_structs:
                interfaces[type_num] = InterfaceGenerator(type_num, all_structs)

        return interfaces


# ╔══════════════════════════════════════════════════════════════════════════════╗
# ║  CLI: Ties both stages together                                             ║
# ╚══════════════════════════════════════════════════════════════════════════════╝

def collect_markdown_files(path: Path) -> list[Path]:
    """Collect .md files from a file or directory."""
    if path.is_file() and path.suffix == ".md":
        return [path]
    if path.is_dir():
        return sorted(path.glob("*.md"))
    return []


def main():
    argp = argparse.ArgumentParser(
        description="SMBIOS Spec → Go Code Pipeline: parse markdown specs and generate Go structs + interfaces"
    )
    argp.add_argument(
        "input",
        help="Path to a markdown file or directory of markdown files",
    )
    argp.add_argument(
        "-o", "--output",
        help="Output directory for generated Go files (default: ./CodeGen)",
        default="./CodeGen",
    )
    argp.add_argument(
        "--json",
        help="Also write intermediate JSON files to the output directory",
        action="store_true",
    )
    args = argp.parse_args()

    input_path = Path(args.input)
    output_dir = Path(args.output)
    md_files = collect_markdown_files(input_path)

    if not md_files:
        print(f"No .md files found at {input_path}", file=sys.stderr)
        sys.exit(1)

    print(f"Found {len(md_files)} markdown file(s)", file=sys.stderr)
    print(file=sys.stderr)

    # ── Stage 1: Parse all markdown files ────────────────────────────────
    codegen = CodeGenerator(output_dir)

    for md_file in md_files:
        parser = SMBIOSParser(str(md_file))
        parser.parse()
        parser.print_stats()

        # Feed parsed data directly into the code generator (no JSON round-trip)
        codegen.add_version(parser.to_dict())

        # Optionally dump JSON for debugging / standalone use
        if args.json:
            os.makedirs(output_dir, exist_ok=True)
            json_name = md_file.stem + ".json"
            json_path = output_dir / json_name
            json_path.write_text(parser.to_json(), encoding="utf-8")
            print(f"  Wrote {json_path}", file=sys.stderr)

        print(file=sys.stderr)

    # ── Stage 2: Generate Go code ────────────────────────────────────────
    print("Generating Go code...", file=sys.stderr)
    codegen.generate()
    print("\nDone.", file=sys.stderr)


if __name__ == "__main__":
    main()