package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"unsafe"

	smbioslib "github.com/grep-michael/SMBIOS_parser/SMBiosLib"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	smbios_bytes, eps_bytes := buildByteArrays()

	fmt.Printf("SMBIOS len: %d\n", len(smbios_bytes))
	//fmt.Printf("%v\n", smbios_bytes[:50])
	//fmt.Printf("%v\n", smbios_bytes[:23])
	//fmt.Printf("%v\n", smbios_bytes[23:50])

	fmt.Printf("SMBIOS_EPS len: %d\n", len(eps_bytes))
	//fmt.Printf("%v\n", eps_bytes[:10])
	eps := buildEPS(eps_bytes)
	log.Printf("Entry:\n\t%+v\n", eps)
	smb_structures := parseHeaders(smbios_bytes)
	fmt.Println(len(smb_structures))

	for _, s := range smb_structures {
		switch s.Body.(type) {
		case *smbioslib.ProcessorInfo:
			fmt.Printf("%+v\n", s.Body)
		}
	}
}

func parseHeaders(smbios []byte) []*smbioslib.StructureChunk {
	var structs []*smbioslib.StructureChunk
	offset := 0

	for offset < len(smbios) {
		if offset+int(unsafe.Sizeof(smbioslib.StructureHeader{})) > len(smbios) {
			break
		}

		header := &smbioslib.StructureHeader{}
		buf := bytes.NewReader(smbios[offset:])
		err := binary.Read(buf, binary.LittleEndian, header)

		if err != nil {
			log.Printf("Error reading into header: %+v\n", err)
			os.Exit(1)
		}
		if header.Length == 0 || offset+int(header.Length) > len(smbios) {
			log.Printf("Invaild Header Length")
			os.Exit(1)
		}

		//log.Printf("Header found: %+v\n", header)
		scanOffset := offset + int(header.Length)
		// Find double null terminator
		/*
			apples SMB struct is all fucked up and retarded, i genuinely think they might John Apple himself said
			"Lets make our own fucked up reatrded version of SMBIOs just to fuck with people"
		*/
		for scanOffset < len(smbios)-1 {
			if smbios[scanOffset] == 0x00 && smbios[scanOffset+1] == 0x00 {
				scanOffset += 2
				break
			}
			scanOffset++
		}

		chunk, err := smbioslib.ParseChunk(header, smbios[offset:scanOffset])
		if err == nil {
			structs = append(structs, chunk)
		}

		offset = scanOffset

	}

	return structs
}

func buildEPS(data []byte) *smbioslib.EntryPointStruct {
	var entry smbioslib.EntryPointStruct

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
