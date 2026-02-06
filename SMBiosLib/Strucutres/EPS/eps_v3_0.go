package eps

type EntryPointStruct_3_0 struct {
	AnchorString          AnchorString64bit
	Checksum              byte
	Length                byte
	MajorVersion          byte
	MinorVersion          byte
	Docrev                byte
	Revision              byte
	Reserved              byte
	MaxTableSize          uint32
	StructureTableAddress uint64
}
type AnchorString64bit [5]byte

func (bs AnchorString64bit) String() string {
	return string(bs[:])
}
