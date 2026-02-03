package structures

import structuretypes "github.com/grep-michael/SMBIOS_parser/SMBiosLib/StructureTypes"

type StructureHeader struct {
	Type   structuretypes.StructureType
	Length byte // depending on the struct there will be a default length assuming no extensions bytes are used
	Handle uint16
}

type TestStruct struct {
	StructureHeader
}
