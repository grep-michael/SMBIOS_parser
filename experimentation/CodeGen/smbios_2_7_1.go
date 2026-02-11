package GeneratedCode

type S_0_BIOSInformationIndicator struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Vendor byte //STRING
	BIOSVersion byte //STRING
	BIOSStartingAddressSegment uint16 //
	BIOSReleaseDate byte //STRING
	BIOSROMSize byte //
	BIOSCharacteristics uint64 //Bit Field
	BIOSCharacteristicsExtensionBytes interface{} //Bit Field Type:Zero or more BYTES
	SystemBIOSMajorRelease byte //
	SystemBIOSMinorRelease byte //
	EmbeddedControllerFirmwareMajorRelease byte //
	EmbeddedControllerFirmwareMinorRelease byte //
}
type S_1_SystemInformationIndicator struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Manufacturer byte //STRING
	ProductName byte //STRING
	Version byte //STRING
	SerialNumber byte //STRING
	UUID interface{} // Type:16 BYTES
	WakeupType byte //ENUM
	SKUNumber byte //STRING
	Family byte //STRING
}
type S_3_SystemEnclosureIndicator struct {
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
type S_4_ProcessorInformationIndicator struct {
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
}
type S_16_PhysicalMemoryArraytype struct {
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
type S_17_MemoryDevicetype struct {
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
	ConfiguredMemoryClockSpeed uint16 //
}
type S_22_PortableBatteryindicator struct {
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
type S_24_HardwareSecurityindicator struct {
	Type byte //
	Length byte //
	Handle uint16 //
	HardwareSecuritySettings byte //
}
type S_39_PowerSupplyStructureindicator struct {
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
type S_126_Inactivestructureindicator struct {
	Type byte //
	Length byte //
	Handle uint16 //
}
type S_127_Endoftableindicator struct {
	Type byte //
	Length byte //
	Handle uint16 //
}
