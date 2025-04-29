package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu   sync.Mutex
	Data map[string]cacheEntry
}

func NewCache(interval time.Duration) Cache {
	result_cache := Cache { Data: make(map[string]cacheEntry),}
	cache_reaploop := func(interval time.Duration) {
		result_cache.reapLoop(interval)
	}
	go cache_reaploop(interval)
	return result_cache
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	cache_entry, ok := c.Data[key]
	c.mu.Unlock()
	if !ok {
		return nil, false
	}
	return cache_entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		<- ticker.C
		runCacheCleaner(c, interval)
	}
}

func runCacheCleaner(c *Cache, interval time.Duration) {
	c.mu.Lock()
	for key, cache_entry := range c.Data {
		past_time := time.Now().Add(-1 * interval)
		if cache_entry.createdAt.Before(past_time) {
			delete(c.Data, key)
		}
	}
	c.mu.Unlock()
}
