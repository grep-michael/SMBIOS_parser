package smbios

import (
	"fmt"
)


// Platform Firmware Information (Type 0) structure (Type 0)
type Type0PlatformFirmwareInformation struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Vendor byte //STRING
	FirmwareVersion byte //STRING
	BIOSStartingAddressSegment uint16 //
	FirmwareReleaseDate byte //STRING
	FirmwareROMSize byte //
	FirmwareCharacteristics uint64 //Bit Field
	FirmwareCharacteristicsExtensionBytes [2]byte //Bit Field
	PlatformFirmwareMajorRelease byte //
	PlatformFirmwareMinorRelease byte //
	EmbeddedControllerFirmwareMajorRelease byte //
	EmbeddedControllerFirmwareMinorRelease byte //
	ExtendedFirmwareROMSize uint16 //Bit Field
}
func (s *Type0PlatformFirmwareInformation) GetVendor(strings []string) string {
	idx := int(s.Vendor)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type0PlatformFirmwareInformation) GetFirmwareVersion(strings []string) string {
	idx := int(s.FirmwareVersion)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type0PlatformFirmwareInformation) GetFirmwareReleaseDate(strings []string) string {
	idx := int(s.FirmwareReleaseDate)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}
// Firmware Characteristics
const (
	Type0_FirmwareCharacteristics_Bit0 = 1 << 0 // Reserved.
	Type0_FirmwareCharacteristics_Bit1 = 1 << 1 // Reserved.
	Type0_FirmwareCharacteristics_Bit2 = 1 << 2 // Unknown.
	Type0_FirmwareCharacteristics_Bit3 = 1 << 3 // Firmware Characteristics are not supported.
	Type0_FirmwareCharacteristics_Bit4 = 1 << 4 // ISA is supported.
	Type0_FirmwareCharacteristics_Bit5 = 1 << 5 // MCA is supported.
	Type0_FirmwareCharacteristics_Bit6 = 1 << 6 // EISA is supported.
	Type0_FirmwareCharacteristics_Bit7 = 1 << 7 // PCI is supported.
	Type0_FirmwareCharacteristics_Bit8 = 1 << 8 // PC card (PCMCIA) is supported.
	Type0_FirmwareCharacteristics_Bit9 = 1 << 9 // Plug and Play is supported.
	Type0_FirmwareCharacteristics_Bit10 = 1 << 10 // APM is supported.
	Type0_FirmwareCharacteristics_Bit11 = 1 << 11 // Firmware is upgradeable (Flash).
	Type0_FirmwareCharacteristics_Bit12 = 1 << 12 // Firmware shadowing is allowed.
	Type0_FirmwareCharacteristics_Bit13 = 1 << 13 // VL-VESA is supported.
	Type0_FirmwareCharacteristics_Bit14 = 1 << 14 // ESCD support is available.
	Type0_FirmwareCharacteristics_Bit15 = 1 << 15 // Boot from CD is supported.
	Type0_FirmwareCharacteristics_Bit16 = 1 << 16 // Selectable boot is supported.
	Type0_FirmwareCharacteristics_Bit17 = 1 << 17 // Firmware ROM is socketed (e.g., PLCC or SOP socket).
	Type0_FirmwareCharacteristics_Bit18 = 1 << 18 // Boot from PC card (PCMCIA) is supported.
	Type0_FirmwareCharacteristics_Bit19 = 1 << 19 // EDD specification is supported.
	Type0_FirmwareCharacteristics_Bit20 = 1 << 20 // Int 13h — Japanese floppy for NEC 9800 1.2 MB (3.5", 1K bytes/sector, 360 RPM) i
	Type0_FirmwareCharacteristics_Bit21 = 1 << 21 // Int 13h — Japanese floppy for Toshiba 1.2 MB (3.5", 360 RPM) is supported.
	Type0_FirmwareCharacteristics_Bit22 = 1 << 22 // Int 13h — 5.25" / 360 KB floppy services are supported.
	Type0_FirmwareCharacteristics_Bit23 = 1 << 23 // Int 13h — 5.25" /1.2 MB floppy services are supported.
	Type0_FirmwareCharacteristics_Bit24 = 1 << 24 // Int 13h — 3.5" / 720 KB floppy services are supported.
	Type0_FirmwareCharacteristics_Bit25 = 1 << 25 // Int 13h — 3.5" / 2.88 MB floppy services are supported.
	Type0_FirmwareCharacteristics_Bit26 = 1 << 26 // Int 5h, print screen Service is supported.
	Type0_FirmwareCharacteristics_Bit27 = 1 << 27 // Int 9h, 8042 keyboard services are supported.
	Type0_FirmwareCharacteristics_Bit28 = 1 << 28 // Int 14h, serial services are supported.
	Type0_FirmwareCharacteristics_Bit29 = 1 << 29 // Int 17h, printer services are supported.
	Type0_FirmwareCharacteristics_Bit30 = 1 << 30 // Int 10h, CGA/Mono Video Services are supported.
	Type0_FirmwareCharacteristics_Bit31 = 1 << 31 // NEC PC-98.
)

// Firmware Characteristics Extension Byte 1
const (
	Type0_FirmwareCharacteristicsExtensionByte1_Bit0 = 1 << 0 // ACPI is supported.
	Type0_FirmwareCharacteristicsExtensionByte1_Bit1 = 1 << 1 // USB Legacy is supported.
	Type0_FirmwareCharacteristicsExtensionByte1_Bit2 = 1 << 2 // AGP is supported.
	Type0_FirmwareCharacteristicsExtensionByte1_Bit3 = 1 << 3 // I2O boot is supported.
	Type0_FirmwareCharacteristicsExtensionByte1_Bit4 = 1 << 4 // LS-120 SuperDisk boot is supported.
	Type0_FirmwareCharacteristicsExtensionByte1_Bit5 = 1 << 5 // ATAPI ZIP drive boot is supported.
	Type0_FirmwareCharacteristicsExtensionByte1_Bit6 = 1 << 6 // 1394 boot is supported.
	Type0_FirmwareCharacteristicsExtensionByte1_Bit7 = 1 << 7 // Smart battery is supported.
)

// Firmware Characteristics Extension Byte 2
const (
	Type0_FirmwareCharacteristicsExtensionByte2_Bit0 = 1 << 0 // BIOS Boot Specification is supported.
	Type0_FirmwareCharacteristicsExtensionByte2_Bit1 = 1 << 1 // Function key-initiated network service boot is supported. When function key-unin
	Type0_FirmwareCharacteristicsExtensionByte2_Bit2 = 1 << 2 // Enable targeted content distribution. The manufacturer has ensured that the SMBI
	Type0_FirmwareCharacteristicsExtensionByte2_Bit3 = 1 << 3 // UEFI Specification is supported.
	Type0_FirmwareCharacteristicsExtensionByte2_Bit4 = 1 << 4 // SMBIOS table describes a virtual machine. (If this bit is not set, no inference 
	Type0_FirmwareCharacteristicsExtensionByte2_Bit5 = 1 << 5 // Manufacturing mode is supported. (Manufacturing mode is a special boot mode, not
	Type0_FirmwareCharacteristicsExtensionByte2_Bit6 = 1 << 6 // Manufacturing mode is enabled.
	Type0_FirmwareCharacteristicsExtensionByte2_Bit7 = 1 << 7 // Reserved for future assignment by this specification.
)
func (s *Type0PlatformFirmwareInformation) GetFirmwareCharacteristicsFlags() []string {
	v := s.FirmwareCharacteristics
	var flags []string
	if v&(1<<0) != 0 { flags = append(flags, "Reserved") }
	if v&(1<<1) != 0 { flags = append(flags, "Reserved") }
	if v&(1<<2) != 0 { flags = append(flags, "Unknown") }
	if v&(1<<3) != 0 { flags = append(flags, "Firmware Characteristics are not supported") }
	if v&(1<<4) != 0 { flags = append(flags, "ISA is supported") }
	if v&(1<<5) != 0 { flags = append(flags, "MCA is supported") }
	if v&(1<<6) != 0 { flags = append(flags, "EISA is supported") }
	if v&(1<<7) != 0 { flags = append(flags, "PCI is supported") }
	if v&(1<<8) != 0 { flags = append(flags, "PC card (PCMCIA) is supported") }
	if v&(1<<9) != 0 { flags = append(flags, "Plug and Play is supported") }
	if v&(1<<10) != 0 { flags = append(flags, "APM is supported") }
	if v&(1<<11) != 0 { flags = append(flags, "Firmware is upgradeable (Flash)") }
	if v&(1<<12) != 0 { flags = append(flags, "Firmware shadowing is allowed") }
	if v&(1<<13) != 0 { flags = append(flags, "VL-VESA is supported") }
	if v&(1<<14) != 0 { flags = append(flags, "ESCD support is available") }
	if v&(1<<15) != 0 { flags = append(flags, "Boot from CD is supported") }
	if v&(1<<16) != 0 { flags = append(flags, "Selectable boot is supported") }
	if v&(1<<17) != 0 { flags = append(flags, "Firmware ROM is socketed (e") }
	if v&(1<<18) != 0 { flags = append(flags, "Boot from PC card (PCMCIA) is supported") }
	if v&(1<<19) != 0 { flags = append(flags, "EDD specification is supported") }
	if v&(1<<20) != 0 { flags = append(flags, "Int 13h — Japanese floppy for NEC 9800 1") }
	if v&(1<<21) != 0 { flags = append(flags, "Int 13h — Japanese floppy for Toshiba 1") }
	if v&(1<<22) != 0 { flags = append(flags, "Int 13h — 5") }
	if v&(1<<23) != 0 { flags = append(flags, "Int 13h — 5") }
	if v&(1<<24) != 0 { flags = append(flags, "Int 13h — 3") }
	if v&(1<<25) != 0 { flags = append(flags, "Int 13h — 3") }
	if v&(1<<26) != 0 { flags = append(flags, "Int 5h, print screen Service is supported") }
	if v&(1<<27) != 0 { flags = append(flags, "Int 9h, 8042 keyboard services are supported") }
	if v&(1<<28) != 0 { flags = append(flags, "Int 14h, serial services are supported") }
	if v&(1<<29) != 0 { flags = append(flags, "Int 17h, printer services are supported") }
	if v&(1<<30) != 0 { flags = append(flags, "Int 10h, CGA/Mono Video Services are supported") }
	if v&(1<<31) != 0 { flags = append(flags, "NEC PC-98") }
	return flags
}

func (s *Type0PlatformFirmwareInformation) GetFirmwareCharacteristicsExtensionBytesFlags() []string {
	v := byte(s.FirmwareCharacteristicsExtensionBytes)
	var flags []string
	if v&(1<<0) != 0 { flags = append(flags, "ACPI is supported") }
	if v&(1<<1) != 0 { flags = append(flags, "USB Legacy is supported") }
	if v&(1<<2) != 0 { flags = append(flags, "AGP is supported") }
	if v&(1<<3) != 0 { flags = append(flags, "I2O boot is supported") }
	if v&(1<<4) != 0 { flags = append(flags, "LS-120 SuperDisk boot is supported") }
	if v&(1<<5) != 0 { flags = append(flags, "ATAPI ZIP drive boot is supported") }
	if v&(1<<6) != 0 { flags = append(flags, "1394 boot is supported") }
	if v&(1<<7) != 0 { flags = append(flags, "Smart battery is supported") }
	return flags
}

