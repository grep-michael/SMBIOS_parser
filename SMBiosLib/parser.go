package smbioslib

import (
	"bytes"
	"encoding/binary"
	"log"
)

type StructureChunk struct {
	Header  StructureHeader
	Body    interface{}
	Strings []string
}

func ParseChunk(header *StructureHeader, body []byte) (*StructureChunk, error) {
	chunk := &StructureChunk{}
	/*
		We define a chunck from the start of the header (type field) to the double null termination
	*/
	var structPtr interface{}
	switch header.Type {
	case 4:
		structPtr = &ProcessorInfo{}
	default:
		structPtr = &TestStruct{}
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

func parseNullTerminatedStrings(data []byte) []string {
	var strings []string
	start := 0

	for i := 0; i < len(data)-1; i++ {
		if data[i] == 0x00 {
			if data[i+1] == 0x00 {
				// Double null - end of strings
				if i > start {
					strings = append(strings, string(data[start:i]))
				}
				break
			}
			// Single null - end of current string
			if i > start {
				strings = append(strings, string(data[start:i]))
			}
			start = i + 1
		}
	}

	return strings
}
