package processor

import (
	"fmt"

	utility "github.com/grep-michael/SMBIOS_parser/SMBiosLib/Utility"
)

type Processor struct {
	Processor ProcessorInfo
	Strings   []string
}

type ProcessorInfo interface {
	GetFamily() string
	GetManufacturer([]string) string
	GetSpeed() uint16
	GetProcessorUpgrade() string
	GetSerial([]string) string
	GetAssetTag([]string) string
	GetPartNumber([]string) string
	GetCoreCount() byte
	GetThreadCount() byte
}

func ParseChunk(data []byte) (*Processor, error) {
	if data[0] != 4 {
		return nil, fmt.Errorf("Not Processor Type, data Misaligned or wrong struct type, type is %d expected 4", data[0])
	}
	processor := &Processor{}
	switch data[1] {
	case 0x30:
		//3.0
		info := &Processor3_0Fixed{}
		err := utility.ReadIntoStruct(data, info)
		strings := utility.ParseNullTerminatedStrings(data[int(data[1]):])
		processor.Strings = strings
		processor.Processor = info
		return processor, err
	case 0x1A:
		//2.0
		fallthrough
	case 0x23:
		//2.3
		fallthrough
	case 0x28:
		//2.5
		fallthrough
	case 0x2A:
		//2.6
		fallthrough
	default:
		return nil, fmt.Errorf("Unspported Processor Struct 0x%x", data[1])
	}
}