func (s *Type0PlatformFirmwareInformation) GetExtendedFirmwareROMSizeFlags() []string {
	v := uint64(s.ExtendedFirmwareROMSize)
	var flags []string
	if v&(1<<0) != 0 { flags = append(flags, "Reserved") }
	if v&(1<<1) != 0 { flags = append(flags, "Reserved") }
	if v&(1<<2) != 0 { flags = append(flags, "Unknown") }
	if v&(1<<3) != 0 { flags = append(flags, "Firmware Characteristics are not supported") }
	if v&(1<<4) != 0 { flags = append(flags, "ISA is supported") }
	if v&(1<<5) != 0 { flags = append(flags, "MCA is supported") }
	if v&(1<<6) != 0 { flags = append(flags, "EISA is supported") }
	if v&(1<<7) != 0 { flags = append(flags, "PCI is supported") }
	if v&(1<<8) != 0 { flags = append(flags, "PC card (PCMCIA) is supported") }
	if v&(1<<9) != 0 { flags = append(flags, "Plug and Play is supported") }
	if v&(1<<10) != 0 { flags = append(flags, "APM is supported") }
	if v&(1<<11) != 0 { flags = append(flags, "Firmware is upgradeable (Flash)") }
	if v&(1<<12) != 0 { flags = append(flags, "Firmware shadowing is allowed") }
	if v&(1<<13) != 0 { flags = append(flags, "VL-VESA is supported") }
	if v&(1<<14) != 0 { flags = append(flags, "ESCD support is available") }
	if v&(1<<15) != 0 { flags = append(flags, "Boot from CD is supported") }
	if v&(1<<16) != 0 { flags = append(flags, "Selectable boot is supported") }
	if v&(1<<17) != 0 { flags = append(flags, "Firmware ROM is socketed (e") }
	if v&(1<<18) != 0 { flags = append(flags, "Boot from PC card (PCMCIA) is supported") }
	if v&(1<<19) != 0 { flags = append(flags, "EDD specification is supported") }
	if v&(1<<20) != 0 { flags = append(flags, "Int 13h — Japanese floppy for NEC 9800 1") }
	if v&(1<<21) != 0 { flags = append(flags, "Int 13h — Japanese floppy for Toshiba 1") }
	if v&(1<<22) != 0 { flags = append(flags, "Int 13h — 5") }
	if v&(1<<23) != 0 { flags = append(flags, "Int 13h — 5") }
	if v&(1<<24) != 0 { flags = append(flags, "Int 13h — 3") }
	if v&(1<<25) != 0 { flags = append(flags, "Int 13h — 3") }
	if v&(1<<26) != 0 { flags = append(flags, "Int 5h, print screen Service is supported") }
	if v&(1<<27) != 0 { flags = append(flags, "Int 9h, 8042 keyboard services are supported") }
	if v&(1<<28) != 0 { flags = append(flags, "Int 14h, serial services are supported") }
	if v&(1<<29) != 0 { flags = append(flags, "Int 17h, printer services are supported") }
	if v&(1<<30) != 0 { flags = append(flags, "Int 10h, CGA/Mono Video Services are supported") }
	if v&(1<<31) != 0 { flags = append(flags, "NEC PC-98") }
	return flags
}

