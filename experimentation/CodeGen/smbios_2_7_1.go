package GeneratedCode

type S_BIOSInformationType0 struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Vendor byte //STRING
	BIOSVersion byte //STRING
	BIOSStartingAddressSegment uint16 //
	BIOSReleaseDate byte //STRING
	BIOSROMSize byte //
	BIOSCharacteristics uint64 //Bit Field
	BIOSCharacteristicsExtensionBytes Zero or more BYTES //Bit Field
	SystemBIOSMajorRelease byte //
	SystemBIOSMinorRelease byte //
	EmbeddedControllerFirmwareMajorRelease byte //
	EmbeddedControllerFirmwareMinorRelease byte //
}
type S_SystemInformationType1 struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Manufacturer byte //STRING
	ProductName byte //STRING
	Version byte //STRING
	SerialNumber byte //STRING
	UUID 16 BYTES //
	WakeupType byte //ENUM
	SKUNumber byte //STRING
	Family byte //STRING
}
type S_BaseboardorModuleInformationType2 struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Manufacturer byte //STRING
	Product byte //STRING
	Version byte //STRING
	SerialNumber byte //STRING
	AssetTag byte //STRING
	FeatureFlags byte //Bit Field
}
type S_SystemEnclosureorChassisType3 struct {
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
	ContainedElements n * m BYTES //
	SKUNumber byte //STRING
}
type S_ProcessorInformationType4 struct {
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
type S_MemoryControllerInformationType5Obsolete struct {
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
	MemoryModuleConfigurationHandles x WORDs //
	EnabledErrorCorrectingCapabilities byte //Bit Field
}
type S_MemoryModuleInformationType6Obsolete struct {
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
	_0  //
	_1  //
}
type S_CacheInformationType7 struct {
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
}
type S_PortConnectorInformationType8 struct {
	Type byte //
	Length byte //
	Handle uint16 //
	InternalReferenceDesignator byte //STRING
	InternalConnectorType byte //ENUM
	ExternalReferenceDesignator byte //STRING
	ExternalConnectorType byte //ENUM
	PortType byte //ENUM
}
type S_SystemSlotsType9 struct {
	Type byte //
	Length byte //
	Handle uint16 //
	SlotDesignation byte //STRING
	SlotType byte //ENUM
	SlotDataBusWidth byte //ENUM
	CurrentUsage byte //ENUM
	SlotLength byte //ENUM
	SlotID uint16 //
	SlotCharacteristics1 byte //Bit Field
	SlotCharacteristics2 byte //Bit Field
	SegmentGroupNumber uint16 //
	BusNumber byte //
	DeviceFunctionNumber byte //
}
type S_OnBoardDevicesInformationType10Obsolete struct {
	Type byte //
	Length byte //
	Handle uint16 //
	DevicenTypenrangesfrom1toNumberofDevices byte //
	DescriptionString byte //STRING
}
type S_OEMStringsType11 struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Count byte //
}
type S_SystemConfigurationOptionsType12 struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Count byte //
}
type S_BIOSLanguageInformationType13 struct {
	Type byte //
	Length byte //
	Handle uint16 //
	InstallableLanguages byte //
	Flags byte //Bit Field
	Reserved 15 BYTES //
	CurrentLanguage byte //STRING
}
type S_GroupAssociationsType14 struct {
	Type byte //
	Length byte //
	Handle uint16 //
	GroupName byte //STRING
	ItemType byte //
	ItemHandle uint16 //
}
type S_SystemEventLogType15 struct {
	Type byte //
	Length byte //
	Handle uint16 //
	LogAreaLength uint16 //
	LogHeaderStartOffset uint16 //
	LogDataStartOffset uint16 //
	AccessMethod byte //
	LogStatus[1] byte //
	LogChangeToken uint32 //
	AccessMethodAddress uint32 //
	LogHeaderFormat byte //ENUM
	NumberofSupportedLogTypeDescriptorsx byte //
	LengthofeachLogTypeDescriptory byte //
	ListofSupportedEventLogTypeDescriptors Varies //
}
type S_PhysicalMemoryArrayType16 struct {
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
type S_MemoryDeviceType17 struct {
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
type S_32BitMemoryErrorInformationType18 struct {
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
type S_MemoryArrayMappedAddressType19 struct {
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
type S_MemoryDeviceMappedAddressType20 struct {
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
type S_Built struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Type_0 byte //ENUM
	Interface byte //ENUM
	NumberofButtons byte //
}
type S_PortableBatteryType22 struct {
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
type S_SystemResetType23 struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Capabilities byte //
	ResetCount uint16 //
	ResetLimit uint16 //
	TimerInterval uint16 //
	Timeout uint16 //
}
type S_HardwareSecurityType24 struct {
	Type byte //
	Length byte //
	Handle uint16 //
	HardwareSecuritySettings byte //
}
type S_SystemPowerControlsType25 struct {
	Type byte //
	Length byte //
	Handle uint16 //
	NextScheduledPoweronMonth byte //
	NextScheduledPoweronDayofmonth byte //
	NextScheduledPoweronHour byte //
	NextScheduledPoweronMinute byte //
	NextScheduledPoweronSecond byte //
}
type S_VoltageProbeType26 struct {
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
type S_CoolingDeviceType27 struct {
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
type S_TemperatureProbeType28 struct {
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
type S_ElectricalCurrentProbeType29 struct {
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
type S_Out struct {
	Type byte //
	Length byte //
	Handle uint16 //
	ManufacturerName byte //STRING
	Connections byte //
}
type S_SystemBootInformationType32 struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Reserved 6 BYTEs //
	BootStatus Length-10 Bytes //
}
type S_64 struct {
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
type S_ManagementDeviceType34 struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Description byte //STRING
	Type_0 byte //
	Address uint32 //
	AddressType byte //
}
type S_ManagementDeviceComponentType35 struct {
	Type byte //
	Length byte //
	Handle uint16 //
	Description byte //STRING
	ManagementDeviceHandle uint16 //
	ComponentHandle uint16 //
	ThresholdHandle uint16 //
}
type S_ManagementDeviceThresholdDataType36 struct {
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
type S_MemoryChannelType37 struct {
	Type byte //
	Length byte //
	Handle uint16 //
	ChannelType byte //
	MaximumChannelLoad byte //
	MemoryDeviceCountn byte //
	Memory1DeviceLoad byte //
	MemoryDevice1Handle uint16 //
	MemoryDevicenLoad byte //
	MemoryDevicenHandle uint16 //
}
type S_IPMIDeviceInformationType38 struct {
	Type byte //
	Length byte //
	Handle uint16 //
	InterfaceType byte //ENUM
	IPMISpecificationRevision byte //
	I2CSlaveAddress byte //
	NVStorageDeviceAddress byte //
	BaseAddress uint64 //
	BaseAddressModifierInterruptInfo byte //
	InterruptNumber byte //
}
type S_SystemPowerSupplyType39 struct {
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
type S_AdditionalInformationType40 struct {
	Type byte //
	Length byte //
	Handle uint16 //
	NumberofAdditionalInformationentriesn byte //
	AdditionalInformationentries Varies //
}
type S_OnboardDevicesExtendedInformationType41 struct {
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
type S_ManagementControllerHostInterfaceType42 struct {
	Type byte //
	Length byte //
	Handle uint16 //
	InterfaceType byte //ENUM
	MCHostInterfaceData n BYTEs //
}
type S_InactiveType126 struct {
	Type byte //
	Length byte //
	Handle uint16 //
}
type S_EndofTableType127 struct {
	Type byte //
	Length byte //
	Handle uint16 //
}
