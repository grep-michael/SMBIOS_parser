#!/usr/bin/env python3
"""
Generate Go code from a parsed SMBIOS JSON file.

Produces for each SMBIOS structure type:
  - A Go struct with fields matching the spec
  - String getter methods for STRING fields (1-based index into []string)
  - Enum lookup maps for each Enum child table
  - Enum resolver methods on the struct for ENUM fields
  - BitField constants for each BitField child table
  - BitField description methods for Bit Field struct fields

Usage:
    python gen_smbios_go.py 3_9.json -o ./output
"""

import re
import json
import os
import sys
import argparse
from pathlib import Path


# ── Go type mapping ──────────────────────────────────────────────────────────

TYPE_MAP = {
    "BYTE": "byte",
    "WORD": "uint16",
    "DWORD": "uint32",
    "QWORD": "uint64",
}

HARDCODED_TYPES = {
    "BIOSCharacteristicsExtensionBytes": "[2]byte",
    "FirmwareCharacteristicsExtensionBytes": "[2]byte",
    "ContainedElements": "[3]byte",
}


def clean_name(name: str) -> str:
    """Remove chars that are invalid in Go identifiers."""
    for c in " ()-,./:'\"":
        name = name.replace(c, "")
    return name


def to_go_type(length: str, field_name: str) -> tuple[str, str]:
    """Map an SMBIOS length string to a Go type. Returns (go_type, extra_comment)."""
    if field_name in HARDCODED_TYPES:
        return HARDCODED_TYPES[field_name], ""

    if length in TYPE_MAP:
        return TYPE_MAP[length], ""

    m = re.match(r"(\d+)\s*bytes?", length, re.IGNORECASE)
    if m:
        return f"[{m.group(1)}]byte", ""

    return "byte", f"Type:{length}"


def parse_hex(s: str) -> int :
    """Parse a hex string like '0Ah' or '00h' to int."""
    m = re.match(r"([0-9A-Fa-f]+)h?$", s.strip())
    if m:
        return int(m.group(1), 16)
    return None


def sanitize_meaning(s: str) -> str:
    """Clean a meaning string for use in a Go identifier."""
    s = s.split(".")[0].split("(")[0].strip()
    s = re.sub(r"[^A-Za-z0-9_ ]", "", s)
    s = clean_name(s)
    return s[:60] if s else "Unknown"


# ── Struct Generation ────────────────────────────────────────────────────────

