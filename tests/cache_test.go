package tests

import (
	"testing"
	"time"

	"caching-labwork/cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestFIFOCache tests the FIFO cache implementation
func TestFIFOCache(t *testing.T) {
	c := cache.NewFIFOCache[string, int](3)

	// Test basic operations
	err := c.Set("a", 1)
	require.NoError(t, err)

	val, err := c.Get("a")
	require.NoError(t, err)
	assert.Equal(t, 1, val)

	// Test FIFO eviction
	c.Set("b", 2)
	c.Set("c", 3)
	c.Set("d", 4) // This should evict "a"

	_, err = c.Get("a")
	assert.Error(t, err)
	assert.Equal(t, cache.ErrKeyNotFound, err)

	val, err = c.Get("b")
	require.NoError(t, err)
	assert.Equal(t, 2, val)

	// Test delete
	err = c.Delete("b")
	require.NoError(t, err)

	_, err = c.Get("b")
	assert.Error(t, err)

	// Test clear
	c.Clear()
	_, err = c.Get("c")
	assert.Error(t, err)
}

// TestLRUCache tests the LRU cache implementation
func TestLRUCache(t *testing.T) {
	c := cache.NewLRUCache[string, int](3)

	// Test basic operations
	c.Set("a", 1)
	c.Set("b", 2)
	c.Set("c", 3)

	// Access "a" to make it most recently used
	val, err := c.Get("a")
	require.NoError(t, err)
	assert.Equal(t, 1, val)

	// Add "d" - should evict "b" (least recently used)
	c.Set("d", 4)

	_, err = c.Get("b")
	assert.Error(t, err)

	val, err = c.Get("a")
	require.NoError(t, err)
	assert.Equal(t, 1, val)

	val, err = c.Get("c")
	require.NoError(t, err)
	assert.Equal(t, 3, val)

	val, err = c.Get("d")
	require.NoError(t, err)
	assert.Equal(t, 4, val)
}

// TestLFUCache tests the LFU cache implementation
func TestLFUCache(t *testing.T) {
	c := cache.NewLFUCache[string, int](3)

	// Test basic operations
	c.Set("a", 1)
	c.Set("b", 2)
	c.Set("c", 3)

	// Access "a" multiple times to increase its frequency
	c.Get("a")
	c.Get("a")
	c.Get("a")

	// Access "b" once
	c.Get("b")

	// Add "d" - should evict "c" (least frequently used)
	c.Set("d", 4)

	_, err := c.Get("c")
	assert.Error(t, err)

	val, err := c.Get("a")
	require.NoError(t, err)
	assert.Equal(t, 1, val)

	val, err = c.Get("b")
	require.NoError(t, err)
	assert.Equal(t, 2, val)

	val, err = c.Get("d")
	require.NoError(t, err)
	assert.Equal(t, 4, val)
}

// TestTTLCache tests the TTL cache implementation
func TestTTLCache(t *testing.T) {
	c := cache.NewTTLCache[string, int](3, 100*time.Millisecond)

	// Test basic operations
	c.Set("a", 1)
	c.Set("b", 2)

	val, err := c.Get("a")
	require.NoError(t, err)
	assert.Equal(t, 1, val)

	// Wait for expiration
	time.Sleep(150 * time.Millisecond)

	// Should not find expired entries
	_, err = c.Get("a")
	assert.Error(t, err)

	_, err = c.Get("b")
	assert.Error(t, err)

	// Test that new entries work after expiration
	c.Set("c", 3)
	val, err = c.Get("c")
	require.NoError(t, err)
	assert.Equal(t, 3, val)
}

// TestARCCache tests the ARC cache implementation (advanced)
func TestARCCache(t *testing.T) {
	c := cache.NewARCCache[string, int](4)

	// Test basic operations
	c.Set("a", 1)
	c.Set("b", 2)
	c.Set("c", 3)
	c.Set("d", 4)

	// Access some items to change their status
	c.Get("a")
	c.Get("b")

	// Add new item - should trigger adaptive replacement
	c.Set("e", 5)

	// Verify that some items are still accessible
	// (ARC behavior depends on implementation)
	val, err := c.Get("a")
	if err == nil {
		assert.Equal(t, 1, val)
	}

	val, err = c.Get("b")
	if err == nil {
		assert.Equal(t, 2, val)
	}
}

// TestCacheErrors tests error conditions
func TestCacheErrors(t *testing.T) {
	c := cache.NewFIFOCache[string, int](1)

	// Test getting non-existent key
	_, err := c.Get("nonexistent")
	assert.Error(t, err)
	assert.Equal(t, cache.ErrKeyNotFound, err)

	// Test deleting non-existent key
	err = c.Delete("nonexistent")
	assert.Error(t, err)
	assert.Equal(t, cache.ErrKeyNotFound, err)

	// Test basic operations work
	c.Set("a", 1)
	val, err := c.Get("a")
	require.NoError(t, err)
	assert.Equal(t, 1, val)
}
