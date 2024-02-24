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
	cacheEntryMap map[string]cacheEntry
	mu            *sync.RWMutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	ce := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.cacheEntryMap[key] = ce
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	cEntry, ok := c.cacheEntryMap[key]
	if !ok {
		// slog.Info("Cache MISS for Key : ", key)
		return nil, false
	}
	// slog.Info("Cache HIT for Key :", key)
	return cEntry.val, true

}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cacheEntryMap: make(map[string]cacheEntry),
		mu:            &sync.RWMutex{},
	}
	ticker := time.NewTicker(interval)
	go c.ReapLoop(*ticker, interval)
	return c
}

func (c *Cache) ReapLoop(t time.Ticker, interval time.Duration) {
	for range t.C {
		for k, v := range c.cacheEntryMap {
			if time.Since(v.createdAt) >= interval {
				c.mu.Lock()
				// slog.Info("Deleting Expired key from cache :", k)
				delete(c.cacheEntryMap, k)
				c.mu.Unlock()
			}
		}
	}
}
