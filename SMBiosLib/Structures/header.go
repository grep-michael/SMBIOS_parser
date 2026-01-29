package structures

type StructureType byte

const (
	TypeBiosInfo                  StructureType = 0
	TypeSystemInfo                StructureType = 1
	TypeSystemEnclosure           StructureType = 3
	TypeProcessorInfo             StructureType = 4
	TypeCacheInfo                 StructureType = 7
	TypeSystemSlots               StructureType = 9
	TypePhysicalMemoryArray       StructureType = 16
	TypeMemoryDevice              StructureType = 17
	TypeMemoryArrayMappedAddress  StructureType = 19
	TypeMemoryDeviceMappedAddress StructureType = 20 //v2.2
	TypeSystemBootInfo            StructureType = 32 //v3.3
)

type StructureHeader struct {
	Type   StructureType
	Length byte // depending on the struct there will be a default length assuming no extensions bytes are used
	Handle uint16
}

type TestStruct struct {
	StructureHeader
}
