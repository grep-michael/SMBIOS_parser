package structures

import (
	structuretypes "github.com/grep-michael/SMBIOS_parser/SMBiosLib/StructureTypes"
)

type MemoryDeviceFixed struct {
	StructureHeader
	PhysicalmemoryArrayHandle Word
	MemoryErrorHandler        Word
	TotalWidth                Word
	DataWidth                 Word
	Size                      Word
	FormFactor                ByteEnum
	DeiveSet                  byte
	DeviceLocator             ByteStringIndex
	BankLocator               ByteStringIndex
	MemoryType                ByteEnum
	TypeDetail                Word
	Speed                     Word
	Manufacturer              ByteStringIndex
	SerialNumber              ByteStringIndex
	AssetTag                  ByteStringIndex
	PartNumber                ByteStringIndex
	Attributes                byte
	ExtendedSize              DWord
	ConfiguredMemorySpeed     Word
	MinVoltage                Word
}

type MemoryDeviceInfo = GenericStruct[MemoryDeviceFixed]

func NewMemoryInfo() SMBiosStruct {
	return &MemoryDeviceInfo{}
}

func init() {
	Register(structuretypes.TypeMemoryDevice, NewMemoryInfo)
}
