package utility

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
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
	value := reflect.ValueOf(obj).Elem()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		if reader.Len() < int(field.Type().Size()) {
			break
		}
		err := binary.Read(reader, binary.LittleEndian, field.Addr().Interface())
		if err != nil {
			return fmt.Errorf("field %d: %w", i, err)
		}
	}
	return nil
}

func PrintObj(obj any) {
	js, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		log.Println("Error unmarshaling")
	}
	fmt.Println(string(js))
}
