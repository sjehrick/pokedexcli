// Package pokecache provides caching functionality for pokeapi clients
package pokecache

import (
	"time"
)

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		interval: interval,
		entry:    make(map[string]cacheEntry),
	}
	go cache.reapLoop()

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	e, ok := c.entry[key]
	if !ok {
		return nil, false
	}
	return e.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for t := range ticker.C {
		c.mu.Lock()
		for k, e := range c.entry {
			if e.createdAt.Before(t.Add(-c.interval)) {
				delete(c.entry, k)
			}
		}
		c.mu.Unlock()
	}
}
