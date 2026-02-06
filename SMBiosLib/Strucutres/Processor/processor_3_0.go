package processor

// Structs and Maps
type Processor3_0Fixed struct {
	Type                     byte
	Length                   byte
	Handle                   uint16
	SocketDesignation        byte //index
	ProcessorType            byte //enum
	ProcessorFamily          byte //enum
	ProcessorManufacturer    byte // index
	ProcessorID              uint64
	ProcessorVersion         byte //index
	Voltage                  byte
	ExternalClock            uint16
	MaxSpeed                 uint16 //in MHz
	CurrentSpeed             uint16
	Status                   byte
	ProcessorUpgrade         byte
	CacheOneHandler          uint16
	CacheTwoHandler          uint16
	CacheThreeHandler        uint16
	SerialNum                byte //index
	AssetTag                 byte //index
	PartNumber               byte //index
	CoreCount                byte
	CoreEnabled              byte
	ThreadCount              byte
	ProcessorCharacteristics uint16
	ProcessorFamilyTwo       uint16
	CoreCountTwo             byte
	CoreEnabledTwo           byte
	ThreadCountTwo           byte
}

var Processor3_0UpgradeMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "Daughter Board",
	0x04: "ZIF Socket",
	0x05: "Replaceable Piggy Back",
	0x06: "None",
	0x07: "LIF Socket",
	0x08: "Slot 1",
	0x09: "Slot 2",
	0x0A: "370-pin socket",
	0x0B: "Slot A",
	0x0C: "Slot M",
	0x0D: "Socket 423",
	0x0E: "Socket A (Socket 462)",
	0x0F: "Socket 478",
	0x10: "Socket 754",
	0x11: "Socket 940",
	0x12: "Socket 939",
	0x13: "Socket mPGA604",
	0x14: "Socket LGA771",
	0x15: "Socket LGA775",
	0x16: "Socket S1",
	0x17: "Socket AM2",
	0x18: "Socket F (1207)",
	0x19: "Socket LGA1366",
	0x1A: "Socket G34",
	0x1B: "Socket AM3",
	0x1C: "Socket C32",
	0x1D: "Socket LGA1156",
	0x1E: "Socket LGA1567",
	0x1F: "Socket PGA988A",
	0x20: "Socket BGA1288",
	0x21: "Socket rPGA988B",
	0x22: "Socket BGA1023",
	0x23: "Socket BGA1224",
	0x24: "Socket LGA1155",
	0x25: "Socket LGA1356",
	0x26: "Socket LGA2011",
	0x27: "Socket FS1",
	0x28: "Socket FS2",
	0x29: "Socket FM1",
	0x2A: "Socket FM2",
	0x2B: "Socket LGA2011-3",
	0x2C: "Socket LGA1356-3",
	0x2D: "Socket LGA1150",
	0x2E: "Socket BGA1168",
	0x2F: "Socket BGA1234",
	0x30: "Socket BGA1364",
	0x31: "Socket AM4",
	0x32: "Socket LGA1151",
	0x33: "Socket BGA1356",
	0x34: "Socket BGA1440",
	0x35: "Socket BGA1515",
	0x36: "Socket LGA3647-1",
	0x37: "Socket SP3",
	0x38: "Socket SP3r2",
	0x39: "Socket LGA2066",
	0x3A: "Socket BGA1392",
	0x3B: "Socket BGA1510",
	0x3C: "Socket BGA1528",
}
var Processor3_0FamilyMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "8086",
	0x04: "80286",
	0x05: "Intel386™ processor",
	0x06: "Intel486™ processor",
	0x07: "8087",
	0x08: "80287",
	0x09: "80387",
	0x0A: "80487",
	0x0B: "Intel® Pentium® processor",
	0x0C: "Pentium® Pro processor",
	0x0D: "Pentium® II processor",
	0x0E: "Pentium® processor with MMX™ technology",
	0x0F: "Intel® Celeron® processor",
	0x10: "Pentium® II Xeon™ processor",
	0x11: "Pentium® III processor",
	0x12: "M1 Family",
	0x13: "M2 Family",
	0x14: "Intel® Celeron® M processor",
	0x15: "Intel® Pentium® 4 HT processor",
	0x18: "AMD Duron™ Processor Family",
	0x19: "K5 Family",
	0x1A: "K6 Family",
	0x1B: "K6-2",
	0x1C: "K6-3",
	0x1D: "AMD Athlon™ Processor Family",
	0x1E: "AMD29000 Family",
	0x1F: "K6-2+",
	0x20: "Power PC Family",
	0x21: "Power PC 601",
	0x22: "Power PC 603",
	0x23: "Power PC 603+",
	0x24: "Power PC 604",
	0x25: "Power PC 620",
	0x26: "Power PC x704",
	0x27: "Power PC 750",
	0x28: "Intel® Core™ Duo processor",
	0x29: "Intel® Core™ Duo mobile processor",
	0x2A: "Intel® Core™ Solo mobile processor",
	0x2B: "Intel® Atom™ processor",
	0x2C: "Intel® Core™ M processor",
	0x2D: "Intel(R) Core(TM) m3 processor",
	0x2E: "Intel(R) Core(TM) m5 processor",
	0x2F: "Intel(R) Core(TM) m7 processor",
	0x30: "Alpha Family",
	0x31: "Alpha 21064",
	0x32: "Alpha 21066",
	0x33: "Alpha 21164",
	0x34: "Alpha 21164PC",
	0x35: "Alpha 21164a",
	0x36: "Alpha 21264",
	0x37: "Alpha 21364",
	0x38: "AMD Turion™ II Ultra Dual-Core Mobile M Processor Family",
	0x39: "AMD Turion™ II Dual-Core Mobile M Processor Family",
	0x3A: "AMD Athlon™ II Dual-Core M Processor Family",
	0x3B: "AMD Opteron™ 6100 Series Processor",
	0x3C: "AMD Opteron™ 4100 Series Processor",
	0x3D: "AMD Opteron™ 6200 Series Processor",
	0x3E: "AMD Opteron™ 4200 Series Processor",
	0x3F: "AMD FX™ Series Processor",
	0x40: "MIPS Family",
	0x41: "MIPS R4000",
	0x42: "MIPS R4200",
	0x43: "MIPS R4400",
	0x44: "MIPS R4600",
	0x45: "MIPS R10000",
	0x46: "AMD C-Series Processor",
	0x47: "AMD E-Series Processor",
	0x48: "AMD A-Series Processor",
	0x49: "AMD G-Series Processor",
	0x4A: "AMD Z-Series Processor",
	0x4B: "AMD R-Series Processor",
	0x4C: "AMD Opteron™ 4300 Series Processor",
	0x4D: "AMD Opteron™ 6300 Series Processor",
	0x4E: "AMD Opteron™ 3300 Series Processor",
	0x4F: "AMD FirePro™ Series Processor",
	0x50: "SPARC Family",
	0x51: "SuperSPARC",
	0x52: "microSPARC II",
	0x53: "microSPARC IIep",
	0x54: "UltraSPARC",
	0x55: "UltraSPARC II",
	0x56: "UltraSPARC Iii",
	0x57: "UltraSPARC III",
	0x58: "UltraSPARC IIIi",
	0x60: "68040 Family",
	0x61: "68xxx",
	0x62: "68000",
	0x63: "68010",
	0x64: "68020",
	0x65: "68030",
	0x66: "AMD Athlon(TM) X4 Quad-Core Processor Family",
	0x67: "AMD Opteron(TM) X1000 Series Processor",
	0x68: "AMD Opteron(TM) X2000 Series APU",
	0x69: "AMD Opteron(TM) A-Series Processor",
	0x6A: "AMD Opteron(TM) X3000 Series APU",
	0x6B: "AMD Zen Processor Family",
	0x70: "Hobbit Family",
	0x78: "Crusoe™ TM5000 Family",
	0x79: "Crusoe™ TM3000 Family",
	0x7A: "Efficeon™ TM8000 Family",
	0x80: "Weitek",
	0x82: "Itanium™ processor",
	0x83: "AMD Athlon™ 64 Processor Family",
	0x84: "AMD Opteron™ Processor Family",
	0x85: "AMD Sempron™ Processor Family",
	0x86: "AMD Turion™ 64 Mobile Technology",
	0x87: "Dual-Core AMD Opteron™ Processor Family",
	0x88: "AMD Athlon™ 64 X2 Dual-Core Processor Family",
	0x89: "AMD Turion™ 64 X2 Mobile Technology",
	0x8A: "Quad-Core AMD Opteron™ Processor Family",
	0x8B: "Third-Generation AMD Opteron™ Processor Family",
	0x8C: "AMD Phenom™ FX Quad-Core Processor Family",
	0x8D: "AMD Phenom™ X4 Quad-Core Processor Family",
	0x8E: "AMD Phenom™ X2 Dual-Core Processor Family",
	0x8F: "AMD Athlon™ X2 Dual-Core Processor Family",
	0x90: "PA-RISC Family",
	0x91: "PA-RISC 8500",
	0x92: "PA-RISC 8000",
	0x93: "PA-RISC 7300LC",
	0x94: "PA-RISC 7200",
	0x95: "PA-RISC 7100LC",
	0x96: "PA-RISC 7100",
	0xA0: "V30 Family",
	0xA1: "Quad-Core Intel® Xeon® processor 3200 Series",
	0xA2: "Dual-Core Intel® Xeon® processor 3000 Series",
	0xA3: "Quad-Core Intel® Xeon® processor 5300 Series",
	0xA4: "Dual-Core Intel® Xeon® processor 5100 Series",
	0xA5: "Dual-Core Intel® Xeon® processor 5000 Series",
	0xA6: "Dual-Core Intel® Xeon® processor LV",
	0xA7: "Dual-Core Intel® Xeon® processor ULV",
	0xA8: "Dual-Core Intel® Xeon® processor 7100 Series",
	0xA9: "Quad-Core Intel® Xeon® processor 5400 Series",
	0xAA: "Quad-Core Intel® Xeon® processor",
	0xAB: "Dual-Core Intel® Xeon® processor 5200 Series",
	0xAC: "Dual-Core Intel® Xeon® processor 7200 Series",
	0xAD: "Quad-Core Intel® Xeon® processor 7300 Series",
	0xAE: "Quad-Core Intel® Xeon® processor 7400 Series",
	0xAF: "Multi-Core Intel® Xeon® processor 7400 Series",
	0xB0: "Pentium® III Xeon™ processor",
	0xB1: "Pentium® III Processor with Intel® SpeedStep™ Technology",
	0xB2: "Pentium® 4 Processor",
	0xB3: "Intel® Xeon® processor",
	0xB4: "AS400 Family",
	0xB5: "Intel® Xeon™ processor MP",
	0xB6: "AMD Athlon™ XP Processor Family",
	0xB7: "AMD Athlon™ MP Processor Family",
	0xB8: "Intel® Itanium® 2 processor",
	0xB9: "Intel® Pentium® M processor",
	0xBA: "Intel® Celeron® D processor",
	0xBB: "Intel® Pentium® D processor",
	0xBC: "Intel® Pentium® Processor Extreme Edition",
	0xBD: "Intel® Core™ Solo Processor",
	0xBE: "Reserved",
	0xBF: "Intel® Core™ 2 Duo Processor",
	0xC0: "Intel® Core™ 2 Solo processor",
	0xC1: "Intel® Core™ 2 Extreme processor",
	0xC2: "Intel® Core™ 2 Quad processor",
	0xC3: "Intel® Core™ 2 Extreme mobile processor",
	0xC4: "Intel® Core™ 2 Duo mobile processor",
	0xC5: "Intel® Core™ 2 Solo mobile processor",
	0xC6: "Intel® Core™ i7 processor",
	0xC7: "Dual-Core Intel® Celeron® processor",
	0xC8: "IBM390 Family",
	0xC9: "G4",
	0xCA: "G5",
	0xCB: "ESA/390 G6",
	0xCC: "z/Architecture base",
	0xCD: "Intel® Core™ i5 processor",
	0xCE: "Intel® Core™ i3 processor",
	0xCF: "Intel® Core™ i9 processor",
	0xD2: "VIA C7™-M Processor Family",
	0xD3: "VIA C7™-D Processor Family",
	0xD4: "VIA C7™ Processor Family",
	0xD5: "VIA Eden™ Processor Family",
	0xD6: "Multi-Core Intel® Xeon® processor",
	0xD7: "Dual-Core Intel® Xeon® processor 3xxx Series",
	0xD8: "Quad-Core Intel® Xeon® processor 3xxx Series",
	0xD9: "VIA Nano™ Processor Family",
	0xDA: "Dual-Core Intel® Xeon® processor 5xxx Series",
}
var Processor3_0TypeDict = map[byte]string{
	0x01: "Other",
	0x02: "Unknow",
	0x03: "Central",
	0x04: "Math",
	0x05: "DSP",
	0x06: "Video",
}

