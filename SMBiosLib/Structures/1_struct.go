package structures

import structuretypes "github.com/grep-michael/SMBIOS_parser/SMBiosLib/StructureTypes"

type SystemInformationFixed struct {
	StructureHeader
	Manufacturer ByteStringIndex
	ProductName  ByteStringIndex
	Version      ByteStringIndex
	SerialNumber ByteStringIndex
	UUID         UUID
	WakeUpType   byte //enum
	SKUNumber    ByteStringIndex
	Family       ByteStringIndex
}
type UUID struct {
	TimeLow               uint32
	TimeMid               uint16
	TimeHiAndVersion      uint16
	ClockSeqHiAndReserved byte
	ClockSeqLow           byte
	Node                  [6]byte
}

type SystemInformation = GenericStruct[SystemInformationFixed]

func NewSysteminfo() SMBiosStruct {
	return &SystemInformation{}
}
func init() {
	Register(structuretypes.TypeSystemInfo, NewSysteminfo)
}