// System Information (Type 1) (Type 1)
type Type1SystemInformation struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Manufacturer byte //STRING
	ProductName byte //STRING
	Version byte //STRING
	SerialNumber byte //STRING
	UUID [16]byte //
	WakeupType byte //ENUM
	SKUNumber byte //STRING
	Family byte //STRING
}
func (s *Type1SystemInformation) GetManufacturer(strings []string) string {
	idx := int(s.Manufacturer)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type1SystemInformation) GetProductName(strings []string) string {
	idx := int(s.ProductName)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type1SystemInformation) GetVersion(strings []string) string {
	idx := int(s.Version)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type1SystemInformation) GetSerialNumber(strings []string) string {
	idx := int(s.SerialNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type1SystemInformation) GetSKUNumber(strings []string) string {
	idx := int(s.SKUNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type1SystemInformation) GetFamily(strings []string) string {
	idx := int(s.Family)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}
// System: Wake-up Type field
var Type1_WakeupTypeMap = map[byte]string{
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
func (s *Type1SystemInformation) GetWakeupTypeString() string {
	if v, ok := Type1_WakeupTypeMap[s.WakeupType]; ok {
		return v
	}
	return "Unknown"
}

// Baseboard (or Module) Information (Type 2) (Type 2)
type Type2BaseboardInformation struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Manufacturer byte //STRING
	Product byte //STRING
	Version byte //STRING
	SerialNumber byte //STRING
	AssetTag byte //STRING
	FeatureFlags byte //Bit Field
	LocationinChassis byte //STRING
	ChassisHandle uint16 //
	BoardType byte //ENUM
	NumberofContainedObjectHandlesn byte //
	ContainedObjectHandles byte // Type:n WORDs
}
func (s *Type2BaseboardInformation) GetManufacturer(strings []string) string {
	idx := int(s.Manufacturer)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type2BaseboardInformation) GetProduct(strings []string) string {
	idx := int(s.Product)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type2BaseboardInformation) GetVersion(strings []string) string {
	idx := int(s.Version)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type2BaseboardInformation) GetSerialNumber(strings []string) string {
	idx := int(s.SerialNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type2BaseboardInformation) GetAssetTag(strings []string) string {
	idx := int(s.AssetTag)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type2BaseboardInformation) GetLocationinChassis(strings []string) string {
	idx := int(s.LocationinChassis)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}
// Baseboard: Board Type
var Type2_BoardTypeMap = map[byte]string{
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
func (s *Type2BaseboardInformation) GetBoardTypeString() string {
	if v, ok := Type2_BoardTypeMap[s.BoardType]; ok {
		return v
	}
	return "Unknown"
}

// System Enclosure or Chassis (Type 3) (Type 3)
type Type3SystemEnclosure struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Manufacturer byte //STRING
	Type_1 byte //
	Version byte //STRING
	SerialNumber byte //STRING
	AssetTagNumber byte //STRING
	BootupState byte //ENUM
	PowerSupplyState byte //ENUM
	ThermalState byte //ENUM
	SecurityStatus byte //ENUM
	OEMdefined uint32 //
	Height byte //
	NumberofPowerCords byte //
	ContainedElementCountn byte //
	ContainedElementRecordLengthm byte //
	ContainedElements [3]byte //
	SKUNumber byte //STRING
	RackType byte //
	RackHeight byte //
}
func (s *Type3SystemEnclosure) GetManufacturer(strings []string) string {
	idx := int(s.Manufacturer)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type3SystemEnclosure) GetVersion(strings []string) string {
	idx := int(s.Version)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type3SystemEnclosure) GetSerialNumber(strings []string) string {
	idx := int(s.SerialNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type3SystemEnclosure) GetAssetTagNumber(strings []string) string {
	idx := int(s.AssetTagNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type3SystemEnclosure) GetSKUNumber(strings []string) string {
	idx := int(s.SKUNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}
// System Enclosure or Chassis Types
var Type3_SystemEnclosureorChassisTypesMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "Desktop",
	0x04: "Low Profile Desktop",
	0x05: "Pizza Box",
	0x06: "Mini Tower",
	0x07: "Tower",
	0x08: "Portable",
	0x09: "Laptop",
	0x0A: "Notebook",
	0x0B: "Hand Held",
	0x0C: "Docking Station",
	0x0D: "All in One",
	0x0E: "Sub Notebook",
	0x0F: "Space-saving",
	0x10: "Lunch Box",
	0x11: "Main Server Chassis",
	0x12: "Expansion Chassis",
	0x13: "SubChassis",
	0x14: "Bus Expansion Chassis",
	0x15: "Peripheral Chassis",
	0x16: "RAID Chassis",
	0x17: "Rack Mount Chassis",
	0x18: "Sealed-case PC",
	0x19: "Multi-system chassis",
	0x1A: "Compact PCI",
	0x1B: "Advanced TCA",
	0x1C: "Blade",
	0x1D: "Blade Enclosure",
	0x1E: "Tablet",
	0x1F: "Convertible",
	0x20: "Detachable",
	0x21: "IoT Gateway",
	0x22: "Embedded PC",
	0x23: "Mini PC",
	0x24: "Stick PC",
}

// System Enclosure or Chassis States
var Type3_SystemEnclosureorChassisStatesMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "Safe",
	0x04: "Warning",
	0x05: "Critical",
	0x06: "Non-recoverable",
}

// System Enclosure or Chassis Security Status field
var Type3_SystemEnclosureorChassisSecurityStatusMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "None",
	0x04: "External interface locked out",
	0x05: "External interface enabled",
}
func (s *Type3SystemEnclosure) GetSecurityStatusString() string {
	if v, ok := Type3_SystemEnclosureorChassisSecurityStatusMap[s.SecurityStatus]; ok {
		return v
	}
	return "Unknown"
}

// Processor Information (Type 4) (Type 4)
type Type4ProcessorInformation struct {
	Type byte //
	Length byte //
	Handle uint16 //
	SocketDesignation byte //STRING
	ProcessorType byte //ENUM
	ProcessorFamily byte //ENUM
	ProcessorManufacturer byte //STRING
	ProcessorID uint64 //
	ProcessorVersion byte //STRING
	Voltage byte //
	ExternalClock uint16 //
	MaxSpeed uint16 //
	CurrentSpeed uint16 //
	Status byte //
	ProcessorUpgrade byte //ENUM
	L1CacheHandle uint16 //
	L2CacheHandle uint16 //
	L3CacheHandle uint16 //
	SerialNumber byte //STRING
	AssetTag byte //STRING
	PartNumber byte //STRING
	CoreCount byte //
	CoreEnabled byte //
	ThreadCount byte //
	ProcessorCharacteristics uint16 //Bit Field
	ProcessorFamily2 uint16 //
	CoreCount2 uint16 //
	CoreEnabled2 uint16 //
	ThreadCount2 uint16 //
	ThreadEnabled uint16 //
	SocketType byte //STRING
}
func (s *Type4ProcessorInformation) GetSocketDesignation(strings []string) string {
	idx := int(s.SocketDesignation)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type4ProcessorInformation) GetProcessorManufacturer(strings []string) string {
	idx := int(s.ProcessorManufacturer)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type4ProcessorInformation) GetProcessorVersion(strings []string) string {
	idx := int(s.ProcessorVersion)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type4ProcessorInformation) GetSerialNumber(strings []string) string {
	idx := int(s.SerialNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type4ProcessorInformation) GetAssetTag(strings []string) string {
	idx := int(s.AssetTag)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type4ProcessorInformation) GetPartNumber(strings []string) string {
	idx := int(s.PartNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type4ProcessorInformation) GetSocketType(strings []string) string {
	idx := int(s.SocketType)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}
// Processor Information: Processor Type field
var Type4_ProcessorTypeMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "Central Processor",
	0x04: "Math Processor",
	0x05: "DSP Processor",
	0x06: "Video Processor",
}

// Processor Information: Processor Family field
var Type4_ProcessorFamilyMap = map[byte]string{
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
	0x10: "Pentium® II Xeon® processor",
	0x11: "Pentium® III processor",
	0x12: "M1 Family",
	0x13: "M2 Family",
	0x14: "Intel® Celeron® M processor",
	0x15: "Intel® Pentium® 4 HT processor",
	0x16: "Intel® Processor",
	0x17: "Available for assignment",
	0x18: "AMD Duron™ Processor Family [1]",
	0x19: "K5 Family [1]",
	0x1A: "K6 Family [1]",
	0x1B: "K6-2 [1]",
	0x1C: "K6-3 [1]",
	0x1D: "AMD Athlon™ Processor Family [1]",
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
	0x2D: "Intel® Core™ m3 processor",
	0x2E: "Intel® Core™ m5 processor",
	0x2F: "Intel® Core™ m7 processor",
	0x30: "Alpha Family [2]",
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
	0x53: "microSPARC IIIp",
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
	0x81: "Available for assignment",
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
	0xB0: "Pentium® III Xeon® processor",
	0xB1: "Pentium® III Processor with Intel® SpeedStep™ Technology",
	0xB2: "Pentium® 4 Processor",
	0xB3: "Intel® Xeon® processor",
	0xB4: "AS400 Family",
	0xB5: "Intel® Xeon® processor MP",
	0xB6: "AMD Athlon™ XP Processor Family",
	0xB7: "AMD Athlon™ MP Processor Family",
	0xB8: "Intel® Itanium® 2 processor",
	0xB9: "Intel® Pentium® M processor",
	0xBA: "Intel® Celeron® D processor",
	0xBB: "Intel® Pentium® D processor",
	0xBC: "Intel® Pentium® Processor Extreme Edition",
	0xBD: "Intel® Core™ Solo Processor",
	0xBE: "Reserved [3]",
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
	0xD0: "Intel® Xeon® D Processor family",
	0xD1: "Available for assignment",
	0xD2: "VIA C7™-M Processor Family",
	0xD3: "VIA C7™-D Processor Family",
	0xD4: "VIA C7™ Processor Family",
	0xD5: "VIA Eden™ Processor Family",
	0xD6: "Multi-Core Intel® Xeon® processor",
	0xD7: "Dual-Core Intel® Xeon® processor 3xxx Series",
	0xD8: "Quad-Core Intel® Xeon® processor 3xxx Series",
	0xD9: "VIA Nano™ Processor Family",
	0xDA: "Dual-Core Intel® Xeon® processor 5xxx Series",
	0xDB: "Quad-Core Intel® Xeon® processor 5xxx Series",
	0xDC: "Available for assignment",
	0xDD: "Dual-Core Intel® Xeon® processor 7xxx Series",
	0xDE: "Quad-Core Intel® Xeon® processor 7xxx Series",
	0xDF: "Multi-Core Intel® Xeon® processor 7xxx Series",
	0xE0: "Multi-Core Intel® Xeon® processor 3400 Series",
	0xE4: "AMD Opteron™ 3000 Series Processor",
	0xE5: "AMD Sempron™ II Processor",
	0xE6: "Embedded AMD Opteron™ Quad-Core Processor Family",
	0xE7: "AMD Phenom™ Triple-Core Processor Family",
	0xE8: "AMD Turion™ Ultra Dual-Core Mobile Processor Family",
	0xE9: "AMD Turion™ Dual-Core Mobile Processor Family",
	0xEA: "AMD Athlon™ Dual-Core Processor Family",
	0xEB: "AMD Sempron™ SI Processor Family",
	0xEC: "AMD Phenom™ II Processor Family",
	0xED: "AMD Athlon™ II Processor Family",
	0xEE: "Six-Core AMD Opteron™ Processor Family",
	0xEF: "AMD Sempron™ M Processor Family",
	0xFA: "i860",
	0xFB: "i960",
	0xFE: "Indicator to obtain the processor family from the Processor Family 2 field",
	0xFF: "Reserved",
	0x100: "ARMv7",
	0x101: "ARMv8",
	0x102: "ARMv9",
	0x103: "Reserved for future use by ARM",
	0x104: "SH-3",
	0x105: "SH-4",
	0x118: "ARM",
	0x119: "StrongARM",
	0x12C: "6x86",
	0x12D: "MediaGX",
	0x12E: "MII",
	0x140: "WinChip",
	0x15E: "DSP",
	0x1F4: "Video Processor",
	0x200: "RISC-V RV32",
	0x201: "RISC-V RV64",
	0x202: "RISC-V RV128",
	0x258: "LoongArch",
	0x259: "Loongson™ 1 Processor Family",
	0x25A: "Loongson™ 2 Processor Family",
	0x25B: "Loongson™ 3 Processor Family",
	0x25C: "Loongson™ 2K Processor Family",
	0x25D: "Loongson™ 3A Processor Family",
	0x25E: "Loongson™ 3B Processor Family",
	0x25F: "Loongson™ 3C Processor Family",
	0x260: "Loongson™ 3D Processor Family",
	0x261: "Loongson™ 3E Processor Family",
	0x262: "Dual-Core Loongson™ 2K Processor 2xxx Series",
	0x26C: "Quad-Core Loongson™ 3A Processor 5xxx Series",
	0x26D: "Multi-Core Loongson™ 3A Processor 5xxx Series",
	0x26E: "Quad-Core Loongson™ 3B Processor 5xxx Series",
	0x26F: "Multi-Core Loongson™ 3B Processor 5xxx Series",
	0x270: "Multi-Core Loongson™ 3C Processor 5xxx Series",
	0x271: "Multi-Core Loongson™ 3D Processor 5xxx Series",
	0x300: "Intel® Core™ 3",
	0x301: "Intel® Core™ 5",
	0x302: "Intel® Core™ 7",
	0x303: "Intel® Core™ 9",
	0x304: "Intel® Core™ Ultra 3",
	0x305: "Intel® Core™ Ultra 5",
	0x306: "Intel® Core™ Ultra 7",
	0x307: "Intel® Core™ Ultra 9",
}

// Processor Information: Processor Upgrade field
var Type4_ProcessorUpgradeMap = map[byte]string{
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
	0x3D: "Socket LGA4189",
	0x3E: "Socket LGA1200",
	0x3F: "Socket LGA4677",
	0x40: "Socket LGA1700",
	0x41: "Socket BGA1744",
	0x42: "Socket BGA1781",
	0x43: "Socket BGA1211",
	0x44: "Socket BGA2422",
	0x45: "Socket LGA1211",
	0x46: "Socket LGA2422",
	0x47: "Socket LGA5773",
	0x48: "Socket BGA5773",
	0x49: "Socket AM5",
	0x4A: "Socket SP5",
	0x4B: "Socket SP6",
	0x4C: "Socket BGA883",
	0x4D: "Socket BGA1190",
	0x4E: "Socket BGA4129",
	0x4F: "Socket LGA4710",
	0x50: "Socket LGA7529",
	0x51: "Socket BGA1964",
	0x52: "Socket BGA1792",
	0x53: "Socket BGA2049",
	0x54: "Socket BGA2551",
	0x55: "Socket LGA1851",
	0x56: "Socket BGA2114",
	0x57: "Socket BGA2833",
	0xFF: "Use this when no other valid enumeration is available.",
}
func (s *Type4ProcessorInformation) GetProcessorTypeString() string {
	if v, ok := Type4_ProcessorTypeMap[s.ProcessorType]; ok {
		return v
	}
	return "Unknown"
}

func (s *Type4ProcessorInformation) GetProcessorFamilyString() string {
	if v, ok := Type4_ProcessorFamilyMap[s.ProcessorFamily]; ok {
		return v
	}
	return "Unknown"
}

func (s *Type4ProcessorInformation) GetProcessorUpgradeString() string {
	if v, ok := Type4_ProcessorUpgradeMap[s.ProcessorUpgrade]; ok {
		return v
	}
	return "Unknown"
}
// 
const (
	Type4_Child_Bit0 = 1 << 0 // Reserved
	Type4_Child_Bit1 = 1 << 1 // Unknown
	Type4_Child_Bit2 = 1 << 2 // 64-bit Capable
	Type4_Child_Bit3 = 1 << 3 // Multi-Core
	Type4_Child_Bit4 = 1 << 4 // Hardware Thread
	Type4_Child_Bit5 = 1 << 5 // Execute Protection
	Type4_Child_Bit6 = 1 << 6 // Enhanced Virtualization
	Type4_Child_Bit7 = 1 << 7 // Power/Performance Control
	Type4_Child_Bit8 = 1 << 8 // 128-bit Capable
	Type4_Child_Bit9 = 1 << 9 // Arm64 SoC ID
)

// Memory Controller Information (Type 5, Obsolete) structure (Type 5)
type Type5MemoryController struct {
	Type byte //
	Length byte //
	Handle uint16 //
	ErrorDetectingMethod byte //ENUM
	ErrorCorrectingCapability byte //Bit Field
	SupportedInterleave byte //ENUM
	CurrentInterleave byte //ENUM
	MaximumMemoryModuleSize byte //
	SupportedSpeeds uint16 //Bit Field
	SupportedMemoryTypes uint16 //Bit Field
	MemoryModuleVoltage byte //Bit Field
	NumberofAssociatedMemorySlotsx byte //
	MemoryModuleConfigurationHandles byte // Type:x WORDs
	EnabledErrorCorrectingCapabilities byte //Bit Field
}
// Memory Controller Error Detecting Method field
var Type5_MemoryControllerErrorDetectingMethodMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "None",
	0x04: "8-bit Parity",
	0x05: "32-bit ECC",
	0x06: "64-bit ECC",
	0x07: "128-bit ECC",
	0x08: "CRC",
}

// Memory Controller Information: Interleave Support field
var Type5_InterleaveSupportMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "One-Way Interleave",
	0x04: "Two-Way Interleave",
	0x05: "Four-Way Interleave",
	0x06: "Eight-Way Interleave",
	0x07: "Sixteen-Way Interleave",
}
func (s *Type5MemoryController) GetErrorDetectingMethodString() string {
	if v, ok := Type5_MemoryControllerErrorDetectingMethodMap[s.ErrorDetectingMethod]; ok {
		return v
	}
	return "Unknown"
}

func (s *Type5MemoryController) GetSupportedInterleaveString() string {
	if v, ok := Type5_InterleaveSupportMap[s.SupportedInterleave]; ok {
		return v
	}
	return "Unknown"
}

func (s *Type5MemoryController) GetCurrentInterleaveString() string {
	if v, ok := Type5_InterleaveSupportMap[s.CurrentInterleave]; ok {
		return v
	}
	return "Unknown"
}
// Memory Controller Error Correcting Capability field
const (
	Type5_MemoryControllerErrorCorrectingCapability_Bit0 = 1 << 0 // Other
	Type5_MemoryControllerErrorCorrectingCapability_Bit1 = 1 << 1 // Unknown
	Type5_MemoryControllerErrorCorrectingCapability_Bit2 = 1 << 2 // None
	Type5_MemoryControllerErrorCorrectingCapability_Bit3 = 1 << 3 // Single-Bit Error Correcting
	Type5_MemoryControllerErrorCorrectingCapability_Bit4 = 1 << 4 // Double-Bit Error Correcting
	Type5_MemoryControllerErrorCorrectingCapability_Bit5 = 1 << 5 // Error Scrubbing
)

// Memory Controller Information: Memory Speeds Bit field
const (
	Type5_MemorySpeedsBit_Bit0 = 1 << 0 // Other
	Type5_MemorySpeedsBit_Bit1 = 1 << 1 // Unknown
	Type5_MemorySpeedsBit_Bit2 = 1 << 2 // 70ns
	Type5_MemorySpeedsBit_Bit3 = 1 << 3 // 60ns
	Type5_MemorySpeedsBit_Bit4 = 1 << 4 // 50ns
)
func (s *Type5MemoryController) GetErrorCorrectingCapabilityFlags() []string {
	v := s.ErrorCorrectingCapability
	var flags []string
	if v&(1<<0) != 0 { flags = append(flags, "Other") }
	if v&(1<<1) != 0 { flags = append(flags, "Unknown") }
	if v&(1<<2) != 0 { flags = append(flags, "None") }
	if v&(1<<3) != 0 { flags = append(flags, "Single-Bit Error Correcting") }
	if v&(1<<4) != 0 { flags = append(flags, "Double-Bit Error Correcting") }
	if v&(1<<5) != 0 { flags = append(flags, "Error Scrubbing") }
	return flags
}

func (s *Type5MemoryController) GetSupportedSpeedsFlags() []string {
	v := s.SupportedSpeeds
	var flags []string
	if v&(1<<0) != 0 { flags = append(flags, "Other") }
	if v&(1<<1) != 0 { flags = append(flags, "Unknown") }
	if v&(1<<2) != 0 { flags = append(flags, "70ns") }
	if v&(1<<3) != 0 { flags = append(flags, "60ns") }
	if v&(1<<4) != 0 { flags = append(flags, "50ns") }
	return flags
}

func (s *Type5MemoryController) GetSupportedMemoryTypesFlags() []string {
	v := byte(s.SupportedMemoryTypes)
	var flags []string
	if v&(1<<0) != 0 { flags = append(flags, "Other") }
	if v&(1<<1) != 0 { flags = append(flags, "Unknown") }
	if v&(1<<2) != 0 { flags = append(flags, "None") }
	if v&(1<<3) != 0 { flags = append(flags, "Single-Bit Error Correcting") }
	if v&(1<<4) != 0 { flags = append(flags, "Double-Bit Error Correcting") }
	if v&(1<<5) != 0 { flags = append(flags, "Error Scrubbing") }
	return flags
}

func (s *Type5MemoryController) GetMemoryModuleVoltageFlags() []string {
	v := s.MemoryModuleVoltage
	var flags []string
	if v&(1<<0) != 0 { flags = append(flags, "Other") }
	if v&(1<<1) != 0 { flags = append(flags, "Unknown") }
	if v&(1<<2) != 0 { flags = append(flags, "None") }
	if v&(1<<3) != 0 { flags = append(flags, "Single-Bit Error Correcting") }
	if v&(1<<4) != 0 { flags = append(flags, "Double-Bit Error Correcting") }
	if v&(1<<5) != 0 { flags = append(flags, "Error Scrubbing") }
	return flags
}

func (s *Type5MemoryController) GetEnabledErrorCorrectingCapabilitiesFlags() []string {
	v := s.EnabledErrorCorrectingCapabilities
	var flags []string
	if v&(1<<0) != 0 { flags = append(flags, "Other") }
	if v&(1<<1) != 0 { flags = append(flags, "Unknown") }
	if v&(1<<2) != 0 { flags = append(flags, "None") }
	if v&(1<<3) != 0 { flags = append(flags, "Single-Bit Error Correcting") }
	if v&(1<<4) != 0 { flags = append(flags, "Double-Bit Error Correcting") }
	if v&(1<<5) != 0 { flags = append(flags, "Error Scrubbing") }
	return flags
}

// Memory Module Information (Type 6, Obsolete) structure (Type 6)
type Type6MemoryModuleConfiguration struct {
	Type byte //
	Length byte //
	Handle uint16 //
	SocketDesignation byte //STRING
	BankConnections byte //
	CurrentSpeed byte //
	CurrentMemoryType uint16 //Bit Field
	InstalledSize byte //
	EnabledSize byte //
	ErrorStatus byte //
}
func (s *Type6MemoryModuleConfiguration) GetSocketDesignation(strings []string) string {
	idx := int(s.SocketDesignation)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}
// Memory Module Information: Memory Types
const (
	Type6_MemoryTypes_Bit0 = 1 << 0 // Other
	Type6_MemoryTypes_Bit1 = 1 << 1 // Unknown
	Type6_MemoryTypes_Bit2 = 1 << 2 // Standard
	Type6_MemoryTypes_Bit3 = 1 << 3 // Fast Page Mode
	Type6_MemoryTypes_Bit4 = 1 << 4 // EDO
	Type6_MemoryTypes_Bit5 = 1 << 5 // Parity
	Type6_MemoryTypes_Bit6 = 1 << 6 // ECC
	Type6_MemoryTypes_Bit7 = 1 << 7 // SIMM
	Type6_MemoryTypes_Bit8 = 1 << 8 // DIMM
	Type6_MemoryTypes_Bit9 = 1 << 9 // Burst EDO
	Type6_MemoryTypes_Bit10 = 1 << 10 // SDRAM
)
func (s *Type6MemoryModuleConfiguration) GetCurrentMemoryTypeFlags() []string {
	v := s.CurrentMemoryType
	var flags []string
	if v&(1<<0) != 0 { flags = append(flags, "Other") }
	if v&(1<<1) != 0 { flags = append(flags, "Unknown") }
	if v&(1<<2) != 0 { flags = append(flags, "Standard") }
	if v&(1<<3) != 0 { flags = append(flags, "Fast Page Mode") }
	if v&(1<<4) != 0 { flags = append(flags, "EDO") }
	if v&(1<<5) != 0 { flags = append(flags, "Parity") }
	if v&(1<<6) != 0 { flags = append(flags, "ECC") }
	if v&(1<<7) != 0 { flags = append(flags, "SIMM") }
	if v&(1<<8) != 0 { flags = append(flags, "DIMM") }
	if v&(1<<9) != 0 { flags = append(flags, "Burst EDO") }
	if v&(1<<10) != 0 { flags = append(flags, "SDRAM") }
	return flags
}

// Cache Information (Type 7) (Type 7)
type Type7CacheInformation struct {
	Type byte //
	Length byte //
	Handle uint16 //
	SocketDesignation byte //STRING
	CacheConfiguration uint16 //
	MaximumCacheSize uint16 //
	InstalledSize uint16 //
	SupportedSRAMType uint16 //Bit Field
	CurrentSRAMType uint16 //Bit Field
	CacheSpeed byte //
	ErrorCorrectionType byte //ENUM
	SystemCacheType byte //ENUM
	Associativity byte //ENUM
	MaximumCacheSize2 uint32 //Bit Field
	InstalledCacheSize2 uint32 //Bit Field
}
func (s *Type7CacheInformation) GetSocketDesignation(strings []string) string {
	idx := int(s.SocketDesignation)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}
// Cache Information: Error Correction Type field
var Type7_ErrorCorrectionTypeMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "None",
	0x04: "Parity",
	0x05: "Single-bit ECC",
	0x06: "Multi-bit ECC",
}

// Cache Information: System Cache Type Field
var Type7_SystemCacheTypeMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "Instruction",
	0x04: "Data",
	0x05: "Unified",
}