type Processor3_0Characteristics struct {
	Reserved0               bool
	Unknown                 bool
	Capable64Bit            bool
	MultiCore               bool
	HardwareThread          bool
	ExecuteProtection       bool
	EnhancedVirtualization  bool
	PowerPerformanceControl bool
	Capable128Bit           bool
	Reserved9to15           uint8
}

func ParseProcessorCharacteristics(characteristics uint16) Processor3_0Characteristics {
	return Processor3_0Characteristics{
		Reserved0:               (characteristics>>0)&0x01 == 1,
		Unknown:                 (characteristics>>1)&0x01 == 1,
		Capable64Bit:            (characteristics>>2)&0x01 == 1,
		MultiCore:               (characteristics>>3)&0x01 == 1,
		HardwareThread:          (characteristics>>4)&0x01 == 1,
		ExecuteProtection:       (characteristics>>5)&0x01 == 1,
		EnhancedVirtualization:  (characteristics>>6)&0x01 == 1,
		PowerPerformanceControl: (characteristics>>7)&0x01 == 1,
		Capable128Bit:           (characteristics>>8)&0x01 == 1,
		Reserved9to15:           uint8((characteristics >> 9) & 0x7F), // Extract bits 9:15 (7 bits)
	}
}

// Interface functions
func (p *Processor3_0Fixed) GetFamily() string {
	if v, exists := Processor3_0FamilyMap[p.ProcessorFamily]; exists {
		return v
	}
	return "Unknown"
}

func (p *Processor3_0Fixed) GetManufacturer(strings []string) string {
	return strings[int(p.ProcessorManufacturer)]
}

func (p *Processor3_0Fixed) GetSpeed() uint16 {
	return p.CurrentSpeed
}

func (p *Processor3_0Fixed) GetProcessorUpgrade() string {
	if upgrade, exists := Processor3_0UpgradeMap[p.ProcessorUpgrade]; exists {
		return upgrade
	}
	return "Unknown"
}

func (p *Processor3_0Fixed) GetSerial(strings []string) string {
	return strings[int(p.SerialNum)]
}

func (p *Processor3_0Fixed) GetAssetTag(strings []string) string {
	return strings[int(p.AssetTag)]
}

func (p *Processor3_0Fixed) GetPartNumber(strings []string) string {

	return strings[int(p.PartNumber)]
}
func (p *Processor3_0Fixed) GetCoreCount() byte {
	if p.CoreCountTwo != 0 {
		return p.CoreCountTwo
	}
	return p.CoreCount
}
func (p *Processor3_0Fixed) GetThreadCount() byte {
	if p.ThreadCountTwo != 0 {
		return p.ThreadCountTwo
	}
	return p.ThreadCount
}
