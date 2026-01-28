package smbioslib

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

/*
	This was made to parse SMBIO data from MacOS ioreg AppleSMBIOS field, however apples own smbios says its running v3.3 but it follows the v2.3 EPS struct
	meaning idk what standard theyre actually following so we're just gunna make this for 2.3 and pray its correct

	I was wrong, the eps header appears to be 32bit instead of the 64bit header i expected from a 64bit system
	What the hell man

	https://www.dmtf.org/sites/default/files/standards/documents/DSP0134_3.3.0.pdf
	3.3 standard

	https://www.dmtf.org/sites/default/files/standards/documents/DSP0130.pdf
	Documentation on the 2.3 SMBIOS standard
*/

type word uint16
type dword uint32
type qword uint64

/*
	EPS Types
*/

type AnchorString [4]byte

func (bs AnchorString) String() string {
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

type EntryPointStruct struct {
	ArchorString         AnchorString            //4 bytes
	Checksum             uint8                   //byte
	EntryPointLength     uint8                   //byte
	MajorVer             uint8                   //byte
	MinorVer             uint8                   //byte
	MaxStructSize        word                    //word
	EntryPointRevision   uint8                   //bytes
	FormattedArea        [5]byte                 //5 bytes
	IntermediateAnrchor  IntermediateAnrchorType //5 bytes
	IntermediateChecksum uint8                   //byte
	StructTableLen       word                    //word
	StructTableAddress   dword                   //dword
	NumOfStructs         word                    //word
	BCDRev               BCDRevType              //byte
}

func (eps *EntryPointStruct) Check() bool {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, eps)
	data := buf.Bytes()
	var sum uint8
	for _, b := range data {
		sum += b
	}
	log.Printf("EPS Check: %d\n", sum)
	return sum == 0
}

/*
	SMBIOS Structures
*/

type ByteString byte

func (bs ByteString) String() string {
	return string(bs)
}

type StructureType byte

const (
	TypeBiosInfo                  StructureType = 0
	TypeSystemInfo                StructureType = 1
	TypeSystemEnclosure           StructureType = 3
	TypeProcessorInfo             StructureType = 4
	TypeCacheInfo                 StructureType = 7
	TypeSystemSlots               StructureType = 9
	TypePhysicalMemoryArray       StructureType = 16
	TypeMemoryDevice              StructureType = 17
	TypeMemoryArrayMappedAddress  StructureType = 19
	TypeMemoryDeviceMappedAddress StructureType = 20
)

type StructureHeader struct {
	Type   StructureType
	Length byte // depending on the struct there will be a default length assuming no extensions bytes are used
	Handle word
}

type TestStruct struct {
	StructureHeader
}

/*
Processor Information
*/

type ProcessorInfo struct {
	StructureHeader
	SocketDesignation        byte
	ProcessorType            byte //enum
	ProcessorFamily          byte //enum
	ProcessorManufacturer    byte
	ProcessorID              qword
	ProcessorVersion         byte
	Voltage                  byte
	ExternalClock            word
	MaxSpeed                 word //in MHz
	CurrentSpeed             word
	Status                   byte
	ProcessorUpgrade         byte //enum
	CacheOneHandler          word
	CacheTwoHandler          word
	CacheThreeHandler        word
	SerialNum                byte
	AssetTag                 byte
	PartNumber               byte
	CoreCount                byte
	CoreEnabled              byte
	ThreadCount              byte
	ProcessorCharacteristics word
	ProcessorFamilyTwo       word
	CoreCountTwo             byte
	CoreEnabledTwo           byte
	ThreadCountTwo           byte
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
	BiosStartAddressSegment  word
	BiosReleaseDate          ByteString
	BiosRomSize              byte
	BiosCharacteristics      qword
	CharacteristicsExtension []byte
}
