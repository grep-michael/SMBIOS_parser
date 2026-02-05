package structures

import structuretypes "github.com/grep-michael/SMBIOS_parser/SMBiosLib/StructureTypes"

type BaseboardInfoFixed struct {
	StructureHeader
	Manufacturer                   ByteStringIndex
	Product                        ByteStringIndex
	Version                        ByteStringIndex
	SerialNumber                   ByteStringIndex
	AssetTag                       ByteStringIndex
	FeatureFlag                    byte
	LocationInChassis              ByteStringIndex
	ChassisHandle                  Word
	BoardType                      ByteEnum
	NumberOfContainedObjectHandles byte
}

type BaseboardInfo = GenericStruct[BaseboardInfoFixed]

func NewBasebaordInfo() SMBiosStruct {
	return &BaseboardInfo{}
}
func init() {
	Register(structuretypes.TypeBaseboardInfo, NewBasebaordInfo)
}
