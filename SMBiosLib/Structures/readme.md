# FOR SMBIOS v3.3

this folder contains the structs we use to parse bytes into golang structs, and then parse those into actual objects 

#### files are named by what struct type they are, for example type four is cpu info, see DSP0134_3.3.0.pdf page 28, or header.go

```Golang
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
```
Currently supported Types

### checklist
 * [x] TypeBiosInfo
 * [ ] TypeSystemInfo
 * [ ] TypeSystemEnclosure
 * [x] TypeProcessorInfo
 * [ ] TypeCacheInfo
 * [ ] TypeSystemSlots
 * [ ] TypePhysicalMemoryArray
 * [ ] TypeMemoryDevice
 * [ ] TypeMemoryArrayMappedAddress
 * [ ] TypeMemoryDeviceMappedAddress
 * [ ] TypeSystemBootInfo