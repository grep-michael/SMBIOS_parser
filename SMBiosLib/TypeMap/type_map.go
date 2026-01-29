package typemap

import (
	"log"
	"sync"
)

type TypedFunc[T any] func() *T

type Registry struct {
	newFuncs map[int]func() interface{}
}

var theRegistry = &Registry{
	newFuncs: make(map[int]func() interface{}),
}
var (
	mu sync.Mutex
)

func Register[T any](typeInt int, newFunc TypedFunc[T]) {
	mu.Lock()
	defer mu.Unlock()
	_, exists := theRegistry.newFuncs[typeInt]
	if exists {
		log.Printf("attempted to add interface for already existing type %d\n", typeInt)
		return
	}
	log.Printf("Registering type %d\n", typeInt)
	theRegistry.newFuncs[typeInt] = func() interface{} { return newFunc() }
}

func GetStructForType(typeInt int) (interface{}, bool) {
	newFunc, exists := theRegistry.newFuncs[typeInt]
	if !exists {
		return nil, false
	}
	return newFunc(), true
}
