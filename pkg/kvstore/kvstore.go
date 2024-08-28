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

type KVStore struct {
	store sync.Map
}

func NewKVStore() *KVStore {
	return &KVStore{}
}

func (kvs *KVStore) Get(key string) (string, bool) {
	value, exists := kvs.store.Load(key)
	if exists {
		return value.(string), true
	}

	return "", false
}

func (kvs *KVStore) Set(key, value string) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}

	kvs.store.Store(key, value)
	log.Printf("Set key: %s with value: %s \n", key, value)

	return nil
}

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
