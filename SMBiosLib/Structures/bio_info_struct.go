package structures

type ByteString byte

func (bs ByteString) String() string {
	return string(bs)
}

/*
BIOS Information
Default Header length: 12
*/
func (bi BiosInfo) GetDefaultLength() uint8 {
	return 0x12
}

type BiosInfo struct {
	StructureHeader
	Vendor                   ByteString
	BiosVersion              ByteString
	BiosStartAddressSegment  uint16
	BiosReleaseDate          ByteString
	BiosRomSize              byte
	BiosCharacteristics      uint64
	CharacteristicsExtension []byte
}