// Cache Information: Associativity field
var Type7_AssociativityMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "Direct Mapped",
	0x04: "2-way Set-Associative",
	0x05: "4-way Set-Associative",
	0x06: "Fully Associative",
	0x07: "8-way Set-Associative",
	0x08: "16-way Set-Associative",
	0x09: "12-way Set-Associative",
	0x0A: "24-way Set-Associative",
	0x0B: "32-way Set-Associative",
	0x0C: "48-way Set-Associative",
	0x0D: "64-way Set-Associative",
	0x0E: "20-way Set-Associative",
}
func (s *Type7CacheInformation) GetErrorCorrectionTypeString() string {
	if v, ok := Type7_ErrorCorrectionTypeMap[s.ErrorCorrectionType]; ok {
		return v
	}
	return "Unknown"
}

func (s *Type7CacheInformation) GetSystemCacheTypeString() string {
	if v, ok := Type7_SystemCacheTypeMap[s.SystemCacheType]; ok {
		return v
	}
	return "Unknown"
}

func (s *Type7CacheInformation) GetAssociativityString() string {
	if v, ok := Type7_AssociativityMap[s.Associativity]; ok {
		return v
	}
	return "Unknown"
}
// Cache Information: SRAM Type field
const (
	Type7_SRAMType_Bit0 = 1 << 0 // Other
	Type7_SRAMType_Bit1 = 1 << 1 // Unknown
	Type7_SRAMType_Bit2 = 1 << 2 // Non-Burst
	Type7_SRAMType_Bit3 = 1 << 3 // Burst
)
func (s *Type7CacheInformation) GetSupportedSRAMTypeFlags() []string {
	v := s.SupportedSRAMType
	var flags []string
	if v&(1<<0) != 0 { flags = append(flags, "Other") }
	if v&(1<<1) != 0 { flags = append(flags, "Unknown") }
	if v&(1<<2) != 0 { flags = append(flags, "Non-Burst") }
	if v&(1<<3) != 0 { flags = append(flags, "Burst") }
	return flags
}

func (s *Type7CacheInformation) GetCurrentSRAMTypeFlags() []string {
	v := s.CurrentSRAMType
	var flags []string
	if v&(1<<0) != 0 { flags = append(flags, "Other") }
	if v&(1<<1) != 0 { flags = append(flags, "Unknown") }
	if v&(1<<2) != 0 { flags = append(flags, "Non-Burst") }
	if v&(1<<3) != 0 { flags = append(flags, "Burst") }
	return flags
}

func (s *Type7CacheInformation) GetMaximumCacheSize2Flags() []string {
	v := uint16(s.MaximumCacheSize2)
	var flags []string
	if v&(1<<0) != 0 { flags = append(flags, "Other") }
	if v&(1<<1) != 0 { flags = append(flags, "Unknown") }
	if v&(1<<2) != 0 { flags = append(flags, "Non-Burst") }
	if v&(1<<3) != 0 { flags = append(flags, "Burst") }
	return flags
}

func (s *Type7CacheInformation) GetInstalledCacheSize2Flags() []string {
	v := uint16(s.InstalledCacheSize2)
	var flags []string
	if v&(1<<0) != 0 { flags = append(flags, "Other") }
	if v&(1<<1) != 0 { flags = append(flags, "Unknown") }
	if v&(1<<2) != 0 { flags = append(flags, "Non-Burst") }
	if v&(1<<3) != 0 { flags = append(flags, "Burst") }
	return flags
}

// Port Connector Information (Type 8) (Type 8)
type Type8ConnectorInformation struct {
	Type byte //
	Length byte //
	Handle uint16 //
	InternalReferenceDesignator byte //STRING
	InternalConnectorType byte //ENUM
	ExternalReferenceDesignator byte //STRING
	ExternalConnectorType byte //ENUM
	PortType byte //ENUM
}
func (s *Type8ConnectorInformation) GetInternalReferenceDesignator(strings []string) string {
	idx := int(s.InternalReferenceDesignator)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type8ConnectorInformation) GetExternalReferenceDesignator(strings []string) string {
	idx := int(s.ExternalReferenceDesignator)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}
// Port Information: Connector Types Field
var Type8_ConnectorTypesMap = map[byte]string{
	0x00: "None",
	0x01: "Centronics",
	0x02: "Mini Centronics",
	0x03: "Proprietary",
	0x04: "DB-25 pin male",
	0x05: "DB-25 pin female",
	0x06: "DB-15 pin male",
	0x07: "DB-15 pin female",
	0x08: "DB-9 pin male",
	0x09: "DB-9 pin female",
	0x0A: "RJ-11",
	0x0B: "RJ-45",
	0x0C: "50-pin MiniSCSI",
	0x0D: "Mini-DIN",
	0x0E: "Micro-DIN",
	0x0F: "PS/2",
	0x10: "Infrared",
	0x11: "HP-HIL",
	0x12: "Access Bus (USB)",
	0x13: "SSA SCSI",
	0x14: "Circular DIN-8 male",
	0x15: "Circular DIN-8 female",
	0x16: "On Board IDE",
	0x17: "On Board Floppy",
	0x18: "9-pin Dual Inline (pin 10 cut)",
	0x19: "25-pin Dual Inline (pin 26 cut)",
	0x1A: "50-pin Dual Inline",
	0x1B: "68-pin Dual Inline",
	0x1C: "On Board Sound Input from CD-ROM",
	0x1D: "Mini-Centronics Type-14",
	0x1E: "Mini-Centronics Type-26",
	0x1F: "Mini-jack (headphones)",
	0x20: "BNC",
	0x21: "1394",
	0x22: "SAS/SATA Plug Receptacle",
	0x23: "USB Type-C Receptacle",
	0xA0: "PC-98",
	0xA1: "PC-98Hireso",
	0xA2: "PC-H98",
	0xA3: "PC-98Note",
	0xA4: "PC-98Full",
	0xFF: "Other - Use Reference Designator Strings to supply information.",
}

// Port Types field
var Type8_PortTypesMap = map[byte]string{
	0x00: "None",
	0x01: "Parallel Port XT/AT Compatible",
	0x02: "Parallel Port PS/2",
	0x03: "Parallel Port ECP",
	0x04: "Parallel Port EPP",
	0x05: "Parallel Port ECP/EPP",
	0x06: "Serial Port XT/AT Compatible",
	0x07: "Serial Port 16450 Compatible",
	0x08: "Serial Port 16550 Compatible",
	0x09: "Serial Port 16550A Compatible",
	0x0A: "SCSI Port",
	0x0B: "MIDI Port",
	0x0C: "Joy Stick Port",
	0x0D: "Keyboard Port",
	0x0E: "Mouse Port",
	0x0F: "SSA SCSI",
	0x10: "USB",
	0x11: "FireWire (IEEE P1394)",
	0x12: "PCMCIA Type I",
	0x13: "PCMCIA Type II",
	0x14: "PCMCIA Type III",
	0x15: "Card bus",
	0x16: "Access Bus Port",
	0x17: "SCSI II",
	0x18: "SCSI Wide",
	0x19: "PC-98",
	0x1A: "PC-98-Hireso",
	0x1B: "PC-H98",
	0x1C: "Video Port",
	0x1D: "Audio Port",
	0x1E: "Modem Port",
	0x1F: "Network Port",
	0x20: "SATA",
	0x21: "SAS",
	0x22: "MFDP (Multi-Function Display Port)",
	0x23: "Thunderbolt",
	0xA0: "8251 Compatible",
	0xA1: "8251 FIFO Compatible",
	0xFF: "Other",
}
func (s *Type8ConnectorInformation) GetInternalConnectorTypeString() string {
	if v, ok := Type8_ConnectorTypesMap[s.InternalConnectorType]; ok {
		return v
	}
	return "Unknown"
}

