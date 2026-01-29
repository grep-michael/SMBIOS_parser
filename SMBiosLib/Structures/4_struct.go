package structures

import typemap "github.com/grep-michael/SMBIOS_parser/SMBiosLib/TypeMap"

//ProcessUpgrade is Basically just socket type

type ProcessorInfo struct {
	StructureHeader
	SocketDesignation        byte
	ProcessorType            byte //enum
	ProcessorFamily          byte //enum
	ProcessorManufacturer    byte
	ProcessorID              uint64
	ProcessorVersion         byte
	Voltage                  byte
	ExternalClock            uint16
	MaxSpeed                 uint16 //in MHz
	CurrentSpeed             uint16
	Status                   byte
	ProcessorUpgrade         byte //enum
	CacheOneHandler          uint16
	CacheTwoHandler          uint16
	CacheThreeHandler        uint16
	SerialNum                byte
	AssetTag                 byte
	PartNumber               byte
	CoreCount                byte
	CoreEnabled              byte
	ThreadCount              byte
	ProcessorCharacteristics uint16
	ProcessorFamilyTwo       uint16
	CoreCountTwo             byte
	CoreEnabledTwo           byte
	ThreadCountTwo           byte
}

func NewProssorInfo() *ProcessorInfo {
	return &ProcessorInfo{}
}

func init() {
	typemap.Register(4, NewProssorInfo)
}
