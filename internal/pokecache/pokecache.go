package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	items    map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		items:    make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	c.items[key] = cacheEntry{createdAt: time.Now(), val: value}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	if _, ok := c.items[key]; ok {
		val := c.items[key].val
		c.mu.Unlock()
		return val, true
	} else {
		c.mu.Unlock()
		return nil, false
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.items {
			if time.Since(v.createdAt) > interval {
				delete(c.items, k)
			}
		}
		c.mu.Unlock()
	}
}
