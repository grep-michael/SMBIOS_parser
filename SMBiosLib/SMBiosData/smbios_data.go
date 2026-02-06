package smbiosdata

import (
	"encoding/base64"
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
	if len(data.DMI_TABLE_Bytes) <= 1 {
		err := data.DecodeBase64Fields()
		if err != nil {
			return err
		}
	}
	data.DMITable = dmitabel.NewDMITable()
	return data.DMITable.BuildStructs(data.DMI_TABLE_Bytes)
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
