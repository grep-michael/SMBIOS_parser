package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/grep-michael/SMBIOS_parser/SMBiosLib/Parsers"
	structuretypes "github.com/grep-michael/SMBIOS_parser/SMBiosLib/StructureTypes"
	structs_lib "github.com/grep-michael/SMBIOS_parser/SMBiosLib/Structures"
)

//100000000000000
//10000000000000

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(file)

	smbios_bytes, eps_bytes := buildByteArrays()
	//smbios_bytes, eps_bytes, err := loadLocalSMBIOS()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	fmt.Printf("SMBIOS len: %d\n", len(smbios_bytes))
	fmt.Printf("SMBIOS_EPS len: %d\n", len(eps_bytes))

	eps := buildEPS(eps_bytes)
	log.Printf("EPS ↓\n\t%+v\n", eps)
	chunks := parsers.FindChunks(smbios_bytes)
	fmt.Printf("SMBIO data ↓\n\tChunk Count: %d\n\tExpected Count: %d\n", len(chunks), eps.NumOfStructs)

	for _, chunk := range chunks {
		parsers.ParseStruct(chunk, smbios_bytes)
	}
	type_arg, _ := strconv.Atoi(os.Args[1])
	struct_type_arg := structuretypes.StructureType(type_arg)

	fmt.Println()
	fmt.Printf("All %d structs: %d\n", type_arg, len(parsers.StructureMap[struct_type_arg]))

	for _, structure := range parsers.StructureMap[struct_type_arg] {
		struct_info := structure.(*structs_lib.MemoryDeviceInfo)
		fmt.Println(struct_info.Data.Size)
		printObj(struct_info)

	}

}

func printObj(obj any) {
	json, _ := json.MarshalIndent(obj, "", "  ")
	fmt.Println(string(json))
}

func loadLocalSMBIOS() (dmi_table []byte, eps []byte, err error) {
	if dmi_table, err = os.ReadFile("/sys/firmware/dmi/tables/DMI"); err == nil {
		return
	}
	if eps, err = os.ReadFile("/sys/firmware/dmi/tables/smbios_entry_point"); err == nil {
		return
	}
	return
}

func buildEPS(data []byte) *structs_lib.EntryPointStruct {
	var entry structs_lib.EntryPointStruct

	data, err := base64.StdEncoding.DecodeString(SMBIOS_EPS)
	if err != nil {
		log.Printf("Error Decoding EPS: %v\n", err)
		os.Exit(1)
	}

	buf := bytes.NewReader(data)
	err = binary.Read(buf, binary.LittleEndian, &entry)
	if err != nil {
		log.Printf("Error Reading into struct: %v\n", err)
		os.Exit(1)
	}
	return &entry
}

func buildByteArrays() (smbios []byte, eps []byte) {
	var err error
	eps, err = base64.StdEncoding.DecodeString(SMBIOS_EPS)
	if err != nil {
		log.Printf("Error Decoding EPS: %v\n", err)
		os.Exit(1)
	}

	smbios, err = base64.StdEncoding.DecodeString(SMBIOS)
	if err != nil {
		log.Printf("Error Decoding EPS: %v\n", err)
		os.Exit(1)
	}
	return
}
