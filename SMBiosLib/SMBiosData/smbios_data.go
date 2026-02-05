package smbiosdata

import (
	"encoding/base64"
	"log"

	eps "github.com/grep-michael/SMBIOS_parser/SMBiosLib/Strucutres/EPS"
)

type SMBiosData struct {
	DMI_TABLE_b64S  string //base64 encoded DMI Table
	EPS_b64S        string //base64 encoded EPS data
	DMI_TABLE_Bytes []byte
	EPS_Bytes       []byte
	EPS             *eps.EntryPointStruct
	Chunks          []StructureChunk
}

type StructureChunk struct {
	StructType          int
	FriendlyName        string
	Start               int //position inside the smbios array
	StructureSegmentEnd int //position of structures section end, i.e headers length value
	End                 int //position from start to the double null terminators

}

func (data *SMBiosData) DecodeBase64Fields() (err error) {
	data.DMI_TABLE_Bytes, err = base64.StdEncoding.DecodeString(data.DMI_TABLE_b64S)
	if err != nil {
		log.Printf("Error Decoding DMI Table: %v\n", err)
		return err
	}
	data.EPS_Bytes, err = base64.StdEncoding.DecodeString(data.EPS_b64S)
	if err != nil {
		log.Printf("Error Decoding EPS: %v\n", err)
		return err
	}
	return nil
}
func (data *SMBiosData) LoadDMITableChunks() error {
	data.Chunks = make([]StructureChunk, 0)
	if len(data.DMI_TABLE_Bytes) <= 1 {
		err := data.DecodeBase64Fields()
		if err != nil {
			log.Println("Error decoding base64 fields")
			return err
		}
	}

	index := 0
	for index < len(data.DMI_TABLE_Bytes) {
		chunk := StructureChunk{}

		chunk.Start = index
		chunk.StructType = int(data.DMI_TABLE_Bytes[index])

		length := data.DMI_TABLE_Bytes[index+1]
		chunk.StructureSegmentEnd = int(length) + index

		segment_end := chunk.StructureSegmentEnd
		for segment_end < len(data.DMI_TABLE_Bytes)-1 {
			if data.DMI_TABLE_Bytes[segment_end] == 0x00 && data.DMI_TABLE_Bytes[segment_end+1] == 0x00 {
				segment_end += 2
				break
			}
			segment_end++
		}
		chunk.End = segment_end
		data.Chunks = append(data.Chunks, chunk)
		index = segment_end
	}
	return nil
}
func (data *SMBiosData) LoadEPSStruct() error {
	if len(data.EPS_Bytes) <= 1 {
		err := data.DecodeBase64Fields()
		if err != nil {
			log.Println("Error decoding base64 fields")
			return err
		}
	}
	data.EPS = eps.NewEPS(data.EPS_Bytes)
	return nil
}
