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
	data map[string]cacheEntry
}

func NewCache(base_time time.Duration) Cache {
	var result_cache Cache
	//TODO actual solution
	return result_cache
}

func (c *Cache) Add(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntity{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entity, exists := c.data[key]
	if !exists {
		return nil, false
	}
	return entity.val, true
}
