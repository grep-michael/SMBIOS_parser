package dmitabel

import (
	"log"

	processor "github.com/grep-michael/SMBIOS_parser/SMBiosLib/Strucutres/Processor"
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
	Processors  []*processor.Processor
	chunks      []StructureChunk
	rawDMITable []byte
}

func NewDMITable() *DMITable {
	table := &DMITable{
		chunks: make([]StructureChunk, 0),
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
		switch chunk.StructType {
		case 4:
			proc, err := processor.ParseChunk(chunk.Data)
			if err != nil {
				log.Printf("Failed to parse processor chunk: %+v\n", err)
				continue
			}
			table.Processors = append(table.Processors, proc)
		default:
			//log.Printf("No parsing for chunk of type %d %s\n", chunk.StructType, chunk.FriendlyName)
		}
	}
}
