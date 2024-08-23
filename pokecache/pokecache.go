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
	Map map[string]cacheEntry
	mu  *sync.RWMutex
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		Map: map[string]cacheEntry{},
		mu:  &sync.RWMutex{},
	}
	ticker := time.NewTicker(interval)
	cache.reapLoop(ticker.C, interval)
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Lock()
	c.Map[key] = entry
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	entry, ok := c.Map[key]
	c.mu.RUnlock()

	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(channel <-chan time.Time, interval time.Duration) {
	go func() {
		for range channel {
			for key, val := range c.Map {
				if time.Since(val.createdAt) >= interval {
					c.mu.Lock()
					delete(c.Map, key)
					c.mu.Unlock()
				}
			}
		}
	}()
}
