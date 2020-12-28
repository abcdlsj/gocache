package gocache

import (
	"gocache/lru"
	"sync"
)

// 实现单机缓存，加锁支持并发
type cache struct {
	mu         sync.Mutex
	lru        *lru.Cache
	cacheBytes int
}

func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// 延迟初始化，对象初始化会延迟到第一次使用该对象
	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes)
		c.lru.OnEvicted = nil
	}
	c.lru.Add(key, value)
}

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}
	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), ok
	}

	return
}
