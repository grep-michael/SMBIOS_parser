package structures

import structuretypes "github.com/grep-michael/SMBIOS_parser/SMBiosLib/StructureTypes"

//ProcessUpgrade is Basically just socket type

type ProcessorInfoFixed struct {
	StructureHeader
	SocketDesignation        ByteString
	ProcessorType            ByteEnum
	ProcessorFamily          ByteEnum
	ProcessorManufacturer    ByteString
	ProcessorID              uint64
	ProcessorVersion         ByteString
	Voltage                  byte
	ExternalClock            uint16
	MaxSpeed                 uint16 //in MHz
	CurrentSpeed             uint16
	Status                   byte
	ProcessorUpgrade         ByteEnum
	CacheOneHandler          uint16
	CacheTwoHandler          uint16
	CacheThreeHandler        uint16
	SerialNum                ByteStringIndex
	AssetTag                 ByteStringIndex
	PartNumber               ByteStringIndex
	CoreCount                byte
	CoreEnabled              byte
	ThreadCount              byte
	ProcessorCharacteristics uint16
	ProcessorFamilyTwo       uint16
	CoreCountTwo             byte
	CoreEnabledTwo           byte
	ThreadCountTwo           byte
}

type ProcessorInfo = GenericStruct[BaseboardInfoFixed]

func NewProssorInfo() SMBiosStruct {
	return &ProcessorInfo{}
}

func init() {
	Register(structuretypes.TypeProcessorInfo, NewProssorInfo)
}
