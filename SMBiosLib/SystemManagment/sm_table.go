package sm

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"

	eps "github.com/grep-michael/SMBIOS_parser/SMBiosLib/Structures/EPS"
	"github.com/grep-michael/SMBIOS_parser/SMBiosLib/Structures/StructureParsing"
)

type SMTable struct {
	DMI_TABLE_Bytes []byte
	EPS_Bytes       []byte
	EPS             *eps.EntryPointStruct
	Structures      map[int][]*parsing.ParsedChunk
}

func NewSMBiosData(eps_bytes []byte, dmiTable_bytes []byte) *SMTable {
	return &SMTable{
		EPS_Bytes:       eps_bytes,
		DMI_TABLE_Bytes: dmiTable_bytes,
	}
}
func NewSMBiosDataB64(eps_b64string string, dmiTable_b64string string) (*SMTable, error) {
	tableBytes, err := base64.StdEncoding.DecodeString(dmiTable_b64string)
	if err != nil {
		log.Printf("Error Decoding DMI Table: %v\n", err)
		return nil, err
	}
	headerBytes, err := base64.StdEncoding.DecodeString(eps_b64string)
	if err != nil {
		log.Printf("Error Decoding EPS: %v\n", err)
		return nil, err
	}

	return &SMTable{
		EPS_Bytes:       tableBytes,
		DMI_TABLE_Bytes: headerBytes,
	}, nil
}

/*
Fresh load structures every time we marashal, this is because they are interface{} objs and the marshaler converts them to maps everytime
*/
func (table *SMTable) UnmarshalJSON(data []byte) error {
	anon := struct {
		DMI_TABLE_Bytes []byte
		EPS_Bytes       []byte
		EPS             *eps.EntryPointStruct
		Structures      map[int][]*parsing.ParsedChunk `json:"-"`
	}{}
	if err := json.Unmarshal(data, &anon); err != nil {
		return err
	}
	table.DMI_TABLE_Bytes = anon.DMI_TABLE_Bytes
	table.EPS_Bytes = anon.EPS_Bytes
	table.EPS = anon.EPS
	return table.LoadStructures()
}

func (data *SMTable) LoadStructures() error {
	if len(data.DMI_TABLE_Bytes) <= 1 {
		return fmt.Errorf("DMI_TABLE_Bytes Empty, failed to load")
	}
	chunks, err := parsing.ParseSMBiosBytes(data.DMI_TABLE_Bytes)
	if err != nil {
		return err
	}
	data.Structures = chunks
	return nil
}
func (data *SMTable) LoadEPS() error {
	if data.EPS_Bytes == nil {
		return fmt.Errorf("Populate .EPS_Bytes First")
	}
	data.EPS = eps.NewEPS(data.EPS_Bytes)
	log.Printf("Built New EPS: Version %d\n", data.EPS.Version)
	return nil
}
func (data *SMTable) VerifyDMITable() error {
	if data.EPS == nil {
		if err := data.LoadEPS(); err != nil {
			return err
		}
	}
	return data.EPS.VerifyDMITable(data.DMI_TABLE_Bytes)
}
