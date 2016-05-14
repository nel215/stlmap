package stlmap

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

func TestSetAndGetAndDelete(t *testing.T) {
	m := New()
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
	do := func(wg *sync.WaitGroup, m *StripedLockedMap) {
		key := fmt.Sprintf("%d", rand.Int()%100)
		m.Set(key, 1)
		m.Get(key)
		m.Delete(key)
		wg.Done()
	}
	m := New()
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go do(wg, m)
	}
	wg.Wait()
}
