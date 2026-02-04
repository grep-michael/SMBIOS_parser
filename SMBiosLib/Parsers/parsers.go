package parsers

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"log"
	"strconv"

	structuretypes "github.com/grep-michael/SMBIOS_parser/SMBiosLib/StructureTypes"
	structures "github.com/grep-michael/SMBIOS_parser/SMBiosLib/Structures"
)

// global functions used by all separate parsers

type StructureChunk struct {
	StructType          structuretypes.StructureType
	FriendlyName        string
	Start               int //position inside the smbios array
	StructureSegmentEnd int //position of structures section end, i.e headers length value
	End                 int //position from start to the double null terminators

}

var StructureMap = make(map[structuretypes.StructureType][]structures.SMBiosStruct)

func FindChunks(smbios_raw_bytes []byte) []StructureChunk {
	/*
		we assume the data is acutal smbios starting with a header
		meaning we can assume byte[2] is the length of the first struct
	*/

	var chunks []StructureChunk

	index := 0
	for index < len(smbios_raw_bytes) {
		chunk := StructureChunk{}

		chunk.Start = index
		chunk.StructType = structuretypes.StructureType(smbios_raw_bytes[index])

		friendly_name, ok := structuretypes.TypeNumToFriendlyNameMap[int(chunk.StructType)]
		if ok {
			chunk.FriendlyName = friendly_name
		} else {
			chunk.FriendlyName = strconv.Itoa(int(chunk.StructType))
		}

		length := smbios_raw_bytes[index+1]
		chunk.StructureSegmentEnd = int(length) + index

		segment_end := chunk.StructureSegmentEnd
		for segment_end < len(smbios_raw_bytes)-1 {
			if smbios_raw_bytes[segment_end] == 0x00 && smbios_raw_bytes[segment_end+1] == 0x00 {
				segment_end += 2
				break
			}
			segment_end++
		}
		chunk.End = segment_end
		chunks = append(chunks, chunk)
		index = segment_end
	}
	return chunks
}

func ParseStruct(chunk StructureChunk, smbios_raw_bytes []byte) {
	data := smbios_raw_bytes[chunk.Start:chunk.End]
	structPtr, ok := structures.GetStructForType(chunk.StructType)
	if !ok {
		log.Printf("No data Struct for type '%s'\n", chunk.FriendlyName)
		return
	}
	err := structPtr.Parse(data)
	if err != nil {
		log.Printf("Errored trying to read data into %d structure: %+v\n", chunk.StructType, err)
		return
	}

	StructureMap[chunk.StructType] = append(StructureMap[chunk.StructType], structPtr)
	log.Printf("%+v\n", structPtr)
}

func BuildEPS(base64_enocded string) (*structures.EntryPointStruct, error) {
	var entry structures.EntryPointStruct

	data, err := base64.StdEncoding.DecodeString(base64_enocded)
	if err != nil {
		log.Printf("Error Decoding EPS: %v\n", err)
		return nil, err
	}

	buf := bytes.NewReader(data)
	err = binary.Read(buf, binary.LittleEndian, &entry)
	if err != nil {
		log.Printf("Error Reading into struct: %v\n", err)
		return nil, err
	}
	return &entry, nil
}

func ParseDMITable(base64_enocded string) (map[structuretypes.StructureType][]structures.SMBiosStruct, error) {
	dmi_table, err := base64.StdEncoding.DecodeString(base64_enocded)
	if err != nil {
		log.Printf("Error Decoding DMI: %v\n", err)
		return nil, err
	}

	chunks := FindChunks(dmi_table)
	log.Printf("SMBIO data ↓\n\tChunk Count: %d\n", len(chunks))

	for _, chunk := range chunks {
		ParseStruct(chunk, dmi_table)
	}
	return StructureMap, nil
}
