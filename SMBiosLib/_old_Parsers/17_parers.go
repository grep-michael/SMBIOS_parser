package parsers

import "fmt"

var MemoryFormFactorMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "SIMM",
	0x04: "SIP",
	0x05: "Chip",
	0x06: "DIP",
	0x07: "ZIP",
	0x08: "Proprietary Card",
	0x09: "DIMM",
	0x0A: "TSOP",
	0x0B: "Row of chips",
	0x0C: "RIMM",
	0x0D: "SODIMM",
	0x0E: "SRIMM",
	0x0F: "FB-DIMM",
	0x10: "Die",
}

var MemoryTypeMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "DRAM",
	0x04: "EDRAM",
	0x05: "VRAM",
	0x06: "SRAM",
	0x07: "RAM",
	0x08: "ROM",
	0x09: "FLASH",
	0x0A: "EEPROM",
	0x0B: "FEPROM",
	0x0C: "EPROM",
	0x0D: "CDRAM",
	0x0E: "3DRAM",
	0x0F: "SDRAM",
	0x10: "SGRAM",
	0x11: "RDRAM",
	0x12: "DDR",
	0x13: "DDR2",
	0x14: "DDR2 FB-DIMM",
	0x18: "DDR3",
	0x19: "FBD2",
	0x1A: "DDR4",
	0x1B: "LPDDR",
	0x1C: "LPDDR2",
	0x1D: "LPDDR3",
	0x1E: "LPDDR4",
	0x1F: "Logical non-volatile device",
	0x20: "HBM (High Bandwidth Memory)",
	0x21: "HBM2 (High Bandwidth Memory Generation 2)",
}

func GetMemorySize(size uint16) string {
	list := []string{"KB", "MB", "GB"}
	index := 0
	is_megaBytes := 0x2000&0x8000 == 0
	if is_megaBytes {
		index = 1
	} else {
		index = 0
	}

	for size > 1000 {
		size = size / 1000
		index += 1
	}
	return fmt.Sprintf("%d %s", size, list[index])
}

func GetMemoryFormFactor(b byte) string {
	if value, exists := MemoryFormFactorMap[b]; exists {
		return value
	}
	return fmt.Sprintf("Unknown Form Factor (0x%02X)", b)
}

func GetMemoryType(b byte) string {
	if value, exists := MemoryTypeMap[b]; exists {
		return value
	}
	return fmt.Sprintf("Unknown Memory Type (0x%02X)", b)
}
