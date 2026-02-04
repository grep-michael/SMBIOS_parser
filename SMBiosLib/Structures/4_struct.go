package structures

import structuretypes "github.com/grep-michael/SMBIOS_parser/SMBiosLib/StructureTypes"

//ProcessUpgrade is Basically just socket type

type ProcessorInfoFixed struct {
	StructureHeader
	SocketDesignation        ByteString
	ProcessorType            ByteEnum
	ProcessorFamily          ByteEnum
	ProcessorManufacturer    ByteString
	ProcessorID              QWord
	ProcessorVersion         ByteString
	Voltage                  byte
	ExternalClock            Word
	MaxSpeed                 Word //in MHz
	CurrentSpeed             Word
	Status                   byte
	ProcessorUpgrade         ByteEnum
	CacheOneHandler          Word
	CacheTwoHandler          Word
	CacheThreeHandler        Word
	SerialNum                ByteStringIndex
	AssetTag                 ByteStringIndex
	PartNumber               ByteStringIndex
	CoreCount                byte
	CoreEnabled              byte
	ThreadCount              byte
	ProcessorCharacteristics Word
	ProcessorFamilyTwo       Word
	CoreCountTwo             byte
	CoreEnabledTwo           byte
	ThreadCountTwo           byte
}

type ProcessorInfo = GenericStruct[ProcessorInfoFixed]

func NewProssorInfo() SMBiosStruct {
	return &ProcessorInfo{}
}

func init() {
	Register(structuretypes.TypeProcessorInfo, NewProssorInfo)
}
