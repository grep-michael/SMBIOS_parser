package structures

type CatchAll struct {
	StructureHeader
}

// Byte that represents a string index
type ByteStringIndex byte

// An actual byte that is suppose to be a string
type ByteString byte

func (bs ByteStringIndex) String() string {
	return string(bs)
}
