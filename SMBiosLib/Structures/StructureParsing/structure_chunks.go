package parsing

import (
	"log"
)

type RawChunk struct {
	StructType          int
	Length              int
	FriendlyName        string
	Data                []byte
	Start               int //position inside the smbios array
	StructureSegmentEnd int //position of structures section end, i.e headers length value
	End                 int //position from start to the double null terminators
}

type ParsedChunk struct {
	StructType byte
	Strings    []string
	Data       interface{}
}

func ParseSMBiosBytes(data []byte) (map[int][]*ParsedChunk, error) {
	rawChunks, err := buildRawChunkList(data)
	if err != nil {
		return nil, err
	}
	parsedChunks, err := buildParsedChunkMap(rawChunks)
	return parsedChunks, err
}
func buildRawChunkList(data []byte) ([]RawChunk, error) {
	var rawChunkList []RawChunk
	index := 0
	for index < len(data) {
		chunk := RawChunk{}

		chunk.Start = index
		chunk.StructType = int(data[index])
		chunk.FriendlyName = TypeNumToFriendlyNameMap[chunk.StructType]
		chunk.Length = int(data[index+1])
		chunk.StructureSegmentEnd = int(chunk.Length) + index

		segment_end := chunk.StructureSegmentEnd
		for segment_end < len(data)-1 {
			if data[segment_end] == 0x00 && data[segment_end+1] == 0x00 {
				segment_end += 2
				break
			}
			segment_end++
		}
		chunk.End = segment_end
		chunk.Data = data[chunk.Start:chunk.End]
		rawChunkList = append(rawChunkList, chunk)
		index = segment_end
	}
	return rawChunkList, nil

}
func buildParsedChunkMap(chunks []RawChunk) (map[int][]*ParsedChunk, error) {
	parsedChunks := make(map[int][]*ParsedChunk)
	for _, chunk := range chunks {
		prased_chunk, err := ParseChunk(byte(chunk.StructType), byte(chunk.Length), chunk.Data)
		if err != nil {
			log.Printf("Failed to parse chunk %d\n", chunk.StructType)
			continue
		}
		parsedChunks[chunk.StructType] = append(parsedChunks[chunk.StructType], prased_chunk)
		//log.Printf("Found struct %d\n", chunk.StructType)
	}
	return parsedChunks, nil
}