func (s *Type8ConnectorInformation) GetExternalConnectorTypeString() string {
	if v, ok := Type8_ConnectorTypesMap[s.ExternalConnectorType]; ok {
		return v
	}
	return "Unknown"
}

func (s *Type8ConnectorInformation) GetPortTypeString() string {
	if v, ok := Type8_PortTypesMap[s.PortType]; ok {
		return v
	}
	return "Unknown"
}

// System Slots (Type 9) structure (Type 9)
type Type9SystemSlot struct {
	Type byte //
	Length byte //
}
// System Slots: Slot Type field
var Type9_SlotTypeMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "ISA",
	0x04: "MCA",
	0x05: "EISA",
	0x06: "PCI",
	0x07: "PC Card (PCMCIA)",
	0x08: "VL-VESA",
	0x09: "Proprietary",
	0x0A: "Processor Card Slot",
	0x0B: "Proprietary Memory Card Slot",
	0x0C: "I/O Riser Card Slot",
	0x0D: "NuBus",
	0x0E: "PCI - 66MHz Capable",
	0x0F: "AGP",
	0x10: "AGP 2X",
	0x11: "AGP 4X",
	0x12: "PCI-X",
	0x13: "AGP 8X",
	0x14: "M.2 Socket 1-DP (Mechanical Key A)",
	0x15: "M.2 Socket 1-SD (Mechanical Key E)",
	0x16: "M.2 Socket 2 (Mechanical Key B)",
	0x17: "M.2 Socket 3 (Mechanical Key M)",
	0x18: "MXM Type I",
	0x19: "MXM Type II",
	0x1A: "MXM Type III (standard connector)",
	0x1B: "MXM Type III (HE connector)",
	0x1C: "MXM Type IV",
	0x1D: "MXM 3.0 Type A",
	0x1E: "MXM 3.0 Type B",
	0x1F: "PCI Express Gen 2 SFF-8639 (U.2)",
	0x20: "PCI Express Gen 3 SFF-8639 (U.2)",
	0x21: "PCI Express Mini 52-pin (CEM spec. 2.0) with bottom-side keep-outs. Use Slot Length field value 03h (short length) for \"half-Mini card\"-only support, 04h (long length) for \"full-Mini card\" or dual support.",
	0x22: "PCI Express Mini 52-pin (CEM spec. 2.0) without bottom-side keep-outs. Use Slot Length field value 03h (short length) for \"half-Mini card\"-only support, 04h (long length) for \"full-Mini card\" or dual support.",
	0x23: "PCI Express Mini 76-pin (CEM spec. 2.0) Corresponds to Display-Mini card.",
	0x24: "PCI Express Gen 4 SFF-8639 (U.2)",
	0x25: "PCI Express Gen 5 SFF-8639 (U.2)",
	0x26: "OCP NIC 3.0 Small Form Factor (SFF)",
	0x27: "OCP NIC 3.0 Large Form Factor (LFF)",
	0x28: "OCP NIC Prior to 3.0",
	0x30: "CXL Flexbus 1.0 (deprecated, see note below)",
	0xA0: "PC-98/C20",
	0xA1: "PC-98/C24",
	0xA2: "PC-98/E",
	0xA3: "PC-98/Local Bus",
	0xA4: "PC-98/Card",
	0xA5: "PCI Express (see note below)",
	0xA6: "PCI Express x1",
	0xA7: "PCI Express x2",
	0xA8: "PCI Express x4",
	0xA9: "PCI Express x8",
	0xAA: "PCI Express x16",
	0xAB: "PCI Express Gen 2 (see note below)",
	0xAC: "PCI Express Gen 2 x1",
	0xAD: "PCI Express Gen 2 x2",
	0xAE: "PCI Express Gen 2 x4",
	0xAF: "PCI Express Gen 2 x8",
	0xB0: "PCI Express Gen 2 x16",
	0xB1: "PCI Express Gen 3 (see note below)",
	0xB2: "PCI Express Gen 3 x1",
	0xB3: "PCI Express Gen 3 x2",
	0xB4: "PCI Express Gen 3 x4",
	0xB5: "PCI Express Gen 3 x8",
	0xB6: "PCI Express Gen 3 x16",
	0xB8: "PCI Express Gen 4 (see note below)",
	0xB9: "PCI Express Gen 4 x1",
	0xBA: "PCI Express Gen 4 x2",
	0xBB: "PCI Express Gen 4 x4",
	0xBC: "PCI Express Gen 4 x8",
	0xBD: "PCI Express Gen 4 x16",
	0xBE: "PCI Express Gen 5 (see note below)",
	0xBF: "PCI Express Gen 5 x1",
	0xC0: "PCI Express Gen 5 x2",
	0xC1: "PCI Express Gen 5 x4",
	0xC2: "PCI Express Gen 5 x8",
	0xC3: "PCI Express Gen 5 x16",
	0xC4: "PCI Express Gen 6 and Beyond (see Slot Information and Slot Physical Width fields for more details)",
	0xC5: "Enterprise and Datacenter 1U E1 Form Factor Slot (EDSFF E1.S, E1.L)",
	0xC6: "Enterprise and Datacenter 3\" E3 Form Factor Slot (EDSFF E3.S, E3.L)",
}

// System Slots: Slot Width field
var Type9_SlotWidthMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "8 bit",
	0x04: "16 bit",
	0x05: "32 bit",
	0x06: "64 bit",
	0x07: "128 bit",
	0x08: "1x or x1",
	0x09: "2x or x2",
	0x0A: "4x or x4",
	0x0B: "8x or x8",
	0x0C: "12x or x12",
	0x0D: "16x or x16",
	0x0E: "32x or x32",
}

// System Slots: Current Usage field
var Type9_CurrentUsageMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "Available",
	0x04: "In use",
	0x05: "Unavailable",
}

// System Slots: Slot Length field
var Type9_SlotLengthMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "Short Length",
	0x04: "Long Length",
	0x05: "2.5\" drive form factor",
	0x06: "3.5\" drive form factor",
}

// 
var Type9_ChildMap = map[byte]string{
	0x00: "Not applicable",
	0x01: "Other",
	0x02: "Unknown",
	0x03: "Full height",
	0x04: "Low-profile",
}
// Slot Characteristics 1 field
const (
	Type9_SlotCharacteristics1_Bit0 = 1 << 0 // Characteristics unknown.
	Type9_SlotCharacteristics1_Bit1 = 1 << 1 // Provides 5.0 volts.
	Type9_SlotCharacteristics1_Bit2 = 1 << 2 // Provides 3.3 volts.
	Type9_SlotCharacteristics1_Bit3 = 1 << 3 // Slot’s opening is shared with another slot (for example, PCI/EISA shared slot).
	Type9_SlotCharacteristics1_Bit4 = 1 << 4 // PC Card slot supports PC Card-16.
	Type9_SlotCharacteristics1_Bit5 = 1 << 5 // PC Card slot supports CardBus.
	Type9_SlotCharacteristics1_Bit6 = 1 << 6 // PC Card slot supports Zoom Video.
	Type9_SlotCharacteristics1_Bit7 = 1 << 7 // PC Card slot supports Modem Ring Resume.
)

// Slot Characteristics 2
const (
	Type9_SlotCharacteristics2_Bit0 = 1 << 0 // PCI slot supports Power Management Event (PME#) signal.
	Type9_SlotCharacteristics2_Bit1 = 1 << 1 // Slot supports hot-plug devices.
	Type9_SlotCharacteristics2_Bit2 = 1 << 2 // PCI slot supports SMBus signal.
	Type9_SlotCharacteristics2_Bit3 = 1 << 3 // PCIe slot supports bifurcation. This slot can partition its lanes into two or mo
	Type9_SlotCharacteristics2_Bit4 = 1 << 4 // Slot supports async/surprise removal, such as removal without prior notification
	Type9_SlotCharacteristics2_Bit5 = 1 << 5 // Flexbus slot, CXL 1.0 capable
	Type9_SlotCharacteristics2_Bit6 = 1 << 6 // Flexbus slot, CXL 2.0 capable
	Type9_SlotCharacteristics2_Bit7 = 1 << 7 // Flexbus slot, CXL 3.0 capable
)

// On Board Devices Information (Type 10, Obsolete) structure (Type 10)
type Type10OnBoardDevicesInformation struct {
	Type byte //
	Length byte //
	Handle uint16 //
	DeviceTypenrangesfrom1toNumberofDevices byte //
	DescriptionString byte //STRING
}
func (s *Type10OnBoardDevicesInformation) GetDescriptionString(strings []string) string {
	idx := int(s.DescriptionString)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}
// Onboard Device Types
var Type10_OnboardDeviceTypesMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "Video",
	0x04: "SCSI Controller",
	0x05: "Ethernet",
	0x06: "Token Ring",
	0x07: "Sound",
	0x08: "PATA Controller",
	0x09: "SATA Controller",
	0x0A: "SAS Controller",
}

// OEM Strings (Type 11) (Type 11)
type Type11OEMStrings struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Count byte //
}

// System Configuration Options (Type 12) structure (Type 12)
type Type12ConfigurationInformation struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Count byte //
}

// Firmware Language Information (Type 13) structure (Type 13)
type Type13LanguageInformation struct {
	Type byte //
	Length byte //
	Handle uint16 //
	InstallableLanguages byte //
	Flags byte //Bit Field
	Reserved [15]byte //
	CurrentLanguage byte //STRING
}
func (s *Type13LanguageInformation) GetCurrentLanguage(strings []string) string {
	idx := int(s.CurrentLanguage)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

// Group Associations (Type 14) structure (Type 14)
type Type14GroupAssociations struct {
	Type byte //
	Length byte //
	Handle uint16 //
	GroupName byte //STRING
	ItemType byte //
	ItemHandle uint16 //
}
func (s *Type14GroupAssociations) GetGroupName(strings []string) string {
	idx := int(s.GroupName)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

// System Event Log (Type 15) (Type 15)
type Type15EventLog struct {
	Type byte //
	Length byte //
	Handle uint16 //
	LogAreaLength uint16 //
	LogHeaderStartOffset uint16 //
	LogDataStartOffset uint16 //
	AccessMethod byte //
	LogStatus byte //
	LogChangeToken uint32 //
	AccessMethodAddress uint32 //
	LogHeaderFormat byte //ENUM
	NumberofSupportedLogTypeDescriptorsx byte //
	LengthofeachLogTypeDescriptory byte //
	ListofSupportedEventLogTypeDescriptors byte // Type:Varies
}
// Log Header format
var Type15_LogHeaderformatMap = map[byte]string{
	0x00: "No header (for example, the header is 0 bytes in length)",
	0x01: "Type 1 log header; see 7.16.5.1",
}
func (s *Type15EventLog) GetLogHeaderFormatString() string {
	if v, ok := Type15_LogHeaderformatMap[s.LogHeaderFormat]; ok {
		return v
	}
	return "Unknown"
}

// Physical Memory Array (Type 16) structure (Type 16)
type Type16PhysicalMemoryArray struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Location byte //ENUM
	Use byte //ENUM
	MemoryErrorCorrection byte //ENUM
	MaximumCapacity uint32 //
	MemoryErrorInformationHandle uint16 //
	NumberofMemoryDevices uint16 //
	ExtendedMaximumCapacity uint64 //
}
// Memory Array: Location field
var Type16_LocationMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "System board or motherboard",
	0x04: "ISA add-on card",
	0x05: "EISA add-on card",
	0x06: "PCI add-on card",
	0x07: "MCA add-on card",
	0x08: "PCMCIA add-on card",
	0x09: "Proprietary add-on card",
	0x0A: "NuBus",
	0xA0: "PC-98/C20 add-on card",
	0xA1: "PC-98/C24 add-on card",
	0xA2: "PC-98/E add-on card",
	0xA3: "PC-98/Local bus add-on card",
	0xA4: "CXL add-on card",
}

