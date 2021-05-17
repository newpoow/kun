package maps

import (
	"sync"

	"github.com/yaoapp/kun/interfaces"
)

// MapStrAnySync type of sync.Map
type MapStrAnySync struct {
	*sync.Map
}

// MakeStrSync create a new instance
func MakeStrSync() MapStrAnySync {
	return MapStrAnySync{
		Map: &sync.Map{},
	}
}

// MakeStrAnySync create a new instance
func MakeStrAnySync() MapStrAnySync {
	return MapStrAnySync{
		Map: &sync.Map{},
	}
}

// Set set the value for a key
func (m MapStrAnySync) Set(key string, value interface{}) {
	m.Store(key, value)
}

// Get get the value of the given key
func (m MapStrAnySync) Get(key string) interface{} {
	value, has := m.Load(key)
	if has {
		return value
	}
	return nil
}

// Del deletes the value for a key.
func (m MapStrAnySync) Del(key string) {
	m.Delete(key)
}

// GetOrSet returns the existing value for the key if present. Otherwise, it stores and returns the given value.
func (m MapStrAnySync) GetOrSet(key string, value interface{}) interface{} {
	value, load := m.LoadOrStore(key, value)
	if load {
		return value
	}

	return value
}

// GetAndDel deletes the value for a key, returning the previous value if any. The loaded result reports whether the key was present.
func (m MapStrAnySync) GetAndDel(key string) interface{} {
	value := m.Get(key)
	m.Del(key)
	return value
}

// Range calls f sequentially for each key and value present in the map. If f returns false, range stops the iteration.
func (m MapStrAnySync) Range(cb func(key string, value interface{}) bool) {
	m.Map.Range(func(key, value interface{}) bool {
		keyStr := key.(string)
		return cb(keyStr, value)
	})
}

//IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.
func (m MapStrAnySync) IsEmpty() bool {
	has := false
	m.Range(func(key string, value interface{}) bool {
		has = true
		return false
	})
	return has
}

// Merge merges hash maps
func (m MapStrAnySync) Merge(maps ...interfaces.MapStr) {
	for _, new := range maps {
		new.Range(func(key string, value interface{}) bool {
			m.Set(key, value)
			return true
		})
	}
}
