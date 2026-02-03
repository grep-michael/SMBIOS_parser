package structures

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

type SMBiosStruct interface {
	Parse([]byte) error
}

func As[T any](s SMBiosStruct) (*GenericStruct[T], bool) {
	v, ok := s.(*GenericStruct[T])
	return v, ok
}

type GenericStruct[T any] struct {
	Data    T
	Strings []string
}

func (g *GenericStruct[T]) Parse(data []byte) error {
	r := bytes.NewReader(data)
	err := binary.Read(r, binary.LittleEndian, &g.Data)
	if err != nil {
		return err
	}
	g.Strings = parseNullTerminatedStrings(data[data[1]:])
	return nil
}
func parseNullTerminatedStrings(data []byte) []string {
	var strings []string
	start := 0

	for i := 0; i < len(data)-1; i++ {
		if data[i] == 0x00 {
			if data[i+1] == 0x00 { //double null, end of section
				if i > start {
					strings = append(strings, string(data[start:i]))
				}
				break
			}
			if i > start {
				strings = append(strings, string(data[start:i]))
			}
			start = i + 1
		}
	}
	return strings
}

// bytes that suppose to be enums
type ByteEnum byte

// Byte that represents a string index
type ByteStringIndex byte

func (byt ByteStringIndex) String() string {
	i := int(byt)
	return strconv.Itoa(i)
}

// An actual byte that is suppose to be a string
type ByteString byte

func (byt ByteString) String() string {
	return string(byt)
}
