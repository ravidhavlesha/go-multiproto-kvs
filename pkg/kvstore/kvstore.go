package kvstore

import "sync"

type KVStore struct {
	store map[string]string
	mu    sync.RWMutex
}

func NewKVStore() *KVStore {
	return &KVStore{store: make(map[string]string)}
}

func (kvs *KVStore) Get(key string) (string, bool) {
	kvs.mu.RLock()
	defer kvs.mu.RUnlock()

	value, exists := kvs.store[key]
	return value, exists
}

func (kvs *KVStore) Set(key, value string) {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()

	kvs.store[key] = value
}

func (kvs *KVStore) Delete(key string) {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()

	delete(kvs.store, key)
}