class StructGen:
    """Generates a Go struct and its methods from a main SMBIOS table."""

    def __init__(self, table: dict, version: str):
        self.table = table
        self.ver = version.replace(".", "_")
        self.rows = table["rows"]
        # Only keep children that have an explicit Enum or BitField type
        self.children = [
            c for c in table.get("children", [])
            if c.get("type") in ("Enum", "BitField")
        ]
        self.type_num = self.rows[0]["value"] if self.rows else "?"
        self.type_desc = clean_name(self.rows[0].get("description", ""))
        # e.g. Type4ProcessorInformation, Type17MemoryDevice
        desc_short = re.sub(r"(?i)indicator|structure|type", "", self.type_desc).strip()
        self.struct_name = f"Type{self.type_num}{desc_short}"

        # Track fields for dedup and method generation
        self.fields: list[dict] = []  # {name, go_type, comment, is_string, is_enum, is_bitfield}
        self._seen_names: dict[str, int] = {}
        self._build_fields()

    def _build_fields(self):
        for row in self.rows:
            name = clean_name(row["name"])
            if not name:
                continue

            value = row.get("value", "")
            go_type, extra = to_go_type(row["length"], name)

            comment = "//"
            if value in ("STRING", "ENUM", "Bit Field"):
                comment += value
            if extra:
                comment += f" {extra}"

            # Dedup field names
            if name in self._seen_names:
                self._seen_names[name] += 1
                name = f"{name}_{self._seen_names[name]}"
            else:
                self._seen_names[name] = 0

            self.fields.append({
                "name": name,
                "go_type": go_type,
                "comment": comment,
                "is_string": value == "STRING",
                "is_enum": value == "ENUM",
                "is_bitfield": value == "Bit Field",
                "raw_name": row["name"],
            })

    # ── Child table type detection ─────────────────────────────────────

    # Maps the first word of a child table's first header to a Go type
    _HEADER_TYPE_MAP = {
        "byte":  "byte",
        "word":  "uint16",
        "dword": "uint32",
        "qword": "uint64",
        "hex":   "byte",   # e.g. "Hex Value" in processor family
        "bit":   "byte",   # e.g. "Bit Position(s)"
        "value": "byte",   # generic fallback
    }

    @classmethod
    def _child_go_type(cls, child: dict) -> str:
        """Derive the Go map key / cast type from the child's first header."""
        if not child.get("headers"):
            return "byte"
        first_word = child["headers"][0].split()[0].lower()
        return cls._HEADER_TYPE_MAP.get(first_word, "byte")

    # ── Struct definition ────────────────────────────────────────────────

    def gen_struct(self) -> str:
        lines = [f"// {self.table.get('table_name', '')} (Type {self.type_num})"]
        lines.append(f"type {self.struct_name} struct {{")
        for f in self.fields:
            lines.append(f"\t{f['name']} {f['go_type']} {f['comment']}")
        lines.append("}")
        return "\n".join(lines)

    # ── String getters ───────────────────────────────────────────────────

    def gen_string_methods(self) -> str:
        blocks = []
        for f in self.fields:
            if not f["is_string"]:
                continue
            blocks.append(
                f"func (s *{self.struct_name}) Get{f['name']}(strings []string) string {{\n"
                f"\tidx := int(s.{f['name']})\n"
                f"\tif idx > 0 && idx <= len(strings) {{\n"
                f"\t\treturn strings[idx-1]\n"
                f"\t}}\n"
                f"\treturn \"\"\n"
                f"}}"
            )
        return "\n\n".join(blocks)

    # ── Enum maps + resolver methods ─────────────────────────────────────

    def gen_enum_maps(self) -> str:
        """Generate var blocks with map[T]string for each Enum child,
        where T is derived from the first header word."""
        blocks = []
        for child in self.children:
            if child.get("type") != "Enum":
                continue

            map_name = self._enum_map_name(child)
            go_key_type = self._child_go_type(child)
            entries = self._parse_enum_entries(child)
            if not entries:
                continue

            lines = [f"// {child.get('table_name', '')}"]
            lines.append(f"var {map_name} = map[{go_key_type}]string{{")
            for val, meaning in entries:
                meaning_escaped = meaning.replace('"', '\\"')
                lines.append(f'\t0x{val:02X}: "{meaning_escaped}",')
            lines.append("}")
            blocks.append("\n".join(lines))

        return "\n\n".join(blocks)

    def gen_enum_methods(self) -> str:
        """Generate resolver methods that use the enum maps."""
        blocks = []
        enum_children = [c for c in self.children if c.get("type") == "Enum"]

        for f in self.fields:
            if not f["is_enum"]:
                continue

            child = self._match_enum_child(f["raw_name"], enum_children)
            if child is None:
                continue

            map_name = self._enum_map_name(child)
            go_key_type = self._child_go_type(child)

            # Cast the struct field to the map's key type
            cast = f"{go_key_type}(s.{f['name']})" if go_key_type != f["go_type"] else f"s.{f['name']}"
            blocks.append(
                f"func (s *{self.struct_name}) Get{f['name']}String() string {{\n"
                f"\tif v, ok := {map_name}[{cast}]; ok {{\n"
                f"\t\treturn v\n"
                f"\t}}\n"
                f'\treturn "Unknown"\n'
                f"}}"
            )

        return "\n\n".join(blocks)

    # ── BitField constants + description methods ─────────────────────────

    def gen_bitfield_consts(self) -> str:
        """Generate const blocks for BitField children."""
        blocks = []
        for child in self.children:
            if child.get("type") != "BitField":
                continue

            const_prefix = f"Type{self.type_num}_{self._child_short_name(child)}"
            entries = self._parse_bitfield_entries(child)
            if not entries:
                continue

            lines = [f"// {child.get('table_name', '')}"]
            lines.append("const (")
            for bit_num, meaning in entries:
                label = sanitize_meaning(meaning)
                if not label:
                    continue
                lines.append(f"\t{const_prefix}_Bit{bit_num} = 1 << {bit_num} // {meaning[:80]}")
            lines.append(")")
            blocks.append("\n".join(lines))

        return "\n\n".join(blocks)

    def gen_bitfield_methods(self) -> str:
        """Generate methods that return a []string of set flags."""
        blocks = []
        bitfield_children = [c for c in self.children if c.get("type") == "BitField"]

        for f in self.fields:
            if not f["is_bitfield"]:
                continue

            child = self._match_bitfield_child(f["raw_name"], bitfield_children)
            if child is None:
                continue

            entries = self._parse_bitfield_entries(child)
            if not entries:
                continue

            go_cast_type = self._child_go_type(child)
            # Cast the struct field if types differ
            field_expr = f"s.{f['name']}"
            if go_cast_type != f["go_type"]:
                field_expr = f"{go_cast_type}(s.{f['name']})"

            lines = [
                f"func (s *{self.struct_name}) Get{f['name']}Flags() []string {{",
                f"\tv := {field_expr}",
                f"\tvar flags []string",
            ]
            for bit_num, meaning in entries:
                meaning_escaped = meaning.split(".")[0].strip().replace('"', '\\"')[:80]
                lines.append(
                    f'\tif v&(1<<{bit_num}) != 0 {{ flags = append(flags, "{meaning_escaped}") }}'
                )
            lines.append("\treturn flags")
            lines.append("}")
            blocks.append("\n".join(lines))

        return "\n\n".join(blocks)

    # ── Internal helpers ─────────────────────────────────────────────────

    def _enum_map_name(self, child: dict) -> str:
        return f"Type{self.type_num}_{self._child_short_name(child)}Map"

    def _child_short_name(self, child: dict) -> str:
        name = child.get("table_name", "")
        # Strip common prefixes like "Processor Information: "
        name = re.sub(r"^.*:\s*", "", name)
        name = re.sub(r"\s+field$", "", name, flags=re.IGNORECASE)
        return clean_name(name) or "Child"

    def _parse_enum_entries(self, child: dict) -> list[tuple[int, str]]:
        """Extract (int_value, meaning) pairs from an Enum child."""
        entries = []
        for row in child["rows"]:
            # Find the hex value key
            hex_val = None
            meaning = ""
            for k, v in row.items():
                kl = k.lower()
                if ("value" in kl or "hex" in kl) and kl != "decimal_value":
                    hex_val = parse_hex(v)
                if "meaning" in kl:
                    meaning = v.strip()

            if hex_val is not None and meaning:
                # Skip range entries like "06h-7Fh"
                entries.append((hex_val, meaning.split("\n")[0].strip()))

        return entries

    def _parse_bitfield_entries(self, child: dict) -> list[tuple[int, str]]:
        """Extract (bit_number, meaning) pairs from a BitField child."""
        entries = []
        for row in child["rows"]:
            bit_num = None
            meaning = ""
            for k, v in row.items():
                kl = k.lower()
                if "bit" in kl or "position" in kl:
                    m = re.search(r"Bit\s+(\d+)", v, re.IGNORECASE)
                    if m:
                        bit_num = int(m.group(1))
                if "meaning" in kl or ("description" in kl and "bit" not in kl):
                    meaning = v.strip()

            if bit_num is not None and meaning:
                entries.append((bit_num, meaning.split("\n")[0].strip()))

        return entries

    def _match_enum_child(self, field_name: str, children: list[dict]) -> dict :
        """Find the Enum child that best matches a struct field name."""
        field_lower = field_name.lower().replace("-", "").replace(" ", "")
        for c in children:
            child_name = c.get("table_name", "").lower().replace("-", "").replace(" ", "")
            # Check if the field name words appear in the child table name
            if field_lower in child_name:
                return c
        # Fallback: match by word overlap
        field_words = set(field_name.lower().split())
        best, best_score = None, 0
        for c in children:
            child_words = set(c.get("table_name", "").lower().split())
            score = len(field_words & child_words)
            if score > best_score:
                best, best_score = c, score
        return best if best_score >= 1 else None

    def _match_bitfield_child(self, field_name: str, children: list[dict]) -> dict :
        """Find the BitField child that best matches a struct field name."""
        return self._match_enum_child(field_name, children)


