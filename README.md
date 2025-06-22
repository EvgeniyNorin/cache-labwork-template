# Caching Lab Work

## Overview
This lab work focuses on implementing various caching strategies in Go. You will implement the `Cache` interface with different caching algorithms and features.

## Learning Objectives
- Understand different caching strategies (FIFO, LRU, LFU, ARC)
- Implement generic interfaces in Go
- Work with time-based expiration (TTL)
- Write efficient cache implementations

## Cache Interface
```go
type Cache[K comparable, V any] interface {
    Get(key K) (V, error)
    Set(key K, value V) error
    Delete(key K) error
    Clear()
}
```

## Required Implementations

### 1. FIFO Cache (First In, First Out)
- Implement `NewFIFOCache[K comparable, V any](capacity int) Cache[K, V]`
- When cache is full, remove the oldest entry

### 2. LRU Cache (Least Recently Used)
- Implement `NewLRUCache[K comparable, V any](capacity int) Cache[K, V]`
- When cache is full, remove the least recently accessed entry

### 3. LFU Cache (Least Frequently Used)
- Implement `NewLFUCache[K comparable, V any](capacity int) Cache[K, V]`
- When cache is full, remove the entry with the lowest access frequency

### 4. TTL Cache (Time To Live)
- Implement `NewTTLCache[K comparable, V any](capacity int, ttl time.Duration) Cache[K, V]`
- Entries expire after the specified TTL duration
- Expired entries should be removed on access

### 5. ARC Cache (Advanced Replacement Cache) - Advanced Task
- Implement `NewARCCache[K comparable, V any](capacity int) Cache[K, V]`
- Adaptive replacement cache that balances between LRU and LFU

## Error Handling
- `Get` should return an error if the key doesn't exist
- `Set` should return an error if the cache is full and eviction fails
- `Delete` should return an error if the key doesn't exist

## Testing
- Run tests with: `go test ./tests -v`
- Check coverage with: `go test ./tests -cover`
- All tests must pass for full credit

## Submission
- Implement all required cache strategies in `cache/cache.go`
- Ensure all tests pass
- Submit via GitHub Classroom

## Grading Criteria
- Correctness of implementations (60%)
- Test coverage (20%)
- Code quality and efficiency (20%)

## Hints
- Use Go's `container/list` package for linked list operations
- Use `time.Now()` and `time.Duration` for TTL implementation
- Think about the trade-offs between different caching strategies
- Focus on correctness first, then optimize for performance 