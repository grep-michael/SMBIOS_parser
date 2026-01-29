package structures

import (
	typemap "github.com/grep-michael/SMBIOS_parser/SMBiosLib/TypeMap"
)

type BiosInfo struct {
	StructureHeader
	Vendor                                 ByteStringIndex
	BiosVersion                            ByteStringIndex
	BiosStartAddressSegment                uint16
	BiosReleaseDate                        ByteStringIndex
	BiosRomSize                            byte
	BiosCharacteristics                    uint64
	CharacteristicsExtension               uint16 //3.3 defines this as zero or more bytes? at offset 12h, but then also says there more structured data at offset 14h?
	SystemBiosMajorRelease                 byte
	SystemBiosMinorRelease                 byte
	EmbeddedControllerFirmwareMajorRelease byte
	EmbeddedControllerFirmwareMinorRelease byte
	ExtendedBiosRomSize                    uint16
}

func NewBioInfo() *BiosInfo {
	return &BiosInfo{}
}

func init() {
	typemap.Register(0, NewBioInfo)
}
