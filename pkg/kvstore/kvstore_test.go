package kvstore_test

import (
	"testing"

	"github.com/ravidhavlesha/go-multiproto-kvs/pkg/kvstore"
	"github.com/stretchr/testify/assert"
)

func TestKVStore_Set(t *testing.T) {
	store := kvstore.NewKVStore()

	t.Run("Set and Get a valid key-value pair", func(t *testing.T) {
		err := store.Set("foo", "bar")
		assert.NoError(t, err)

		value, exists := store.Get("foo")
		assert.True(t, exists)
		assert.Equal(t, "bar", value)
	})

	t.Run("Set an empty key", func(t *testing.T) {
		err := store.Set("", "bar")
		assert.Error(t, err)
		assert.Equal(t, "key cannot be empty", err.Error())
	})
}

func TestKVStore_Get(t *testing.T) {
	store := kvstore.NewKVStore()

	t.Run("Get a non-existing key", func(t *testing.T) {
		value, exists := store.Get("foo")
		assert.False(t, exists)
		assert.Equal(t, "", value)
	})

}

func TestKVStore_Delete(t *testing.T) {
	store := kvstore.NewKVStore()

	t.Run("Delete an existing key", func(t *testing.T) {
		err := store.Set("foo", "bar")
		assert.NoError(t, err)

		err = store.Delete("foo")
		assert.NoError(t, err)

		_, exists := store.Get("foo")
		assert.False(t, exists)
	})

	t.Run("Delete an non-existing key", func(t *testing.T) {
		err := store.Delete("foo")
		assert.Error(t, err)
		assert.Equal(t, "key does not exists", err.Error())
	})

	t.Run("Delete an empty key", func(t *testing.T) {
		err := store.Delete("")
		assert.Error(t, err)
		assert.Equal(t, "key cannot be empty", err.Error())
	})

}
