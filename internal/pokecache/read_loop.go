package pokecache

import (
	"time"
)

func (c *Cache) readLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			var keysToDelete []string

			c.mu.Lock()
			for key, val := range c.cacheMap {
				if val.createdAt.Compare(time.Now()) <= 0 {
					keysToDelete = append(keysToDelete, key)
				}
			}

			for _, key := range keysToDelete {
				delete(c.cacheMap, key)
			}
			c.mu.Unlock()
		}
	}
}
