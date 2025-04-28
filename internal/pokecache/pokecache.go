package pokecache

import (
	"time"
	"sync"
	//"fmt"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu   sync.Mutex
	data map[string]cacheEntry
}

func NewCache(interval time.Duration) Cache {
	result_cache := Cache { data: make(map[string]cacheEntry),}
	cache_reaploop := func(interval time.Duration) {
		result_cache.reapLoop(interval)
		time.Sleep(interval) //TODO instead of this I need to use time.Ticker with some channel to run the reaploop thing
	}
	go cache_reaploop(interval)
	return result_cache
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cache_entry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return cache_entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	past_time := time.Now().Add(-1 * interval)
	for key, cache_entry := range c.data {
		if cache_entry.createdAt.Before(past_time) {
			delete(c.data, key)
		}
	}
}
