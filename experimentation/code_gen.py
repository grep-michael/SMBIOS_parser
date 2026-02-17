#!/usr/bin/env python3

import argparse, os, json, re
from pathlib import Path

"""
"offset": "00h",
"spec_version": "2.0+",
"name": "Type",
"length": "BYTE",
"value": "0",
"description": "BIOS Information indicator"
"""

TYPE_MAP = {
    "BYTE": "byte",
    "WORD": "uint16",
    "DWORD": "uint32",
    "QWORD": "uint64",
}

DESIRED_TYPES = [
    "0", "1", "3", "4", "16", "17",
    "22", "24", "39", "126", "127",
    "9"
]

# Friendly names for interfaces, keyed by SMBIOS type number
INTERFACE_NAMES = {
    "0":   "BIOS",
    "1":   "System",
    "3":   "Chassis",
    "4":   "Processor",
    "16":  "PhysicalMemoryArray",
    "17":  "MemoryDevice",
    "22":  "PortableBattery",
    "24":  "HardwareSecurity",
    "39":  "PowerSupply",
    "126": "Inactive",
    "127": "EndOfTable",
}


def FilterName(name: str) -> str:
    for c in " ()-,./":
        name = name.replace(c, "")
    return name


# ── Field & Struct Generation (unchanged logic) ─────────────────────────────

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
        """Generate the interface method signature."""
        if self.is_string or self.is_enum:
            return f"{self.method_name}(strings []string) string"
        return f"{self.method_name}() {self.go_type}"

    def impl(self, struct_name: str) -> str:
        """Generate the method implementation for a concrete struct."""
        if self.is_string:
            return (
                f"func (s *{struct_name}) {self.method_name}(strings []string) string {{\n"
                f"\tidx := int(s.{self.field_name})\n"
                f"\tif idx > 0 && idx <= len(strings) {{\n"
                f"\t\treturn strings[idx-1]\n"
                f"\t}}\n"
                f"\treturn \"\"\n"
                f"}}\n"
            )
        if self.is_enum:
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
        """Find fields shared by ALL versions and create methods for them."""
        if not structs:
            return

        # Build a {field_name: StructField} map per struct
        field_maps: list[dict[str, StructField]] = []
        for s in structs:
            fm = {}
            for f in s.fields:
                fm[f.name] = f
            field_maps.append(fm)

        # Intersect field names across all versions
        common_names = set(field_maps[0].keys())
        for fm in field_maps[1:]:
            common_names &= set(fm.keys())

        # Use the first struct's field order to keep things stable
        # Skip Type/Length/Handle — they're structural, not useful as getters
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
        """Generate the Go interface definition."""
        lines = [f"type {self.iface_name} interface {{"]
        for m in self.methods:
            lines.append(f"\t{m.signature()}")
        lines.append("}")
        return "\n".join(lines)

    def gen_wrapper_string(self) -> str:
        """Generate the wrapper struct that holds the interface + strings."""
        return (
            f"type {self.wrapper_name} struct {{\n"
            f"\tInfo    {self.iface_name}\n"
            f"\tStrings []string\n"
            f"}}"
        )

    def gen_impl_string(self, struct_name: str) -> str:
        """Generate all method implementations for a specific struct."""
        lines = []
        for m in self.methods:
            lines.append(m.impl(struct_name))
        return "\n".join(lines)


# ── Code Generator ───────────────────────────────────────────────────────────

class CodeGenerator:
    def __init__(self, Namespace):
        self.input_path: Path = Path(Namespace.input)
        self.dest_dir: Path = Path(Namespace.output).joinpath("CodeGen")
        # {version_str: {type_num: GoStruct}}
        self.version_structs: dict[str, dict[str, GoStruct]] = {}
        self.version_data: dict[str, dict] = {}  # raw json per version

    def generate(self):
        os.makedirs(self.dest_dir, exist_ok=True)

        # Load all JSON files from input (file or directory)
        json_files = self._collect_json_files()
        if not json_files:
            print(f"No JSON files found at {self.input_path}")
            return

        for jf in json_files:
            self._load_version(jf)

        # Generate per-version struct files
        for ver, data in self.version_data.items():
            self._write_struct_file(ver, data)

        # Generate shared interfaces file
        self._write_interfaces_file()

    def _collect_json_files(self) -> list[Path]:
        """Collect JSON files from a file path or directory."""
        if self.input_path.is_file():
            return [self.input_path]
        if self.input_path.is_dir():
            return sorted(self.input_path.glob("*.json"))
        return []

    def _load_version(self, filepath: Path):
        """Load a single JSON file and build GoStructs for it."""
        with open(filepath, "r") as f:
            data = json.load(f)

        version = data["DocumentInfo"]["Version"]
        self.version_data[version] = data

        structs: dict[str, GoStruct] = {}
        for table in data["tables"]:
            gs = GoStruct(table, version)
            if gs.StructNumber in DESIRED_TYPES:
                structs[gs.StructNumber] = gs

        self.version_structs[version] = structs

    def _write_struct_file(self, version: str, data: dict):
        """Write the per-version Go struct file with method implementations."""
        ver_underscore = version.replace(".", "_")
        structs = self.version_structs[version]

        code = "package GeneratedCode\n\n"

        # Write struct definitions
        for type_num in DESIRED_TYPES:
            gs = structs.get(type_num)
            if gs is None:
                continue
            code += gs.gen_struct_string() + "\n"

        # Write method implementations for shared interfaces
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
        print(f"Wrote {outfile}")

    def _write_interfaces_file(self):
        """Write the shared interfaces and wrapper structs to a single file."""
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
        print(f"Wrote {outfile}")

    def _build_interfaces(self) -> dict[str, InterfaceGenerator]:
        """Build SharedInterface objects for each desired type across all versions."""
        interfaces: dict[str, InterfaceGenerator] = {}
        for type_num in DESIRED_TYPES:
            # Collect all GoStructs for this type across versions
            all_structs = []
            for ver_structs in self.version_structs.values():
                gs = ver_structs.get(type_num)
                if gs is not None:
                    all_structs.append(gs)

            if all_structs:
                interfaces[type_num] = InterfaceGenerator(type_num, all_structs)

        return interfaces


# ── CLI ──────────────────────────────────────────────────────────────────────

def main():
    argp = argparse.ArgumentParser(
        description="Generate Go code from parsed SMBIOS JSON files"
    )
    argp.add_argument(
        "input",
        help="Path to a JSON file or directory of JSON files",
    )
    argp.add_argument(
        "-o", "--output",
        help="Output folder (default: ./)",
        default="./",
    )
    args = argp.parse_args()
    generator = CodeGenerator(args)
    generator.generate()


if __name__ == "__main__":
    main()