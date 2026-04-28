package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	interval time.Duration
	mu       sync.Mutex
	entry    map[string]cacheEntry
}
