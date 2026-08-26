package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu       sync.Mutex
	Entries  map[string]cacheEntry
	Interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	t := time.NewTicker(interval)
	cache := Cache{
		Interval: interval,
		Entries:  make(map[string]cacheEntry),
	}
	cache.reapLoop(t)
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Entries[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.Entries[key]
	return v.val, ok
}

func (c *Cache) reapLoop(t *time.Ticker) {
	go func() {
		for {
			now := <-t.C
			c.mu.Lock()
			for k, v := range c.Entries {
				if now.After(v.createdAt.Add(c.Interval)) {
					delete(c.Entries, k)
				}
			}

			c.mu.Unlock()
		}
	}()
}
