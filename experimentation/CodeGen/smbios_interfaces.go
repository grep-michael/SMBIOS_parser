package GeneratedCode

// BIOS
type BIOSInfo interface {
	GetVendor(strings []string) string
	GetBIOSVersion(strings []string) string
	GetBIOSStartingAddressSegment() uint16
	GetBIOSReleaseDate(strings []string) string
	GetBIOSROMSize() byte
	GetBIOSCharacteristics() uint64
	GetBIOSCharacteristicsExtensionBytes() [2]byte
	GetSystemBIOSMajorRelease() byte
	GetSystemBIOSMinorRelease() byte
	GetEmbeddedControllerFirmwareMajorRelease() byte
	GetEmbeddedControllerFirmwareMinorRelease() byte
}

type BIOS struct {
	Info    BIOSInfo
	Strings []string
}

// System
type SystemInfo interface {
	GetManufacturer(strings []string) string
	GetProductName(strings []string) string
	GetVersion(strings []string) string
	GetSerialNumber(strings []string) string
	GetUUID() [16]byte
	GetWakeupType(strings []string) string
	GetSKUNumber(strings []string) string
	GetFamily(strings []string) string
}

type System struct {
	Info    SystemInfo
	Strings []string
}

// Chassis
type ChassisInfo interface {
	GetManufacturer(strings []string) string
	GetType_0() byte
	GetVersion(strings []string) string
	GetSerialNumber(strings []string) string
	GetAssetTagNumber(strings []string) string
	GetBootupState(strings []string) string
	GetPowerSupplyState(strings []string) string
	GetThermalState(strings []string) string
	GetSecurityStatus(strings []string) string
	GetOEMdefined() uint32
	GetHeight() byte
	GetNumberofPowerCords() byte
	GetContainedElementCountn() byte
	GetContainedElementRecordLengthm() byte
	GetContainedElements() [3]byte
}

type Chassis struct {
	Info    ChassisInfo
	Strings []string
}

// Processor
type ProcessorInfo interface {
	GetSocketDesignation(strings []string) string
	GetProcessorType(strings []string) string
	GetProcessorFamily(strings []string) string
	GetProcessorManufacturer(strings []string) string
	GetProcessorID() uint64
	GetProcessorVersion(strings []string) string
	GetVoltage() byte
	GetExternalClock() uint16
	GetMaxSpeed() uint16
	GetCurrentSpeed() uint16
	GetStatus() byte
	GetProcessorUpgrade(strings []string) string
	GetL1CacheHandle() uint16
}

type Processor struct {
	Info    ProcessorInfo
	Strings []string
}

// PhysicalMemoryArray
type PhysicalMemoryArrayInfo interface {
	GetLocation(strings []string) string
	GetUse(strings []string) string
	GetMemoryErrorCorrection(strings []string) string
	GetMaximumCapacity() uint32
	GetMemoryErrorInformationHandle() uint16
	GetNumberofMemoryDevices() uint16
}

type PhysicalMemoryArray struct {
	Info    PhysicalMemoryArrayInfo
	Strings []string
}

// MemoryDevice
type MemoryDeviceInfo interface {
	GetPhysicalMemoryArrayHandle() uint16
	GetMemoryErrorInformationHandle() uint16
	GetTotalWidth() uint16
	GetDataWidth() uint16
	GetSize() uint16
	GetFormFactor(strings []string) string
	GetDeviceSet() byte
	GetDeviceLocator(strings []string) string
	GetBankLocator(strings []string) string
	GetMemoryType(strings []string) string
	GetTypeDetail() uint16
	GetSpeed() uint16
	GetManufacturer(strings []string) string
	GetSerialNumber(strings []string) string
	GetAssetTag(strings []string) string
	GetPartNumber(strings []string) string
}

type MemoryDevice struct {
	Info    MemoryDeviceInfo
	Strings []string
}

// PortableBattery
type PortableBatteryInfo interface {
	GetLocation(strings []string) string
	GetManufacturer(strings []string) string
	GetManufactureDate(strings []string) string
	GetSerialNumber(strings []string) string
	GetDeviceName(strings []string) string
	GetDeviceChemistry(strings []string) string
	GetDesignCapacity() uint16
	GetDesignVoltage() uint16
	GetSBDSVersionNumber(strings []string) string
	GetMaximumErrorinBatteryData() byte
	GetSBDSSerialNumber() uint16
	GetSBDSManufactureDate() uint16
	GetSBDSDeviceChemistry(strings []string) string
	GetDesignCapacityMultiplier() byte
	GetOEMspecific() uint32
}

type PortableBattery struct {
	Info    PortableBatteryInfo
	Strings []string
}

// HardwareSecurity
type HardwareSecurityInfo interface {
	GetHardwareSecuritySettings() byte
}

type HardwareSecurity struct {
	Info    HardwareSecurityInfo
	Strings []string
}

// PowerSupply
type PowerSupplyInfo interface {
	GetPowerUnitGroup() byte
	GetLocation(strings []string) string
	GetDeviceName(strings []string) string
	GetManufacturer(strings []string) string
	GetSerialNumber(strings []string) string
	GetAssetTagNumber(strings []string) string
	GetModelPartNumber(strings []string) string
	GetRevisionLevel(strings []string) string
	GetMaxPowerCapacity() uint16
	GetPowerSupplyCharacteristics() uint16
	GetInputVoltageProbeHandle() uint16
	GetCoolingDeviceHandle() uint16
	GetInputCurrentProbeHandle() uint16
}

type PowerSupply struct {
	Info    PowerSupplyInfo
	Strings []string
}

