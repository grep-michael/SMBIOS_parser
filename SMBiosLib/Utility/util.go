package utility

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func ParseNullTerminatedStrings(data []byte) []string {
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

func ReadIntoStruct(data []byte, obj interface{}) error {
	reader := bytes.NewReader(data)
	if err := binary.Read(reader, binary.LittleEndian, obj); err != nil {
		return fmt.Errorf("failed to read binary data: %w", err)
	}
	return nil
}
