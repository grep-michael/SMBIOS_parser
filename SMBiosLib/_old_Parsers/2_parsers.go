package parsers

type BoardFeatures struct {
	Reserved         uint8
	HotSwappable     bool
	Replaceable      bool
	Removable        bool
	RequiresDaughter bool
	HostingBoard     bool
}

func ParseBoardFeatures(b byte) BoardFeatures {
	return BoardFeatures{
		Reserved:         (b >> 5) & 0x07,
		HotSwappable:     (b>>4)&0x01 == 1,
		Replaceable:      (b>>3)&0x01 == 1,
		Removable:        (b>>2)&0x01 == 1,
		RequiresDaughter: (b>>1)&0x01 == 1,
		HostingBoard:     b&0x01 == 1,
	}
}

var BoardTypeMap = map[byte]string{
	0x01: "Unknown",
	0x02: "Other",
	0x03: "Server Blade",
	0x04: "Connectivity Switch",
	0x05: "System Management Module",
	0x06: "Processor Module",
	0x07: "I/O Module",
	0x08: "Memory Module",
	0x09: "Daughter board",
	0x0A: "Motherboard (includes processor, memory, and I/O)",
	0x0B: "Processor/Memory Module",
	0x0C: "Processor/IO Module",
	0x0D: "Interconnect board",
}