# ── File Generation ──────────────────────────────────────────────────────────

def gen_parser_func(struct_gens: list["StructGen"]) -> str:
    """Generate a ParseChunk function with a switch statement over all types."""
    lines = [
        "// ParsedChunk holds a parsed SMBIOS structure and its string table.",
        "type ParsedChunk struct {",
        "\tStructType byte",
        "\tData       interface{}",
        "\tStrings    []string",
        "}",
        "",
        "// ParseChunk reads an SMBIOS chunk into the appropriate typed struct.",
        "// data is the full chunk bytes (header + fields + string table).",
        "// structType is the SMBIOS structure type, length is the formatted area size.",
        "func ParseChunk(structType byte, length byte, data []byte) (*ParsedChunk, error) {",
        "\tvar obj interface{}",
        "\tswitch structType {",
    ]

    for sg in struct_gens:
        lines.append(f"\tcase {sg.type_num}:")
        lines.append(f"\t\ts := {sg.struct_name}{{}}")
        lines.append(f"\t\tif err := ReadIntoStruct(data[:length], &s); err != nil {{")
        lines.append(f'\t\t\treturn nil, fmt.Errorf("type {sg.type_num}: %w", err)')
        lines.append(f"\t\t}}")
        lines.append(f"\t\tobj = &s")

    lines.append("\tdefault:")
    lines.append("\t\treturn nil, nil")
    lines.append("\t}")
    lines.append("")
    lines.append("\tstrings := ParseNullTerminatedStrings(data[int(length):])")
    lines.append("\treturn &ParsedChunk{StructType: structType, Data: obj, Strings: strings}, nil")
    lines.append("}")

    return "\n".join(lines)


