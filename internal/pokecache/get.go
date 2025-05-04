package pokecache

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.cacheMap[key]
	if !ok {
		return nil, false
	}
	return val.val, true
}
