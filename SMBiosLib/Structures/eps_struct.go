package structures

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

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
	MaxStructSize        Word                    //word
	EntryPointRevision   uint8                   //bytes
	FormattedArea        [5]byte                 //5 bytes
	IntermediateAnrchor  IntermediateAnrchorType //5 bytes
	IntermediateChecksum uint8                   //byte
	StructTableLen       Word                    //word
	StructTableAddress   DWord                   //dword
	NumOfStructs         Word                    //word
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
	return sum == 0
}
