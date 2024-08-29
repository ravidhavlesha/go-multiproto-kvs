package kvstore

import (
	"errors"
	"log"
	"sync"
)

type KVStoreInterface interface {
	Get(key string) (string, bool)
	Set(key, value string) error
	Delete(key string) error
}

// KVStore is a simple in-memory key-value store with thread-safe access.
type KVStore struct {
	store sync.Map
}

// NewKVStore initializes a new KVStore
func NewKVStore() *KVStore {
	return &KVStore{}
}

// Get retrieves a value for a given key from the store. Returns an empty string if the key does not exists.
func (kvs *KVStore) Get(key string) (string, bool) {
	value, exists := kvs.store.Load(key)
	if exists {
		return value.(string), true
	}

	return "", false
}

// Set stores a key-value pair inthe store.
func (kvs *KVStore) Set(key, value string) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}

	kvs.store.Store(key, value)
	log.Printf("Set key: %s with value: %s \n", key, value)

	return nil
}

// Delete removes a key-value pair from the store.
func (kvs *KVStore) Delete(key string) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}

	_, exists := kvs.Get(key)
	if !exists {
		return errors.New("key does not exists")
	}

	kvs.store.Delete(key)
	log.Printf("Delete key: %s\n", key)

	return nil
}
