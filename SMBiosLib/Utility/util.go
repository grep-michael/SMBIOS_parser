package utility

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func ReadIntoStruct(data []byte, obj interface{}) error {
	reader := bytes.NewReader(data)
	if err := binary.Read(reader, binary.LittleEndian, obj); err != nil {
		return fmt.Errorf("failed to read binary data: %w", err)
	}
	return nil
}
