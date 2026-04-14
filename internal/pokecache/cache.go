package pokecache

import (
	"time"
)

func NewCache(interval time.Duration) {
}

func (c *Cache) Add(key string, val []byte) {
}

func (c *Cache) Get(key string) ([]byte, bool) {
}

func (c *Cache) reapLoop() {
}
