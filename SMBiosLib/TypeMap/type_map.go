package typemap

import (
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
	theRegistry.newFuncs[typeInt] = func() interface{} { return newFunc() }
}

func GetStructForType(id int) (interface{}, bool) {
	newFunc, exists := theRegistry.newFuncs[id]
	if !exists {
		return nil, false
	}
	return newFunc(), true
}
