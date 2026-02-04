package structures

import structuretypes "github.com/grep-michael/SMBIOS_parser/SMBiosLib/StructureTypes"

type SystemEnclosureFixed struct {
	StructureHeader
	Manufatuer                   ByteStringIndex
	Type                         ByteEnum
	Version                      ByteStringIndex
	SerialNumber                 ByteStringIndex
	AssetTagNumber               ByteStringIndex
	BootUpState                  ByteEnum
	PowerSupplyState             ByteEnum
	ThermalState                 ByteEnum
	SecurityStatus               ByteEnum
	OEMDefined                   DWord
	Height                       byte
	NumberOfPowerCords           byte
	ContainedElementCount        byte
	ContainedElementRecordLength byte
	// im not gunna try to figure out a method for creating n*m fields, fuck these goofy goobers and their variable length structures
	//probably just make a custom parse function, instead of making it a GenericStruct we make it its own thing, as long as it satisfys the SMBiosStruct interface
	//DSP0134_3.3.0.pdf pg40
}

type SystemEnclosure = GenericStruct[SystemEnclosureFixed]

func NewSystemEnclosure() SMBiosStruct {
	return &SystemEnclosure{}
}
func init() {
	Register(structuretypes.TypeSystemEnclosure, NewSystemEnclosure)
}