// Memory Array: Use field
var Type16_UseMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "System memory",
	0x04: "Video memory",
	0x05: "Flash memory",
	0x06: "Non-volatile RAM",
	0x07: "Cache memory",
}

// Memory Array: Error Correction Types field
var Type16_ErrorCorrectionTypesMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "None",
	0x04: "Parity",
	0x05: "Single-bit ECC",
	0x06: "Multi-bit ECC",
	0x07: "CRC",
}
func (s *Type16PhysicalMemoryArray) GetLocationString() string {
	if v, ok := Type16_LocationMap[s.Location]; ok {
		return v
	}
	return "Unknown"
}

func (s *Type16PhysicalMemoryArray) GetUseString() string {
	if v, ok := Type16_UseMap[s.Use]; ok {
		return v
	}
	return "Unknown"
}

func (s *Type16PhysicalMemoryArray) GetMemoryErrorCorrectionString() string {
	if v, ok := Type16_ErrorCorrectionTypesMap[s.MemoryErrorCorrection]; ok {
		return v
	}
	return "Unknown"
}

// Memory Device (Type 17) (Type 17)
type Type17MemoryDevice struct {
	Type byte //
	Length byte //
	Handle uint16 //
	PhysicalMemoryArrayHandle uint16 //
	MemoryErrorInformationHandle uint16 //
	TotalWidth uint16 //
	DataWidth uint16 //
	Size uint16 //
	FormFactor byte //ENUM
	DeviceSet byte //
	DeviceLocator byte //STRING
	BankLocator byte //STRING
	MemoryType byte //ENUM
	TypeDetail uint16 //Bit Field
	Speed uint16 //
	Manufacturer byte //STRING
	SerialNumber byte //STRING
	AssetTag byte //STRING
	PartNumber byte //STRING
	Attributes byte //
	ExtendedSize uint32 //
	ConfiguredMemorySpeed uint16 //
	Minimumvoltage uint16 //
	Maximumvoltage uint16 //
	Configuredvoltage uint16 //
	MemoryTechnology byte //
	MemoryOperatingModeCapability uint16 //Bit Field
	FirmwareVersion byte //STRING
	ModuleManufacturerID uint16 //
	ModuleProductID uint16 //
	MemorySubsystemControllerManufacturerID uint16 //
	MemorySubsystemControllerProductID uint16 //
	NonvolatileSize uint64 //
	VolatileSize uint64 //
	CacheSize uint64 //
	LogicalSize uint64 //
	ExtendedSpeed uint32 //
	ExtendedConfiguredMemorySpeed uint32 //
	PMIC0ManufacturerID uint16 //
	PMIC0RevisionNumber uint16 //
	RCDManufacturerID uint16 //
	RCDRevisionNumber uint16 //
}
func (s *Type17MemoryDevice) GetDeviceLocator(strings []string) string {
	idx := int(s.DeviceLocator)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type17MemoryDevice) GetBankLocator(strings []string) string {
	idx := int(s.BankLocator)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type17MemoryDevice) GetManufacturer(strings []string) string {
	idx := int(s.Manufacturer)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type17MemoryDevice) GetSerialNumber(strings []string) string {
	idx := int(s.SerialNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type17MemoryDevice) GetAssetTag(strings []string) string {
	idx := int(s.AssetTag)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type17MemoryDevice) GetPartNumber(strings []string) string {
	idx := int(s.PartNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type17MemoryDevice) GetFirmwareVersion(strings []string) string {
	idx := int(s.FirmwareVersion)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}
// Memory Device: Form Factor field
var Type17_FormFactorMap = map[byte]string{
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
	0x11: "CAMM",
	0x12: "CUDIMM",
	0x13: "CSODIMM",
}

// Memory Device: Type
var Type17_TypeMap = map[byte]string{
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
	0x22: "DDR5",
	0x23: "LPDDR5",
	0x24: "HBM3 (High Bandwidth Memory Generation 3)",
	0x25: "MRDIMM",
}

// Memory Device: Memory Technology field
var Type17_MemoryTechnologyMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "DRAM",
	0x04: "NVDIMM-N",
	0x05: "NVDIMM-F",
	0x06: "NVDIMM-P",
	0x07: "Intel® Optane™ persistent memory",
	0x08: "MRDIMM (Deprecated). This value has been deprecated from this table and moved to Memory Device - Type, subclause 7.18.2",
}
func (s *Type17MemoryDevice) GetFormFactorString() string {
	if v, ok := Type17_FormFactorMap[s.FormFactor]; ok {
		return v
	}
	return "Unknown"
}

func (s *Type17MemoryDevice) GetMemoryTypeString() string {
	if v, ok := Type17_TypeMap[s.MemoryType]; ok {
		return v
	}
	return "Unknown"
}
// Memory Device: Type Detail field
const (
	Type17_TypeDetail_Bit0 = 1 << 0 // Reserved, set to 0
	Type17_TypeDetail_Bit1 = 1 << 1 // Other
	Type17_TypeDetail_Bit2 = 1 << 2 // Unknown
	Type17_TypeDetail_Bit3 = 1 << 3 // Fast-paged
	Type17_TypeDetail_Bit4 = 1 << 4 // Static column
	Type17_TypeDetail_Bit5 = 1 << 5 // Pseudo-static
	Type17_TypeDetail_Bit6 = 1 << 6 // RAMBUS
	Type17_TypeDetail_Bit7 = 1 << 7 // Synchronous
	Type17_TypeDetail_Bit8 = 1 << 8 // CMOS
	Type17_TypeDetail_Bit9 = 1 << 9 // EDO
	Type17_TypeDetail_Bit10 = 1 << 10 // Window DRAM
	Type17_TypeDetail_Bit11 = 1 << 11 // Cache DRAM
	Type17_TypeDetail_Bit12 = 1 << 12 // Non-volatile
	Type17_TypeDetail_Bit13 = 1 << 13 // Registered (Buffered)
	Type17_TypeDetail_Bit14 = 1 << 14 // Unbuffered (Unregistered)
	Type17_TypeDetail_Bit15 = 1 << 15 // LRDIMM
)

// Memory Device: Memory Operating Mode Capability
const (
	Type17_MemoryOperatingModeCapability_Bit0 = 1 << 0 // Reserved, set to 0
	Type17_MemoryOperatingModeCapability_Bit1 = 1 << 1 // Other
	Type17_MemoryOperatingModeCapability_Bit2 = 1 << 2 // Unknown
	Type17_MemoryOperatingModeCapability_Bit3 = 1 << 3 // Volatile memory
	Type17_MemoryOperatingModeCapability_Bit4 = 1 << 4 // Byte-accessible persistent memory
	Type17_MemoryOperatingModeCapability_Bit5 = 1 << 5 // Block-accessible persistent memory
	Type17_MemoryOperatingModeCapability_Bit6 = 1 << 6 // Reserved, set to 0
)
func (s *Type17MemoryDevice) GetTypeDetailFlags() []string {
	v := s.TypeDetail
	var flags []string
	if v&(1<<0) != 0 { flags = append(flags, "Reserved, set to 0") }
	if v&(1<<1) != 0 { flags = append(flags, "Other") }
	if v&(1<<2) != 0 { flags = append(flags, "Unknown") }
	if v&(1<<3) != 0 { flags = append(flags, "Fast-paged") }
	if v&(1<<4) != 0 { flags = append(flags, "Static column") }
	if v&(1<<5) != 0 { flags = append(flags, "Pseudo-static") }
	if v&(1<<6) != 0 { flags = append(flags, "RAMBUS") }
	if v&(1<<7) != 0 { flags = append(flags, "Synchronous") }
	if v&(1<<8) != 0 { flags = append(flags, "CMOS") }
	if v&(1<<9) != 0 { flags = append(flags, "EDO") }
	if v&(1<<10) != 0 { flags = append(flags, "Window DRAM") }
	if v&(1<<11) != 0 { flags = append(flags, "Cache DRAM") }
	if v&(1<<12) != 0 { flags = append(flags, "Non-volatile") }
	if v&(1<<13) != 0 { flags = append(flags, "Registered (Buffered)") }
	if v&(1<<14) != 0 { flags = append(flags, "Unbuffered (Unregistered)") }
	if v&(1<<15) != 0 { flags = append(flags, "LRDIMM") }
	return flags
}

func (s *Type17MemoryDevice) GetMemoryOperatingModeCapabilityFlags() []string {
	v := s.MemoryOperatingModeCapability
	var flags []string
	if v&(1<<0) != 0 { flags = append(flags, "Reserved, set to 0") }
	if v&(1<<1) != 0 { flags = append(flags, "Other") }
	if v&(1<<2) != 0 { flags = append(flags, "Unknown") }
	if v&(1<<3) != 0 { flags = append(flags, "Volatile memory") }
	if v&(1<<4) != 0 { flags = append(flags, "Byte-accessible persistent memory") }
	if v&(1<<5) != 0 { flags = append(flags, "Block-accessible persistent memory") }
	if v&(1<<6) != 0 { flags = append(flags, "Reserved, set to 0") }
	return flags
}

// 32-Bit Memory Error Information (Type 18) structure (Type 18)
type Type1832bitMemoryErrorInformation struct {
	Type byte //
	Length byte //
	Handle uint16 //
	ErrorType byte //ENUM
	ErrorGranularity byte //ENUM
	ErrorOperation byte //ENUM
	VendorSyndrome uint32 //
	MemoryArrayErrorAddress uint32 //
	DeviceErrorAddress uint32 //
	ErrorResolution uint32 //
}
// Memory Error: Error Type field
var Type18_ErrorTypeMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "OK",
	0x04: "Bad read",
	0x05: "Parity error",
	0x06: "Single-bit error",
	0x07: "Double-bit error",
	0x08: "Multi-bit error",
	0x09: "Nibble error",
	0x0A: "Checksum error",
	0x0B: "CRC error",
	0x0C: "Corrected single-bit error",
	0x0D: "Corrected error",
	0x0E: "Uncorrectable error",
}

// Memory Error: Error Granularity field
var Type18_ErrorGranularityMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "Device level",
	0x04: "Memory partition level",
}

// Memory Error: Error Operation field
var Type18_ErrorOperationMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "Read",
	0x04: "Write",
	0x05: "Partial write",
}
func (s *Type1832bitMemoryErrorInformation) GetErrorTypeString() string {
	if v, ok := Type18_ErrorTypeMap[s.ErrorType]; ok {
		return v
	}
	return "Unknown"
}

func (s *Type1832bitMemoryErrorInformation) GetErrorGranularityString() string {
	if v, ok := Type18_ErrorGranularityMap[s.ErrorGranularity]; ok {
		return v
	}
	return "Unknown"
}

func (s *Type1832bitMemoryErrorInformation) GetErrorOperationString() string {
	if v, ok := Type18_ErrorOperationMap[s.ErrorOperation]; ok {
		return v
	}
	return "Unknown"
}

// Memory Array Mapped Address (Type 19) structure (Type 19)
type Type19MemoryArrayMappedAddress struct {
	Type byte //
	Length byte //
	Handle uint16 //
	StartingAddress uint32 //
	EndingAddress uint32 //
	MemoryArrayHandle uint16 //
	PartitionWidth byte //
	ExtendedStartingAddress uint64 //
	ExtendedEndingAddress uint64 //
}

// Memory Device Mapped Address (Type 20) (Type 20)
type Type20MemoryDeviceMappedAddress struct {
	Type byte //
	Length byte //
	Handle uint16 //
	StartingAddress uint32 //
	EndingAddress uint32 //
	MemoryDeviceHandle uint16 //
	MemoryArrayMappedAddressHandle uint16 //
	PartitionRowPosition byte //
	InterleavePosition byte //
	InterleavedDataDepth byte //
	ExtendedStartingAddress uint64 //
	ExtendedEndingAddress uint64 //
}

