package pokecache

import (
	"time"
)

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.cacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Unlock()
}
