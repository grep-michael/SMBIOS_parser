package structures

import (
	"log"
	"sync"

	structuretypes "github.com/grep-michael/SMBIOS_parser/SMBiosLib/StructureTypes"
)

/*
	Registry of factory functions to create new SMBio Structures
*/

type Factory func() SMBiosStruct

type Registry struct {
	newFuncs map[structuretypes.StructureType]Factory
}

var theRegistry = &Registry{
	newFuncs: make(map[structuretypes.StructureType]Factory),
}
var (
	mu sync.Mutex
)

func Register(typeInt structuretypes.StructureType, newFunc Factory) {
	mu.Lock()
	defer mu.Unlock()
	_, exists := theRegistry.newFuncs[typeInt]
	if exists {
		log.Printf("attempted to add interface for already existing type %d\n", typeInt)
		return
	}
	theRegistry.newFuncs[typeInt] = newFunc
}

func GetStructForType(typeInt structuretypes.StructureType) (SMBiosStruct, bool) {
	newFunc, exists := theRegistry.newFuncs[typeInt]
	if !exists {
		return nil, false
	}
	return newFunc(), true
}
