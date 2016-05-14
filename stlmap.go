package stlmap

import (
	"sync"
)

const (
	offset32 = uint32(2166136261)
	prime32  = uint32(16777619)
)

type StripedLockedMap struct {
	backends []*Backend
	mod      uint32
}

type Backend struct {
	mu *sync.RWMutex
	mp map[string]interface{}
}

func New() *StripedLockedMap {
	exp := uint(3)
	backends := make([]*Backend, 1<<exp)
	for i := 0; i < (1 << exp); i++ {
		backends[i] = &Backend{
			mu: &sync.RWMutex{},
			mp: make(map[string]interface{}),
		}
	}
	return &StripedLockedMap{
		backends: backends,
		mod:      (1 << exp) - 1,
	}
}

func (smap *StripedLockedMap) backend(key string) *Backend {
	hash := offset32
	for i := 0; i < len(key); i++ {
		hash ^= uint32(key[i])
		hash *= prime32
	}
	return smap.backends[hash&smap.mod]
}

func (smap *StripedLockedMap) Set(key string, value interface{}) {
	b := smap.backend(key)
	b.mu.Lock()
	b.mp[key] = value
	b.mu.Unlock()
}

func (smap *StripedLockedMap) Get(key string) interface{} {
	b := smap.backend(key)
	b.mu.RLock()
	res := b.mp[key]
	b.mu.RUnlock()
	return res
}

func (smap *StripedLockedMap) Delete(key string) {
	b := smap.backend(key)
	b.mu.Lock()
	delete(b.mp, key)
	b.mu.Unlock()
}
