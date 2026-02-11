package GeneratedCode

type SMB3_3_0_S0_biosinformationindicator struct {
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
type SMB3_3_0_S1_systeminformationindicator struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Manufacturer byte //STRING
	ProductName byte //STRING
	Version byte //STRING
	SerialNumber byte //STRING
	UUID interface{} // Type:16 BYTEs
	WakeupType byte //ENUM
	SKUNumber byte //STRING
	Family byte //STRING
}
type SMB3_3_0_S3_systemenclosureindicator struct {
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
	ContainedElements interface{} // Type:n * m BYTES
	SKUNumber byte //STRING
}
type SMB3_3_0_S4_processorinformationindicator struct {
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
}
type SMB3_3_0_S16_physicalmemoryarraytype struct {
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
type SMB3_3_0_S17_memorydevicetype struct {
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
}
type SMB3_3_0_S22_portablebatteryindicator struct {
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
type SMB3_3_0_S24_hardwaresecurityindicator struct {
	Type byte //
	Length byte //
	Handle uint16 //
	HardwareSecuritySettings byte //
}
type SMB3_3_0_S39_powersupplystructureindicator struct {
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
type SMB3_3_0_S126_inactivestructureindicator struct {
	Type byte //
	Length byte //
	Handle uint16 //
}
type SMB3_3_0_S127_endoftableindicator struct {
	Type byte //
	Length byte //
	Handle uint16 //
}
