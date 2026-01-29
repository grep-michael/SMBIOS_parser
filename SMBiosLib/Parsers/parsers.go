package parsers

import (
	"bytes"
	"encoding/binary"
	"github.com/grep-michael/SMBIOS_parser/SMBiosLib/TypeMap"
	"log"
)

// global functions used by all separate parsers

type StructureChunk struct {
	StructType          int
	Start               int //position inside the smbios array
	StructureSegmentEnd int //position of structures section end, i.e headers length value
	End                 int //position from start to the double null terminators

}

var StructureMap map[int][]interface{} = make(map[int][]interface{})

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
		chunk.StructType = int(smbios_raw_bytes[index])
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
	structPtr, ok := typemap.GetStructForType(chunk.StructType)
	if !ok {
		log.Printf("No data Struct for type %d\n", chunk.StructType)
		return
	}

	data_buff := bytes.NewReader(data)
	err := binary.Read(data_buff, binary.LittleEndian, structPtr)
	if err != nil {
		log.Printf("Errored trying to read body buffer into structure: %+v\n", err)
		return
	}
	StructureMap[chunk.StructType] = append(StructureMap[chunk.StructType], structPtr)
	log.Printf("%+v\n", structPtr)
	return
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
