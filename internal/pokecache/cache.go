package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu       sync.Mutex
	cacheMap map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cacheMap: make(map[string]cacheEntry),
	}
	go c.readLoop(interval)
	return c
}
