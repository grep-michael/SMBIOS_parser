package parsers

var WakeupTypeMap = map[byte]string{
	0x00: "Reserved",
	0x01: "Other",
	0x02: "Unknown",
	0x03: "APM Timer",
	0x04: "Modem Ring",
	0x05: "LAN Remote",
	0x06: "Power Switch",
	0x07: "PCI PME#",
	0x08: "AC Power Restored",
}
