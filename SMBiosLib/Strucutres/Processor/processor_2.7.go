package processor

type Processor2_7Fixed struct {
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
