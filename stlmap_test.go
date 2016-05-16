package stlmap

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

func getKeys(n int) []string {
	keys := make([]string, n)
	for i := 0; i < n; i++ {
		keys[i] = fmt.Sprintf("%d", rand.Int()%100)
	}
	return keys
}

func TestNew(t *testing.T) {
	type Case struct {
		Config   *Config
		Expected int
	}
	for _, c := range []Case{
		Case{&Config{}, 1 << 10},
		Case{&Config{1}, 1 << 1},
	} {
		m := New(c.Config)
		if m.BucketSize() != c.Expected {
			t.Errorf("expected %d, but got %d", c.Expected, m.BucketSize())
		}
	}
}

func TestSetAndGetAndDelete(t *testing.T) {
	m := New(&Config{})
	m.Set("key", 1)
	res := m.Get("key")
	if res != 1 {
		t.Error("res must be 1")
	}
	m.Delete("key")
	res = m.Get("key")
	if res != nil {
		t.Error("res must be nil")
	}
}

func TestConcurrent(t *testing.T) {
	do := func(wg *sync.WaitGroup, m *StripedLockedMap, keys []string) interface{} {
		key := keys[rand.Int()%len(keys)]
		m.Set(key, 1)
		res := m.Get(key)
		m.Delete(key)
		wg.Done()
		return res
	}
	m := New(&Config{})
	wg := &sync.WaitGroup{}
	keys := getKeys(1000)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go do(wg, m, keys)
	}
	wg.Wait()
}

func BenchmarkStlmapConcurrent(b *testing.B) {
	do := func(m *StripedLockedMap, keys []string) interface{} {
		key := keys[rand.Int()%len(keys)]
		m.Set(key, 1)
		res := m.Get(key)
		m.Delete(key)
		return res
	}
	m := New(&Config{})
	keys := getKeys(1000)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			do(m, keys)
		}
	})
}

func BenchmarkDefaultMapConcurrent(b *testing.B) {
	do := func(mu *sync.RWMutex, mp map[string]interface{}, keys []string) interface{} {
		key := keys[rand.Int()%len(keys)]
		mu.Lock()
		mp[key] = 1
		mu.Unlock()
		mu.RLock()
		res := mp[key]
		mu.RUnlock()
		mu.Lock()
		delete(mp, key)
		mu.Unlock()
		return res
	}
	mu := &sync.RWMutex{}
	mp := make(map[string]interface{})
	keys := getKeys(1000)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			do(mu, mp, keys)
		}
	})
}

func BenchmarkStlmapConcurrentSet(b *testing.B) {
	do := func(m *StripedLockedMap, keys []string) {
		key := keys[rand.Int()%len(keys)]
		m.Set(key, 1)
	}
	m := New(&Config{})
	keys := getKeys(1000)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			do(m, keys)
		}
	})
}

func BenchmarkDefaultMapConcurrentSet(b *testing.B) {
	do := func(mu *sync.RWMutex, mp map[string]interface{}, keys []string) {
		key := keys[rand.Int()%len(keys)]
		mu.Lock()
		mp[key] = 1
		mu.Unlock()
	}
	mu := &sync.RWMutex{}
	mp := make(map[string]interface{})
	keys := getKeys(1000)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			do(mu, mp, keys)
		}
	})
}

func BenchmarkStlmapConcurrentGet(b *testing.B) {
	do := func(mp *StripedLockedMap, keys []string) interface{} {
		key := keys[rand.Int()%len(keys)]
		return mp.Get(key)
	}
	mp := New(&Config{})
	keys := getKeys(1000)
	for _, k := range keys {
		mp.Set(k, 1)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			do(mp, keys)
		}
	})
}

func BenchmarkDefaultMapConcurrentGet(b *testing.B) {
	do := func(mu *sync.RWMutex, mp map[string]interface{}, keys []string) interface{} {
		key := keys[rand.Int()%len(keys)]
		mu.RLock()
		res := mp[key]
		mu.RUnlock()
		return res
	}
	mu := &sync.RWMutex{}
	mp := make(map[string]interface{})
	keys := getKeys(1000)
	for _, k := range keys {
		mp[k] = 1
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			do(mu, mp, keys)
		}
	})
}

func BenchmarkStlmapConcurrentDelete(b *testing.B) {
	do := func(mp *StripedLockedMap, keys []string) {
		key := keys[rand.Int()%len(keys)]
		mp.Delete(key)
	}
	mp := New(&Config{})
	keys := getKeys(1000)
	for _, k := range keys {
		mp.Set(k, 1)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			do(mp, keys)
		}
	})
}

func BenchmarkDefaultMapConcurrentDelete(b *testing.B) {
	do := func(mu *sync.RWMutex, mp map[string]interface{}, keys []string) {
		key := keys[rand.Int()%len(keys)]
		mu.Lock()
		delete(mp, key)
		mu.Unlock()
	}
	mu := &sync.RWMutex{}
	mp := make(map[string]interface{})
	keys := getKeys(1000)
	for _, k := range keys {
		mp[k] = 1
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			do(mu, mp, keys)
		}
	})
}
