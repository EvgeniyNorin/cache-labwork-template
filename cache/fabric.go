package cache

import "time"

// NewFIFOCache creates a new FIFO (First In, First Out) cache
// TODO: Implement this function
func NewFIFOCache[K comparable, V any](capacity int) Cache[K, V] {
	// Students should implement this
	return &emptyCache[K, V]{}
}

// NewLRUCache creates a new LRU (Least Recently Used) cache
// TODO: Implement this function
func NewLRUCache[K comparable, V any](capacity int) Cache[K, V] {
	// Students should implement this
	return &emptyCache[K, V]{}
}

// NewLFUCache creates a new LFU (Least Frequently Used) cache
// TODO: Implement this function
func NewLFUCache[K comparable, V any](capacity int) Cache[K, V] {
	// Students should implement this
	return &emptyCache[K, V]{}
}

// NewTTLCache creates a new TTL (Time To Live) cache
// TODO: Implement this function
func NewTTLCache[K comparable, V any](capacity int, ttl time.Duration) Cache[K, V] {
	// Students should implement this
	return &emptyCache[K, V]{}
}

// NewARCCache creates a new ARC (Adaptive Replacement Cache)
// TODO: Implement this function (Advanced task)
func NewARCCache[K comparable, V any](capacity int) Cache[K, V] {
	// Students should implement this
	return &emptyCache[K, V]{}
} 