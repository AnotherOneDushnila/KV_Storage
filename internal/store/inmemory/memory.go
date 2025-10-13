package memory

import (
	"fmt"
	"sync"
) 

type MemoryStore struct {
	mu   sync.RWMutex
	data map[string]map[string][]byte
}


func New() *MemoryStore {
	return &MemoryStore{data: make(map[string]map[string][]byte)}
}


func (m *MemoryStore) Put(collection, key string, value []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.data[collection][key]; !exists {
		return fmt.Errorf("key already exists")
	}

	if _, err := m.data[collection]; !err {
		m.data[collection] = make(map[string][]byte)
	}
	m.data[collection][key] = value

	return nil
}


func (m *MemoryStore) Get(collection, key string) ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	coll, err := m.data[collection]

	if !err {
		return nil, fmt.Errorf("collection not found")
	}

	value, err := coll[key]

	if !err {
		return nil, fmt.Errorf("value not found")
	}

	return value, nil
}


func (m *MemoryStore) Delete(collection, key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	coll, err := m.data[collection]

	if !err {
		return nil
	}

	delete(coll, key)

	if len(coll) == 0 {
		delete(m.data, collection)
	}

	return nil
}