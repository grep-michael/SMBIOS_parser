package GeneratedCode

type SMB3_7_1_S0_biosinformationindicator struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Vendor byte //STRING
	BIOSVersion byte //STRING
	BIOSStartingAddressSegment uint16 //
	BIOSReleaseDate byte //STRING
	BIOSROMSize byte //
	BIOSCharacteristics uint64 //Bit Field
	BIOSCharacteristicsExtensionBytes [2]byte //Bit Field Type:Zero or more BYTEs
	SystemBIOSMajorRelease byte //
	SystemBIOSMinorRelease byte //
	EmbeddedControllerFirmwareMajorRelease byte //
	EmbeddedControllerFirmwareMinorRelease byte //
	ExtendedBIOSROMSize uint16 //Bit Field
}
type SMB3_7_1_S1_systeminformationindicator struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Manufacturer byte //STRING
	ProductName byte //STRING
	Version byte //STRING
	SerialNumber byte //STRING
	UUID [16]byte // Type:16 BYTES
	WakeupType byte //ENUM
	SKUNumber byte //STRING
	Family byte //STRING
}
type SMB3_7_1_S3_systemenclosureindicator struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Manufacturer byte //STRING
	Type_0 byte //
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
	ContainedElements [3]byte // Type:n * m BYTES
	SKUNumber byte //STRING
}
type SMB3_7_1_S4_processorinformationindicator struct {
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
}
type SMB3_7_1_S9_systemslotstructureindicator struct {
	Type byte //
	Length byte //
}
type SMB3_7_1_S16_physicalmemoryarraytype struct {
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
type SMB3_7_1_S17_memorydevicetype struct {
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
type SMB3_7_1_S22_portablebatteryindicator struct {
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
type SMB3_7_1_S24_hardwaresecurityindicator struct {
	Type byte //
	Length byte //
	Handle uint16 //
	HardwareSecuritySettings byte //
}
type SMB3_7_1_S39_powersupplystructureindicator struct {
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
type SMB3_7_1_S126_inactivestructureindicator struct {
	Type byte //
	Length byte //
	Handle uint16 //
}
type SMB3_7_1_S127_endoftableindicator struct {
	Type byte //
	Length byte //
	Handle uint16 //
}

// -- BIOS interface methods --

func (s *SMB3_7_1_S0_biosinformationindicator) GetVendor(strings []string) string {
	idx := int(s.Vendor)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S0_biosinformationindicator) GetBIOSStartingAddressSegment() uint16 {
	return s.BIOSStartingAddressSegment
}

func (s *SMB3_7_1_S0_biosinformationindicator) GetEmbeddedControllerFirmwareMajorRelease() byte {
	return s.EmbeddedControllerFirmwareMajorRelease
}

func (s *SMB3_7_1_S0_biosinformationindicator) GetEmbeddedControllerFirmwareMinorRelease() byte {
	return s.EmbeddedControllerFirmwareMinorRelease
}


// -- System interface methods --

func (s *SMB3_7_1_S1_systeminformationindicator) GetManufacturer(strings []string) string {
	idx := int(s.Manufacturer)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S1_systeminformationindicator) GetProductName(strings []string) string {
	idx := int(s.ProductName)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S1_systeminformationindicator) GetVersion(strings []string) string {
	idx := int(s.Version)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S1_systeminformationindicator) GetSerialNumber(strings []string) string {
	idx := int(s.SerialNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S1_systeminformationindicator) GetUUID() [16]byte {
	return s.UUID
}

func (s *SMB3_7_1_S1_systeminformationindicator) GetWakeupType(strings []string) string {
	idx := int(s.WakeupType)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S1_systeminformationindicator) GetSKUNumber(strings []string) string {
	idx := int(s.SKUNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S1_systeminformationindicator) GetFamily(strings []string) string {
	idx := int(s.Family)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}


// -- Chassis interface methods --

func (s *SMB3_7_1_S3_systemenclosureindicator) GetManufacturer(strings []string) string {
	idx := int(s.Manufacturer)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S3_systemenclosureindicator) GetType_0() byte {
	return s.Type_0
}

func (s *SMB3_7_1_S3_systemenclosureindicator) GetVersion(strings []string) string {
	idx := int(s.Version)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S3_systemenclosureindicator) GetSerialNumber(strings []string) string {
	idx := int(s.SerialNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S3_systemenclosureindicator) GetAssetTagNumber(strings []string) string {
	idx := int(s.AssetTagNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S3_systemenclosureindicator) GetBootupState(strings []string) string {
	idx := int(s.BootupState)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S3_systemenclosureindicator) GetPowerSupplyState(strings []string) string {
	idx := int(s.PowerSupplyState)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S3_systemenclosureindicator) GetThermalState(strings []string) string {
	idx := int(s.ThermalState)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S3_systemenclosureindicator) GetSecurityStatus(strings []string) string {
	idx := int(s.SecurityStatus)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S3_systemenclosureindicator) GetOEMdefined() uint32 {
	return s.OEMdefined
}

func (s *SMB3_7_1_S3_systemenclosureindicator) GetHeight() byte {
	return s.Height
}

func (s *SMB3_7_1_S3_systemenclosureindicator) GetNumberofPowerCords() byte {
	return s.NumberofPowerCords
}

func (s *SMB3_7_1_S3_systemenclosureindicator) GetContainedElementCountn() byte {
	return s.ContainedElementCountn
}

func (s *SMB3_7_1_S3_systemenclosureindicator) GetContainedElementRecordLengthm() byte {
	return s.ContainedElementRecordLengthm
}

func (s *SMB3_7_1_S3_systemenclosureindicator) GetContainedElements() [3]byte {
	return s.ContainedElements
}


// -- Processor interface methods --

func (s *SMB3_7_1_S4_processorinformationindicator) GetSocketDesignation(strings []string) string {
	idx := int(s.SocketDesignation)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S4_processorinformationindicator) GetProcessorType(strings []string) string {
	idx := int(s.ProcessorType)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S4_processorinformationindicator) GetProcessorFamily(strings []string) string {
	idx := int(s.ProcessorFamily)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S4_processorinformationindicator) GetProcessorManufacturer(strings []string) string {
	idx := int(s.ProcessorManufacturer)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S4_processorinformationindicator) GetProcessorID() uint64 {
	return s.ProcessorID
}

func (s *SMB3_7_1_S4_processorinformationindicator) GetProcessorVersion(strings []string) string {
	idx := int(s.ProcessorVersion)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S4_processorinformationindicator) GetVoltage() byte {
	return s.Voltage
}

func (s *SMB3_7_1_S4_processorinformationindicator) GetExternalClock() uint16 {
	return s.ExternalClock
}

func (s *SMB3_7_1_S4_processorinformationindicator) GetMaxSpeed() uint16 {
	return s.MaxSpeed
}

func (s *SMB3_7_1_S4_processorinformationindicator) GetCurrentSpeed() uint16 {
	return s.CurrentSpeed
}

func (s *SMB3_7_1_S4_processorinformationindicator) GetStatus() byte {
	return s.Status
}

func (s *SMB3_7_1_S4_processorinformationindicator) GetProcessorUpgrade(strings []string) string {
	idx := int(s.ProcessorUpgrade)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S4_processorinformationindicator) GetL1CacheHandle() uint16 {
	return s.L1CacheHandle
}


// -- SystemSlot interface methods --



// -- PhysicalMemoryArray interface methods --

func (s *SMB3_7_1_S16_physicalmemoryarraytype) GetLocation(strings []string) string {
	idx := int(s.Location)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S16_physicalmemoryarraytype) GetUse(strings []string) string {
	idx := int(s.Use)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S16_physicalmemoryarraytype) GetMemoryErrorCorrection(strings []string) string {
	idx := int(s.MemoryErrorCorrection)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S16_physicalmemoryarraytype) GetMaximumCapacity() uint32 {
	return s.MaximumCapacity
}

func (s *SMB3_7_1_S16_physicalmemoryarraytype) GetMemoryErrorInformationHandle() uint16 {
	return s.MemoryErrorInformationHandle
}

func (s *SMB3_7_1_S16_physicalmemoryarraytype) GetNumberofMemoryDevices() uint16 {
	return s.NumberofMemoryDevices
}


// -- MemoryDevice interface methods --

func (s *SMB3_7_1_S17_memorydevicetype) GetPhysicalMemoryArrayHandle() uint16 {
	return s.PhysicalMemoryArrayHandle
}

func (s *SMB3_7_1_S17_memorydevicetype) GetMemoryErrorInformationHandle() uint16 {
	return s.MemoryErrorInformationHandle
}

func (s *SMB3_7_1_S17_memorydevicetype) GetTotalWidth() uint16 {
	return s.TotalWidth
}

func (s *SMB3_7_1_S17_memorydevicetype) GetDataWidth() uint16 {
	return s.DataWidth
}

func (s *SMB3_7_1_S17_memorydevicetype) GetSize() uint16 {
	return s.Size
}

func (s *SMB3_7_1_S17_memorydevicetype) GetFormFactor(strings []string) string {
	idx := int(s.FormFactor)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S17_memorydevicetype) GetDeviceSet() byte {
	return s.DeviceSet
}

func (s *SMB3_7_1_S17_memorydevicetype) GetDeviceLocator(strings []string) string {
	idx := int(s.DeviceLocator)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S17_memorydevicetype) GetBankLocator(strings []string) string {
	idx := int(s.BankLocator)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S17_memorydevicetype) GetMemoryType(strings []string) string {
	idx := int(s.MemoryType)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S17_memorydevicetype) GetTypeDetail() uint16 {
	return s.TypeDetail
}

func (s *SMB3_7_1_S17_memorydevicetype) GetSpeed() uint16 {
	return s.Speed
}

func (s *SMB3_7_1_S17_memorydevicetype) GetManufacturer(strings []string) string {
	idx := int(s.Manufacturer)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S17_memorydevicetype) GetSerialNumber(strings []string) string {
	idx := int(s.SerialNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S17_memorydevicetype) GetAssetTag(strings []string) string {
	idx := int(s.AssetTag)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S17_memorydevicetype) GetPartNumber(strings []string) string {
	idx := int(s.PartNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}


// -- PortableBattery interface methods --

func (s *SMB3_7_1_S22_portablebatteryindicator) GetLocation(strings []string) string {
	idx := int(s.Location)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S22_portablebatteryindicator) GetManufacturer(strings []string) string {
	idx := int(s.Manufacturer)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S22_portablebatteryindicator) GetManufactureDate(strings []string) string {
	idx := int(s.ManufactureDate)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S22_portablebatteryindicator) GetSerialNumber(strings []string) string {
	idx := int(s.SerialNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S22_portablebatteryindicator) GetDeviceName(strings []string) string {
	idx := int(s.DeviceName)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S22_portablebatteryindicator) GetDeviceChemistry(strings []string) string {
	idx := int(s.DeviceChemistry)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S22_portablebatteryindicator) GetDesignCapacity() uint16 {
	return s.DesignCapacity
}

func (s *SMB3_7_1_S22_portablebatteryindicator) GetDesignVoltage() uint16 {
	return s.DesignVoltage
}

func (s *SMB3_7_1_S22_portablebatteryindicator) GetSBDSVersionNumber(strings []string) string {
	idx := int(s.SBDSVersionNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S22_portablebatteryindicator) GetMaximumErrorinBatteryData() byte {
	return s.MaximumErrorinBatteryData
}

func (s *SMB3_7_1_S22_portablebatteryindicator) GetSBDSSerialNumber() uint16 {
	return s.SBDSSerialNumber
}

func (s *SMB3_7_1_S22_portablebatteryindicator) GetSBDSManufactureDate() uint16 {
	return s.SBDSManufactureDate
}

func (s *SMB3_7_1_S22_portablebatteryindicator) GetSBDSDeviceChemistry(strings []string) string {
	idx := int(s.SBDSDeviceChemistry)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S22_portablebatteryindicator) GetDesignCapacityMultiplier() byte {
	return s.DesignCapacityMultiplier
}

func (s *SMB3_7_1_S22_portablebatteryindicator) GetOEMspecific() uint32 {
	return s.OEMspecific
}


// -- HardwareSecurity interface methods --

func (s *SMB3_7_1_S24_hardwaresecurityindicator) GetHardwareSecuritySettings() byte {
	return s.HardwareSecuritySettings
}


// -- PowerSupply interface methods --

func (s *SMB3_7_1_S39_powersupplystructureindicator) GetPowerUnitGroup() byte {
	return s.PowerUnitGroup
}

func (s *SMB3_7_1_S39_powersupplystructureindicator) GetLocation(strings []string) string {
	idx := int(s.Location)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S39_powersupplystructureindicator) GetDeviceName(strings []string) string {
	idx := int(s.DeviceName)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S39_powersupplystructureindicator) GetManufacturer(strings []string) string {
	idx := int(s.Manufacturer)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S39_powersupplystructureindicator) GetSerialNumber(strings []string) string {
	idx := int(s.SerialNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S39_powersupplystructureindicator) GetAssetTagNumber(strings []string) string {
	idx := int(s.AssetTagNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S39_powersupplystructureindicator) GetModelPartNumber(strings []string) string {
	idx := int(s.ModelPartNumber)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S39_powersupplystructureindicator) GetRevisionLevel(strings []string) string {
	idx := int(s.RevisionLevel)
	if idx > 0 && idx <= len(strings) {
		return strings[idx-1]
	}
	return ""
}

func (s *SMB3_7_1_S39_powersupplystructureindicator) GetMaxPowerCapacity() uint16 {
	return s.MaxPowerCapacity
}

func (s *SMB3_7_1_S39_powersupplystructureindicator) GetPowerSupplyCharacteristics() uint16 {
	return s.PowerSupplyCharacteristics
}

func (s *SMB3_7_1_S39_powersupplystructureindicator) GetInputVoltageProbeHandle() uint16 {
	return s.InputVoltageProbeHandle
}

func (s *SMB3_7_1_S39_powersupplystructureindicator) GetCoolingDeviceHandle() uint16 {
	return s.CoolingDeviceHandle
}

func (s *SMB3_7_1_S39_powersupplystructureindicator) GetInputCurrentProbeHandle() uint16 {
	return s.InputCurrentProbeHandle
}


// -- Inactive interface methods --



// -- EndOfTable interface methods --


