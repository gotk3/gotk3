package callback

import (
	"sync"

	"github.com/gotk3/gotk3/internal/slab"
)

var (
	mutex    sync.RWMutex
	registry slab.Slab
)

func Assign(callback interface{}) uintptr {
	mutex.Lock()
	defer mutex.Unlock()

	return registry.Put(callback)
}

func Get(ptr uintptr) interface{} {
	mutex.RLock()
	defer mutex.RUnlock()

	return registry.Get(ptr)
}

func Delete(ptr uintptr) {
	GetAndDelete(ptr)
}

func GetAndDelete(ptr uintptr) interface{} {
	mutex.Lock()
	defer mutex.Unlock()

	return registry.Pop(ptr)
}