// Built (Type 21)
type Type21BuiltinPointingDevice struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Type_1 byte //ENUM
	Interface byte //ENUM
	NumberofButtons byte //
}
// Pointing Device: Type field
var Type21_TypeMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "Mouse",
	0x04: "Track Ball",
	0x05: "Track Point",
	0x06: "Glide Point",
	0x07: "Touch Pad",
	0x08: "Touch Screen",
	0x09: "Optical Sensor",
}

// Pointing Device: Interface field
var Type21_InterfaceMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "Serial",
	0x04: "PS/2",
	0x05: "Infrared",
	0x06: "HP-HIL",
	0x07: "Bus mouse",
	0x08: "ADB (Apple Desktop Bus)",
	0xA0: "Bus mouse DB-9",
	0xA1: "Bus mouse micro-DIN",
	0xA2: "USB",
	0xA3: "I²C",
	0xA4: "SPI",
}
func (s *Type21BuiltinPointingDevice) GetType_1String() string {
	if v, ok := Type21_TypeMap[s.Type_1]; ok {
		return v
	}
	return "Unknown"
}

func (s *Type21BuiltinPointingDevice) GetInterfaceString() string {
	if v, ok := Type21_InterfaceMap[s.Interface]; ok {
		return v
	}
	return "Unknown"
}

// Portable Battery (Type 22) (Type 22)
type Type22PortableBattery struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Location byte //STRING
	Manufacturer byte //STRING
	ManufactureDate byte //STRING
	SerialNumber byte //STRING
	DeviceName byte //STRING
	DeviceChemistry byte //ENUM
	DesignCapacity uint16 //
	DesignVoltage uint16 //
	SBDSVersionNumber byte //STRING
	MaximumErrorinBatteryData byte //
	SBDSSerialNumber uint16 //
	SBDSManufactureDate uint16 //
	SBDSDeviceChemistry byte //STRING
	DesignCapacityMultiplier byte //
	OEMspecific uint32 //
}
func (s *Type22PortableBattery) GetLocation(strings []string) string {
	idx := int(s.Location)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type22PortableBattery) GetManufacturer(strings []string) string {
	idx := int(s.Manufacturer)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type22PortableBattery) GetManufactureDate(strings []string) string {
	idx := int(s.ManufactureDate)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type22PortableBattery) GetSerialNumber(strings []string) string {
	idx := int(s.SerialNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type22PortableBattery) GetDeviceName(strings []string) string {
	idx := int(s.DeviceName)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type22PortableBattery) GetSBDSVersionNumber(strings []string) string {
	idx := int(s.SBDSVersionNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type22PortableBattery) GetSBDSDeviceChemistry(strings []string) string {
	idx := int(s.SBDSDeviceChemistry)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}
// Portable Battery: Device Chemistry field
var Type22_DeviceChemistryMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "Lead Acid",
	0x04: "Nickel Cadmium",
	0x05: "Nickel metal hydride",
	0x06: "Lithium-ion",
	0x07: "Zinc air",
	0x08: "Lithium Polymer",
}
func (s *Type22PortableBattery) GetDeviceChemistryString() string {
	if v, ok := Type22_DeviceChemistryMap[s.DeviceChemistry]; ok {
		return v
	}
	return "Unknown"
}

// System Reset (Type 23) (Type 23)
type Type23SystemReset struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Capabilities byte //
	ResetCount uint16 //
	ResetLimit uint16 //
	TimerInterval uint16 //
	Timeout uint16 //
}

// Hardware Security (Type 24) structure (Type 24)
type Type24HardwareSecurity struct {
	Type byte //
	Length byte //
	Handle uint16 //
	HardwareSecuritySettings byte //
}

// System Power Controls (Type 25) (Type 25)
type Type25SystemPowerControls struct {
	Type byte //
	Length byte //
	Handle uint16 //
	NextScheduledPoweronMonth byte //
	NextScheduledPoweronDayofmonth byte //
	NextScheduledPoweronHour byte //
	NextScheduledPoweronMinute byte //
	NextScheduledPoweronSecond byte //
}

// Voltage Probe (Type 26) (Type 26)
type Type26VoltageProbe struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Description byte //STRING
	LocationandStatus byte //
	MaximumValue uint16 //
	MinimumValue uint16 //
	Resolution uint16 //
	Tolerance uint16 //
	Accuracy uint16 //
	OEMdefined uint32 //
	NominalValue uint16 //
}
func (s *Type26VoltageProbe) GetDescription(strings []string) string {
	idx := int(s.Description)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

// Cooling Device (Type 27) (Type 27)
type Type27CoolingDevice struct {
	Type byte //
	Length byte //
	Handle uint16 //
	TemperatureProbeHandle uint16 //
	DeviceTypeandStatus byte //
	CoolingUnitGroup byte //
	OEMdefined uint32 //
	NominalSpeed uint16 //
	Description byte //STRING
}
func (s *Type27CoolingDevice) GetDescription(strings []string) string {
	idx := int(s.Description)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

// Temperature Probe (Type 28) (Type 28)
type Type28TemperatureProbe struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Description byte //STRING
	LocationandStatus byte //
	MaximumValue uint16 //
	MinimumValue uint16 //
	Resolution uint16 //
	Tolerance uint16 //
	Accuracy uint16 //
	OEMdefined uint32 //
	NominalValue uint16 //
}
func (s *Type28TemperatureProbe) GetDescription(strings []string) string {
	idx := int(s.Description)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

// Electrical Current Probe (Type 29) structure (Type 29)
type Type29ElectricalCurrentProbe struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Description byte //STRING
	LocationandStatus byte //ENUM
	MaximumValue uint16 //
	MinimumValue uint16 //
	Resolution uint16 //
	Tolerance uint16 //
	Accuracy uint16 //
	OEMdefined uint32 //
	NominalValue uint16 //
}
func (s *Type29ElectricalCurrentProbe) GetDescription(strings []string) string {
	idx := int(s.Description)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

// Out (Type 30)
type Type30OutofBandRemoteAccess struct {
	Type byte //
	Length byte //
	Handle uint16 //
	ManufacturerName byte //STRING
	Connections byte //
}
func (s *Type30OutofBandRemoteAccess) GetManufacturerName(strings []string) string {
	idx := int(s.ManufacturerName)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

// System Boot Information (Type 32) structure (Type 32)
type Type32SystemBootInformationidentifier struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Reserved [6]byte //
	BootStatus byte // Type:Length-10 Bytes
}

// 64 (Type 33)
type Type3364bitMemoryErrorInformation struct {
	Type byte //
	Length byte //
	Handle uint16 //
	ErrorType byte //ENUM
	ErrorGranularity byte //ENUM
	ErrorOperation byte //ENUM
	VendorSyndrome uint32 //
	MemoryArrayErrorAddress uint64 //
	DeviceErrorAddress uint64 //
	ErrorResolution uint32 //
}

// Management Device (Type 34) (Type 34)
type Type34ManagementDevice struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Description byte //STRING
	Type_1 byte //
	Address uint32 //
	AddressType byte //
}
func (s *Type34ManagementDevice) GetDescription(strings []string) string {
	idx := int(s.Description)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}
// Management Device: Type field
var Type34_TypeMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "National Semiconductor LM75",
	0x04: "National Semiconductor LM78",
	0x05: "National Semiconductor LM79",
	0x06: "National Semiconductor LM80",
	0x07: "National Semiconductor LM81",
	0x08: "Analog Devices ADM9240",
	0x09: "Dallas Semiconductor DS1780",
	0x0A: "Maxim 1617",
	0x0B: "Genesys GL518SM",
	0x0C: "Winbond W83781D",
	0x0D: "Holtek HT82H791",
}

// Management Device: Address Type field
var Type34_AddressTypeMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "I/O Port",
	0x04: "Memory",
	0x05: "SM Bus",
}

// Management Device Component (Type 35) (Type 35)
type Type35ManagementDeviceComponent struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Description byte //STRING
	ManagementDeviceHandle uint16 //
	ComponentHandle uint16 //
	ThresholdHandle uint16 //
}
func (s *Type35ManagementDeviceComponent) GetDescription(strings []string) string {
	idx := int(s.Description)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

// Management Device Threshold Data (Type 36) (Type 36)
type Type36ManagementDeviceThresholdData struct {
	Type byte //
	Length byte //
	Handle uint16 //
	LowerThresholdNoncritical uint16 //
	UpperThresholdNoncritical uint16 //
	LowerThresholdCritical uint16 //
	UpperThresholdCritical uint16 //
	LowerThresholdNonrecoverable uint16 //
	UpperThresholdNonrecoverable uint16 //
}

// Memory Channel (Type 37) (Type 37)
type Type37ManagementDeviceThresholdData struct {
	Type byte //
	Length byte //
	Handle uint16 //
	ChannelType byte //
	MaximumChannelLoad byte //
	MemoryDeviceCountn byte //
	Memory1DeviceLoad byte //
	MemoryDevice1Handle uint16 //
	MemoryDeviceLoad byte //
	MemoryDeviceHandle uint16 //
}
// Memory Channel: Channel Type field
var Type37_ChannelTypeMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "Rambus",
	0x04: "SyncLink",
}

// IPMI Device Information (Type 38) (Type 38)
type Type38IPMIDeviceInformation struct {
	Type byte //
	Length byte //
	Handle uint16 //
	InterfaceType byte //ENUM
	IPMISpecificationRevision byte //
	I2CTargetAddress byte //
	NVStorageDeviceAddress byte //
	BaseAddress uint64 //
	BaseAddressModifierInterruptInfo byte //
	InterruptNumber byte //
}
// IPMI Device Information: BMC Interface Type field
var Type38_BMCInterfaceTypeMap = map[byte]string{
	0x00: "Unknown",
	0x01: "KCS: Keyboard Controller Style",
	0x02: "SMIC: Server Management Interface Chip",
	0x03: "BT: Block Transfer",
	0x04: "SSIF: SMBus System Interface",
}
func (s *Type38IPMIDeviceInformation) GetInterfaceTypeString() string {
	if v, ok := Type38_BMCInterfaceTypeMap[s.InterfaceType]; ok {
		return v
	}
	return "Unknown"
}

