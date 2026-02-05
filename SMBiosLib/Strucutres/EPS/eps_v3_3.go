package eps

// the eps table in accordance with the 3.3 smbios standard

import (
	"fmt"
)

// 32 bit Entry Point struct
type EntryPointStruct_2_1 struct {
	ArchorString         AnchorString32bit       //4 bytes
	Checksum             uint8                   //byte
	EntryPointLength     uint8                   //byte
	MajorVer             uint8                   //byte
	MinorVer             uint8                   //byte
	MaxStructSize        uint16                  //word
	EntryPointRevision   uint8                   //bytes
	FormattedArea        [5]byte                 //5 bytes
	IntermediateAnrchor  IntermediateAnrchorType //5 bytes
	IntermediateChecksum uint8                   //byte
	StructTableLen       uint16                  //word
	StructTableAddress   uint32                  //dword
	NumOfStructs         uint16                  //word
	BCDRev               BCDRevType              //byte
}
type AnchorString32bit [4]byte

func (bs AnchorString32bit) String() string {
	return string(bs[:])
}

type IntermediateAnrchorType [5]byte

func (bs IntermediateAnrchorType) String() string {
	return string(bs[:])
}

type BCDRevType byte

func (bs BCDRevType) String() string {
	major := int(bs >> 4)
	minor := int(bs & 0x0F)
	return fmt.Sprintf("M:%d,m:%d", major, minor)
}

// 64 bit Entry Point struct

type EntryPointStruct_3_0 struct {
	AnchorString          AnchorString64bit
	Checksum              byte
	Length                byte
	MajorVersion          byte
	MinorVersion          byte
	Docrev                byte
	Revision              byte
	Reserved              byte
	MaxStructureSize      uint32
	StructureTableAddress uint64
}
type AnchorString64bit [5]byte

func (bs AnchorString64bit) String() string {
	return string(bs[:])
}
