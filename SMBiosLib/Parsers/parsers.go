package parsers

import (
	"bytes"
	"encoding/binary"

	structs_lib "github.com/grep-michael/SMBIOS_parser/SMBiosLib/Structures"
)

// global functions used by all separate parsers

func ParseChunk(header *structs_lib.StructureHeader, body []byte) (*structs_lib.StructureChunk, error) {
	chunk := &structs_lib.StructureChunk{}
	/*
		We define a chunck from the start of the header (type field) to the double null termination
	*/
	var structPtr interface{}
	switch header.Type {
	case 4:
		structPtr = &structs_lib.ProcessorInfo{}
	default:
		structPtr = &structs_lib.TestStruct{}
	}

	body_buffer := bytes.NewReader(body)
	err := binary.Read(body_buffer, binary.LittleEndian, structPtr)
	if err != nil {
		log.Println(body)
		log.Printf("Errored trying to read body buffer into structure: %+v\n", err)

		return nil, err
	}

	chunk.Body = structPtr
	chunk.Strings = parseNullTerminatedStrings(body[header.Length:])

	return chunk, nil

}

// parseNullTerminatedStrings Given a byte array will split strings by null bytes
func parseNullTerminatedStrings(data []byte) []string {
	var strings []string
	start := 0

	for i := 0; i < len(data)-1; i++ {
		if data[i] == 0x00 {
			if data[i+1] == 0x00 { //double null, end of section
				if i > start {
					strings = append(strings, string(data[start:i]))
				}
				break
			}
			if i > start {
				strings = append(strings, string(data[start:i]))
			}
			start = i + 1
		}
	}
	return strings
}
