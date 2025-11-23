package inmemcache

import (
	"backend/internal/application/ports"
	"time"

	"github.com/patrickmn/go-cache"
)

type InMemoryCache struct {
	cache.Cache
}

func NewInMemoryCache(defaultTTL, defaultCleanupInterval time.Duration) ports.CacheStore {
	return &InMemoryCache{
		Cache: *cache.New(defaultTTL, defaultCleanupInterval),
	}
}

func (c *InMemoryCache) Get(k string) (any, bool) {
	return c.Cache.Get(k)
}

func (c *InMemoryCache) Set(key string, value any) error {
	c.SetDefault(key, value)
	return nil
}

func (c *InMemoryCache) SetWithTTL(key string, value any, duration time.Duration) error {
	c.Cache.Set(key, value, duration)
	return nil
}

func (c *InMemoryCache) Delete(key string) error {
	c.Cache.Delete(key)
	return nil
}

func (c *InMemoryCache) Flush() error {
	c.Cache.Flush()
	return nil
}