// System Power Supply (Type 39) (Type 39)
type Type39PowerSupply struct {
	Type byte //
	Length byte //
	Handle uint16 //
	PowerUnitGroup byte //
	Location byte //STRING
	DeviceName byte //STRING
	Manufacturer byte //STRING
	SerialNumber byte //STRING
	AssetTagNumber byte //STRING
	ModelPartNumber byte //STRING
	RevisionLevel byte //STRING
	MaxPowerCapacity uint16 //
	PowerSupplyCharacteristics uint16 //
	InputVoltageProbeHandle uint16 //
	CoolingDeviceHandle uint16 //
	InputCurrentProbeHandle uint16 //
}
func (s *Type39PowerSupply) GetLocation(strings []string) string {
	idx := int(s.Location)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type39PowerSupply) GetDeviceName(strings []string) string {
	idx := int(s.DeviceName)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type39PowerSupply) GetManufacturer(strings []string) string {
	idx := int(s.Manufacturer)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type39PowerSupply) GetSerialNumber(strings []string) string {
	idx := int(s.SerialNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type39PowerSupply) GetAssetTagNumber(strings []string) string {
	idx := int(s.AssetTagNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type39PowerSupply) GetModelPartNumber(strings []string) string {
	idx := int(s.ModelPartNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type39PowerSupply) GetRevisionLevel(strings []string) string {
	idx := int(s.RevisionLevel)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

// Additional Information (Type 40) (Type 40)
type Type40AdditionalInformation struct {
	Type byte //
	Length byte //
	Handle uint16 //
	NumberofAdditionalInformationentriesn byte //
	AdditionalInformationentries byte // Type:Varies
}

// Onboard Devices Extended Information (Type 41) structure (Type 41)
type Type41OnboardDevicesExtendedInformation struct {
	Type byte //
	Length byte //
	Handle uint16 //
	ReferenceDesignation byte //
	DeviceType byte //ENUM
	DeviceTypeInstance byte //
	SegmentGroupNumber uint16 //
	BusNumber byte //
	DeviceFunctionNumber byte //Bit Field
}
// Onboard Device Types field
var Type41_OnboardDeviceTypesMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "Video",
	0x04: "SCSI Controller",
	0x05: "Ethernet",
	0x06: "Token Ring",
	0x07: "Sound",
	0x08: "PATA Controller",
	0x09: "SATA Controller",
	0x0A: "SAS Controller",
	0x0B: "Wireless LAN",
	0x0C: "Bluetooth",
	0x0D: "WWAN",
	0x0E: "eMMC (embedded Multi-Media Controller)",
	0x0F: "NVMe Controller",
	0x10: "UFS Controller",
}
func (s *Type41OnboardDevicesExtendedInformation) GetDeviceTypeString() string {
	if v, ok := Type41_OnboardDeviceTypesMap[s.DeviceType]; ok {
		return v
	}
	return "Unknown"
}

// Management Controller Host Interface (Type 42) structure (Type 42)
type Type42ManagementControllerHostInterface struct {
	Type byte //
	Length byte //
	Handle uint16 //
	InterfaceType byte //ENUM
	InterfaceTypeSpecificDataLength byte //
	InterfaceTypeSpecificData byte // Type:N BYTEs
	NumberofProtocolRecords byte //
	ProtocolRecords byte // Type:M BYTEs
}

// TPM Device (Type 43) (Type 43)
type Type43TPMDevice struct {
	Type byte //
	Length byte //
	Handle uint16 //
	VendorID [4]byte //
	MajorSpecVersion byte //
	MinorSpecVersion byte //
	FirmwareVersion1 uint32 //
	FirmwareVersion2 uint32 //
	Description byte //STRING
	Characteristics uint64 //
	OEMdefined uint32 //
}
func (s *Type43TPMDevice) GetDescription(strings []string) string {
	idx := int(s.Description)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}
// TPM Device Characteristics
const (
	Type43_TPMDeviceCharacteristics_Bit0 = 1 << 0 // Reserved.
	Type43_TPMDeviceCharacteristics_Bit1 = 1 << 1 // Reserved.
	Type43_TPMDeviceCharacteristics_Bit2 = 1 << 2 // TPM Device Characteristics are not supported.
	Type43_TPMDeviceCharacteristics_Bit3 = 1 << 3 // Family configurable via firmware update; for example, switching between TPM 1.2 
	Type43_TPMDeviceCharacteristics_Bit4 = 1 << 4 // Family configurable via platform software support, such as Firmware Setup; for e
	Type43_TPMDeviceCharacteristics_Bit5 = 1 << 5 // Family configurable via OEM proprietary mechanism; for example, switching betwee
)

// Processor Additional Information (Type 44) (Type 44)
type Type44ProcessorAdditionalInformation struct {
	Type byte //
	Length byte //
	Handle uint16 //
	ReferencedHandle uint16 //
	ProcessorSpecificBlock byte // Type:Varies (Y)
}

// Firmware Inventory Information (Type 45) (Type 45)
type Type45FirmwareInventoryInformation struct {
	Type byte //
	Length byte //
	Handle uint16 //
	FirmwareComponentName byte //STRING
	FirmwareVersion byte //STRING
	VersionFormat byte //
	FirmwareID byte //STRING
	FirmwareIDFormat byte //
	ReleaseDate byte //STRING
	Manufacturer byte //STRING
	LowestSupportedFirmwareVersion byte //STRING
	ImageSize uint64 //
	Characteristics uint16 //Bit Field
	State byte //
	NumberofAssociatedComponentsn byte //
	AssociatedComponentHandles byte // Type:n WORDs
}
func (s *Type45FirmwareInventoryInformation) GetFirmwareComponentName(strings []string) string {
	idx := int(s.FirmwareComponentName)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type45FirmwareInventoryInformation) GetFirmwareVersion(strings []string) string {
	idx := int(s.FirmwareVersion)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type45FirmwareInventoryInformation) GetFirmwareID(strings []string) string {
	idx := int(s.FirmwareID)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type45FirmwareInventoryInformation) GetReleaseDate(strings []string) string {
	idx := int(s.ReleaseDate)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type45FirmwareInventoryInformation) GetManufacturer(strings []string) string {
	idx := int(s.Manufacturer)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *Type45FirmwareInventoryInformation) GetLowestSupportedFirmwareVersion(strings []string) string {
	idx := int(s.LowestSupportedFirmwareVersion)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}
// Firmware Inventory State Information
var Type45_FirmwareInventoryStateInformationMap = map[byte]string{
	0x01: "Other",
	0x02: "Unknown",
	0x03: "Disabled: This firmware component is disabled.",
	0x04: "Enabled: This firmware component is enabled.",
	0x05: "Absent: This firmware component is either not present or not detected",
	0x06: "StandbyOffline: This firmware is enabled but awaits an external action to activate it.",
	0x07: "StandbySpare: This firmware is part of a redundancy set and awaits a failover or other external action to activate it.",
	0x08: "UnavailableOffline: This firmware component is present but cannot be used.",
}

// String Property (Type 46) (Type 46)
type Type46StringProperty struct {
	Type byte //
	Length byte //
	Handle uint16 //
	StringPropertyID uint16 //
	StringPropertyValue byte //STRING
	Parenthandle uint16 //
}
func (s *Type46StringProperty) GetStringPropertyValue(strings []string) string {
	idx := int(s.StringPropertyValue)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}
// String Property IDs
var Type46_StringPropertyIDsMap = map[byte]string{
	0x00: "Reserved - do not use",
	0x01: "UEFI device path - string representation of a UEFI device path, as converted by EFI_DEVICE_PATH_TO_TEXT_PROTOCOL. ConvertDevicePathToText() and then converted to UTF-8",
}

// Inactive (Type 126) (Type 126)
type Type126Inactive struct {
	Type byte //
	Length byte //
	Handle uint16 //
}

// End-of-Table (Type 127) structure (Type 127)
type Type127Endoftable struct {
	Type byte //
	Length byte //
	Handle uint16 //
}

// ParsedChunk holds a parsed SMBIOS structure and its string table.
type ParsedChunk struct {
	StructType byte
	Data       interface{}
	Strings    []string
}

// ParseChunk reads an SMBIOS chunk into the appropriate typed struct.
// data is the full chunk bytes (header + fields + string table).
// structType is the SMBIOS structure type, length is the formatted area size.
func ParseChunk(structType byte, length byte, data []byte) (*ParsedChunk, error) {
	var obj interface{}
	switch structType {
	case 0:
		s := Type0PlatformFirmwareInformation{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 0: %w", err)
		}
		obj = &s
	case 1:
		s := Type1SystemInformation{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 1: %w", err)
		}
		obj = &s
	case 2:
		s := Type2BaseboardInformation{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 2: %w", err)
		}
		obj = &s
	case 3:
		s := Type3SystemEnclosure{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 3: %w", err)
		}
		obj = &s
	case 4:
		s := Type4ProcessorInformation{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 4: %w", err)
		}
		obj = &s
	case 5:
		s := Type5MemoryController{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 5: %w", err)
		}
		obj = &s
	case 6:
		s := Type6MemoryModuleConfiguration{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 6: %w", err)
		}
		obj = &s
	case 7:
		s := Type7CacheInformation{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 7: %w", err)
		}
		obj = &s
	case 8:
		s := Type8ConnectorInformation{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 8: %w", err)
		}
		obj = &s
	case 9:
		s := Type9SystemSlot{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 9: %w", err)
		}
		obj = &s
	case 10:
		s := Type10OnBoardDevicesInformation{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 10: %w", err)
		}
		obj = &s
	case 11:
		s := Type11OEMStrings{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 11: %w", err)
		}
		obj = &s
	case 12:
		s := Type12ConfigurationInformation{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 12: %w", err)
		}
		obj = &s
	case 13:
		s := Type13LanguageInformation{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 13: %w", err)
		}
		obj = &s
	case 14:
		s := Type14GroupAssociations{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 14: %w", err)
		}
		obj = &s
	case 15:
		s := Type15EventLog{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 15: %w", err)
		}
		obj = &s
	case 16:
		s := Type16PhysicalMemoryArray{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 16: %w", err)
		}
		obj = &s
	case 17:
		s := Type17MemoryDevice{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 17: %w", err)
		}
		obj = &s
	case 18:
		s := Type1832bitMemoryErrorInformation{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 18: %w", err)
		}
		obj = &s
	case 19:
		s := Type19MemoryArrayMappedAddress{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 19: %w", err)
		}
		obj = &s
	case 20:
		s := Type20MemoryDeviceMappedAddress{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 20: %w", err)
		}
		obj = &s
	case 21:
		s := Type21BuiltinPointingDevice{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 21: %w", err)
		}
		obj = &s
	case 22:
		s := Type22PortableBattery{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 22: %w", err)
		}
		obj = &s
	case 23:
		s := Type23SystemReset{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 23: %w", err)
		}
		obj = &s
	case 24:
		s := Type24HardwareSecurity{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 24: %w", err)
		}
		obj = &s
	case 25:
		s := Type25SystemPowerControls{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 25: %w", err)
		}
		obj = &s
	case 26:
		s := Type26VoltageProbe{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 26: %w", err)
		}
		obj = &s
	case 27:
		s := Type27CoolingDevice{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 27: %w", err)
		}
		obj = &s
	case 28:
		s := Type28TemperatureProbe{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 28: %w", err)
		}
		obj = &s
	case 29:
		s := Type29ElectricalCurrentProbe{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 29: %w", err)
		}
		obj = &s
	case 30:
		s := Type30OutofBandRemoteAccess{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 30: %w", err)
		}
		obj = &s
	case 32:
		s := Type32SystemBootInformationidentifier{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 32: %w", err)
		}
		obj = &s
	case 33:
		s := Type3364bitMemoryErrorInformation{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 33: %w", err)
		}
		obj = &s
	case 34:
		s := Type34ManagementDevice{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 34: %w", err)
		}
		obj = &s
	case 35:
		s := Type35ManagementDeviceComponent{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 35: %w", err)
		}
		obj = &s
	case 36:
		s := Type36ManagementDeviceThresholdData{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 36: %w", err)
		}
		obj = &s
	case 37:
		s := Type37ManagementDeviceThresholdData{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 37: %w", err)
		}
		obj = &s
	case 38:
		s := Type38IPMIDeviceInformation{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 38: %w", err)
		}
		obj = &s
	case 39:
		s := Type39PowerSupply{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 39: %w", err)
		}
		obj = &s
	case 40:
		s := Type40AdditionalInformation{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 40: %w", err)
		}
		obj = &s
	case 41:
		s := Type41OnboardDevicesExtendedInformation{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 41: %w", err)
		}
		obj = &s
	case 42:
		s := Type42ManagementControllerHostInterface{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 42: %w", err)
		}
		obj = &s
	case 43:
		s := Type43TPMDevice{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 43: %w", err)
		}
		obj = &s
	case 44:
		s := Type44ProcessorAdditionalInformation{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 44: %w", err)
		}
		obj = &s
	case 45:
		s := Type45FirmwareInventoryInformation{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 45: %w", err)
		}
		obj = &s
	case 46:
		s := Type46StringProperty{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 46: %w", err)
		}
		obj = &s
	case 126:
		s := Type126Inactive{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 126: %w", err)
		}
		obj = &s
	case 127:
		s := Type127Endoftable{}
		if err := ReadIntoStruct(data[:length], &s); err != nil {
			return nil, fmt.Errorf("type 127: %w", err)
		}
		obj = &s
	default:
		return nil, nil
	}

	strings := ParseNullTerminatedStrings(data[int(length):])
	return &ParsedChunk{StructType: structType, Data: obj, Strings: strings}, nil
}
