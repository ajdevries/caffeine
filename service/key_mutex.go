package service

import "sync"

// KeyedMutex for locking per key
type KeyedMutex struct {
	mutexes sync.Map
}

// Lock for a specific key
func (m *KeyedMutex) Lock(key string) func() {
	value, _ := m.mutexes.LoadOrStore(key, &sync.Mutex{})
	mtx := value.(*sync.Mutex)
	mtx.Lock()

	return func() { mtx.Unlock() }
}
