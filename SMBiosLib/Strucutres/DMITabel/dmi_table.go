package dmitabel

import (
	"log"
)

type StructureChunk struct {
	StructType          int
	Length              int
	FriendlyName        string
	Data                []byte
	Start               int //position inside the smbios array
	StructureSegmentEnd int //position of structures section end, i.e headers length value
	End                 int //position from start to the double null terminators

}

type DMITable struct {
	Structs     map[int][]interface{}
	chunks      []StructureChunk
	rawDMITable []byte
}

func NewDMITable() *DMITable {
	table := &DMITable{
		chunks:  make([]StructureChunk, 0),
		Structs: make(map[int][]interface{}),
	}
	return table
}

func (table *DMITable) BuildStructs(data []byte) error {
	err := table.buildChunkList(data)
	if err != nil {
		return err
	}
	table.parseChunkList()
	return nil
}
func (table *DMITable) buildChunkList(data []byte) error {
	table.rawDMITable = data
	index := 0
	for index < len(data) {
		chunk := StructureChunk{}

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
		table.chunks = append(table.chunks, chunk)
		index = segment_end
	}
	return nil

}
func (table *DMITable) parseChunkList() {
	for _, chunk := range table.chunks {
		prased_chunk, err := ParseChunk(byte(chunk.StructType), byte(chunk.Length), chunk.Data)
		if err != nil {
			log.Printf("Failed to parse chunk %d\n", chunk.StructType)
			continue
		}
		table.Structs[chunk.StructType] = append(table.Structs[chunk.StructType], prased_chunk)
	}
}
