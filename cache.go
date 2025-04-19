package main

import (
	"fmt"
	"sync"
	"time"
)

var GlobalCache = NewCache(time.Minute * 10)

type Cache struct {
	entries map[string]cacheEntry
	mutex   *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

func (cache *Cache) Add(key string, value []byte) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	fmt.Println("Adding cache entry:", key)

	cache.entries[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		value:     value,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()

	fmt.Println("Retrieving cache entry:", key)

	entry, ok := cache.entries[key]
	if !ok {
		fmt.Println("Cache key miss:", key)
		return nil, false
	}
	fmt.Println("Cache key hit:", key)
	return entry.value, true
}

func (cache *Cache) reap(timeout time.Duration) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	fmt.Println("Reaping cache...")

	for key, entry := range cache.entries {
		if time.Since(entry.createdAt) > timeout {
			fmt.Println("Reaping cache key:", key)
			delete(cache.entries, key)
		}
	}
}

func (cache *Cache) startReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		cache.reap(interval)
	}
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entries: map[string]cacheEntry{},
		mutex:   &sync.RWMutex{},
	}

	go cache.startReapLoop(interval)

	return cache
}
