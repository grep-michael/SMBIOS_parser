#!/usr/bin/env python3

import argparse,os,json
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
    "BYTE":"byte",
    "WORD":"uint16",
    "DWORD":"uint32",
    "QWORD":"uint64"
}

def FilterName(name:str) -> str:
    chacters_to_remove = " ()-,./"#[" ","(",")","-",]
    for c in chacters_to_remove:
        name = name.replace(c, "")
    return name

class StructField():
    def __init__(self,row_table:dict):
        self.row:dict = row_table
        self.comment:str = self._gen_comment()
        self.type:str = self._gen_type()
        self.name:str = FilterName(self.row["name"])
    
    def _gen_comment(self):
        comment = "//"
        value = self.row["value"]
        value = value if value in ["STRING", "Bit Field","ENUM"] else ""
        comment += value
        #if len(self.row["description"]) < 75:
        #    comment += " "
        #    comment += self.row["description"] 
        return comment
    
    def _gen_type(self)->str:
        length = self.row["length"]
        return TYPE_MAP.get(length,"byte")

    def gen_field_string(self) -> str:
        return f"{self.name} {self.type} {self.comment}" 

class GoStruct():
    def __init__(self,table:dict):
        self.table = table
        self.StructName: str = ""
        self.fields:list[StructField] = []
        self._gen_name()
        self._gen_fields()

    def _gen_name(self):
        name:str = self.table["table_name"].replace(" ","")
        self.StructName = FilterName(name)
        self.StructName = "S_" + self.StructName.replace("structure","").replace("Structure","")

    def _gen_fields(self):
        rows:list = self.table["rows"]
        field_names = {}
        count = 0
        for row in rows:
            field = StructField(row)
            if field.name not in field_names:
                field_names[field.name] = "present"
            else:
                field.name = field.name + "_" +str(count) 
                count += 1
            
            self.fields.append(field)

    def gen_struct_string(self) -> str:
        definition = f"type {self.StructName} struct {{\n"
        for row in self.fields:
            field = row.gen_field_string()
            definition += f"\t{field}\n"
        definition+= "}"
        return definition

class CodeGenerator():
    def __init__(self,Namespace):
        self.filepath:Path = Path(Namespace.input)
        self.dest_dir:Path = Path(Namespace.output).joinpath("CodeGen")
        self.json_struct:dict = {}
        self.structs:list[GoStruct] = []
    
    def generate(self) -> str:
        os.makedirs(self.dest_dir,exist_ok=True)
        with open(self.filepath, 'r') as file:
            self.json_struct = json.load(file)
        
        tables = self.json_struct["tables"]
        for table in tables:
            self._gen_struct_table(table)
        
        code = "package GeneratedCode\n\n"
        for struct in self.structs:
            code += struct.gen_struct_string()
            code += "\n"
        
        outfile = self.dest_dir.joinpath(self._gen_output_filename())
        with open(outfile,"w") as f:
            f.write(code)
    

    def _gen_output_filename(self) -> str:
        version:str = self.json_struct["DocumentInfo"]["Version"]
        version = version.replace(".","_")
        return f"smbios_{version}.go"

    def _gen_struct_table(self,table:dict):
        struct = GoStruct(table)
        self.structs.append(struct)
        

def main():
    argp = argparse.ArgumentParser(
        description="generate code from parsed json"
    )
    argp.add_argument("input", help="Path to generated json file")
    argp.add_argument(
        "-o", "--output",
        help="Output folder",
        default="./",
    )
    args = argp.parse_args()
    generator = CodeGenerator(args)
    generator.generate()



if __name__ == "__main__":
    main()