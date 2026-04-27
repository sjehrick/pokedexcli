package pokecache

import (
	"time"
)

func NewCache(interval time.Duration) {
	cache := Cache{
		Interval: interval,
		Entry:    make(map[string]cacheEntry),
	}
	go cache.reapLoop()
}

func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Entry[key] = cacheEntry{
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	e, ok := c.Entry[key]
	if !ok {
		return nil, false
	}
	return e.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.Interval)
	for t := range ticker.C {
		for _, e := range c.Entry {
		}
	}
}
