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
	Interval time.Duration
	Mu       sync.Mutex
	Entry    map[string]cacheEntry
}
