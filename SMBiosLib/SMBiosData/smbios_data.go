package smbiosdata

import (
	"encoding/base64"
	"fmt"
	"log"

	dmitabel "github.com/grep-michael/SMBIOS_parser/SMBiosLib/Strucutres/DMITabel"
	eps "github.com/grep-michael/SMBIOS_parser/SMBiosLib/Strucutres/EPS"
)

type SMBiosData struct {
	DMI_TABLE_b64S  string //base64 encoded DMI Table
	EPS_b64S        string //base64 encoded EPS data
	DMI_TABLE_Bytes []byte
	EPS_Bytes       []byte
	EPS             *eps.EntryPointStruct
	DMITable        *dmitabel.DMITable
}

func NewSMBiosData(eps_bytes []byte, dmiTable_bytes []byte) *SMBiosData {
	return &SMBiosData{
		EPS_Bytes:       eps_bytes,
		DMI_TABLE_Bytes: dmiTable_bytes,
	}
}
func NewSMBiosDataB64(eps_b64string string, dmiTable_b64string string) *SMBiosData {
	return &SMBiosData{
		EPS_b64S:       eps_b64string,
		DMI_TABLE_b64S: dmiTable_b64string,
	}
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
func (data *SMBiosData) LoadDMITable() error {

	if len(data.DMI_TABLE_Bytes) <= 1 || data.DMI_TABLE_b64S != "" {
		err := data.DecodeBase64Fields()
		if err != nil {
			return err
		}
	}
	if len(data.DMI_TABLE_Bytes) <= 1 {
		return fmt.Errorf("Populate .DMI_TABLE_Bytes First")
	}
	data.DMITable = dmitabel.NewDMITable()
	return data.DMITable.PraseByteSlice(data.DMI_TABLE_Bytes)
}
func (data *SMBiosData) LoadEPSStruct() error {
	if len(data.EPS_Bytes) <= 1 || data.EPS_b64S != "" {
		err := data.DecodeBase64Fields()
		if err != nil {
			log.Println("Error decoding base64 fields")
			return err
		}
	}
	if data.EPS_Bytes == nil {
		return fmt.Errorf("Populate .EPS_Bytes First")
	}
	data.EPS = eps.NewEPS(data.EPS_Bytes)
	log.Printf("Built New EPS: Version %d\n", data.EPS.Version)
	return nil
}
func (data *SMBiosData) VerifyDMITable() error {
	if data.EPS == nil {
		if err := data.LoadEPSStruct(); err != nil {
			return err
		}
	}
	return data.EPS.VerifyDMITable(data.DMI_TABLE_Bytes)
}
