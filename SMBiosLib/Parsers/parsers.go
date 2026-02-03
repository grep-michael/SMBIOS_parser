package parsers

import (
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
		log.Printf("Errored trying to read body buffer into structure: %+v\n", err)
		return
	}

	StructureMap[chunk.StructType] = append(StructureMap[chunk.StructType], structPtr)
	log.Printf("%+v\n", structPtr)
}
