package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"log"
	"os"

	"github.com/grep-michael/SMBIOS_parser/SMBiosLib/Parsers"
	structs_lib "github.com/grep-michael/SMBIOS_parser/SMBiosLib/Structures"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(file)

	smbios_bytes, eps_bytes := buildByteArrays()

	fmt.Printf("SMBIOS len: %d\n", len(smbios_bytes))
	fmt.Printf("SMBIOS_EPS len: %d\n", len(eps_bytes))

	eps := buildEPS(eps_bytes)
	log.Printf("EPS ↓\n\t%+v\n", eps)
	chunks := parsers.FindChunks(smbios_bytes)
	fmt.Printf("SMBIO data ↓\n\tChunk Count: %d\n\tExpected Count: %d\n", len(chunks), eps.NumOfStructs)

	for _, chunk := range chunks {
		parsers.ParseStruct(chunk, smbios_bytes)
	}

	fmt.Printf("All Bios Info structs: %d\n", len(parsers.StructureMap[0]))
	for _, bioInfo := range parsers.StructureMap[0] {
		bioinfo := bioInfo.(*structs_lib.BiosInfo)
		fmt.Printf("\t%+v\n", bioinfo)
		chars := parsers.ParseBIOSCharacteristics(bioinfo.BiosCharacteristics)
		fmt.Printf("\t%+v\n", chars)
	}

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