def generate_go_file(data: dict, dest_dir: Path) -> Path:
    """Generate a complete Go file from parsed SMBIOS JSON data."""
    version = data["DocumentInfo"]["Version"]
    ver_underscore = version.replace(".", "_")

    imports = (
        'import (\n'
        '\t"fmt"\n'
        ')'
    )
    sections = [f"package smbios\n\n{imports}\n"]

    struct_gens: list[StructGen] = []

    for table in data["tables"]:
        if not table.get("rows"):
            continue

        sg = StructGen(table, version)
        struct_gens.append(sg)

        parts = [sg.gen_struct()]

        string_methods = sg.gen_string_methods()
        if string_methods:
            parts.append(string_methods)

        enum_maps = sg.gen_enum_maps()
        if enum_maps:
            parts.append(enum_maps)

        enum_methods = sg.gen_enum_methods()
        if enum_methods:
            parts.append(enum_methods)

        bitfield_consts = sg.gen_bitfield_consts()
        if bitfield_consts:
            parts.append(bitfield_consts)

        bitfield_methods = sg.gen_bitfield_methods()
        if bitfield_methods:
            parts.append(bitfield_methods)

        sections.append("\n".join(parts))

    # Generate the parser switch function
    sections.append(gen_parser_func(struct_gens))

    os.makedirs(dest_dir, exist_ok=True)
    outfile = dest_dir / f"smbios_{ver_underscore}.go"
    outfile.write_text("\n\n".join(sections) + "\n", encoding="utf-8")
    return outfile


# ── CLI ──────────────────────────────────────────────────────────────────────

def main():
    argp = argparse.ArgumentParser(description="Generate Go code from parsed SMBIOS JSON")
    argp.add_argument("input", help="Path to JSON file")
    argp.add_argument("-o", "--output", help="Output directory", default="./")
    args = argp.parse_args()

    with open(args.input, "r") as f:
        data = json.load(f)

    outfile = generate_go_file(data, Path(args.output))
    print(f"Wrote {outfile}", file=sys.stderr)


if __name__ == "__main__":
    main()