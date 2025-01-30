package pokecache

import (
	"sync"
	"time"
)

// Cache
type Cache struct {
	entries map[string]cacheEntry // The map of cache
	mutex   *sync.Mutex           // The mutex to protect the map
	td      time.Duration         // The life time duration of this cache values
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte // Represents the raw data we're caching
}

// NewCache creates a new cache with a configurable interval (time.Duration)
func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entries: make(map[string]cacheEntry), // Initializes an empty map
		mutex:   &sync.Mutex{},               // Initializes a mutex
		td:      interval,                    // The life time duration of this cache values
	}

	go cache.ReapLoop(interval) // Calls the CleanUp method (we called it ReapLoop), a goroutine to periodically clean the cache's entries that are older than the cache's life time duration.
	return cache
}

// Add cache
func (c *Cache) Add(key string, value []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

// Get cache data
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	entry, exists := c.entries[key]
	if !exists {
		return nil, false
	}

	return entry.val, true
}

// create a method to periodically clean the cache's entries that are older than the cache's life time duration
func (c *Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval) // Create a ticker, that runs every "interval" time, to periodically clean the cache's entries
	go func() {
		for range ticker.C {
			c.mutex.Lock()
			for key, entry := range c.entries {
				if time.Since(entry.createdAt) > c.td {
					delete(c.entries, key)
				}
			}
			c.mutex.Unlock()
		}
	}()
}
