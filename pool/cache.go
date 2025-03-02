package pool

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data        map[string]string
	m           sync.Mutex
	cache       map[string]time.Time
	ttl         time.Duration
	cleanupTick time.Duration
}

func NewCache(ttl, ct time.Duration) *Cache {
	c := &Cache{
		data:        make(map[string]string),
		cache:       make(map[string]time.Time),
		ttl:         ttl,
		cleanupTick: ct,
	}
	go c.cleanup()
	return c
}

func (c *Cache) Get(key string) (string, bool) {
	c.m.Lock()
	defer c.m.Unlock()
	if tim, ok := c.cache[key]; ok {
		if time.Now().After(tim) {
			delete(c.data, key)
			delete(c.cache, key)
			return "", false
		}
	}
	val, exists := c.data[key]
	return val, exists
}

func (c *Cache) Set(k, v string) {
	c.m.Lock()
	defer c.m.Unlock()
	c.data[k] = v
	c.cache[k] = time.Now().Add(c.ttl)
}

func (c *Cache) cleanup() {
	ticker := time.NewTicker(c.cleanupTick)
	defer ticker.Stop()

	for range ticker.C {
		c.cleanExpire()
	}
}

func (c *Cache) cleanExpire() {
	c.m.Lock()
	defer c.m.Unlock()
	timeNow := time.Now()
	for val, tim := range c.cache {
		if timeNow.After(tim) {
			fmt.Printf("TimeOut: clean expired chache: %s\n", val)
			delete(c.data, val)
			delete(c.cache, val)
		}
	}
}
func CacheTest() {
	c := NewCache(3*time.Second, 4*time.Second)
	go func() {

	}()
	c.Set("gcsun", "cool")
	time.Sleep(1 * time.Second)
	val, ok := c.Get("gcsun")
	if ok {
		fmt.Printf("cache success: %s\n", val)
	}
	time.Sleep(4 * time.Second)
	val, ok = c.Get("gcsun")
	if !ok {
		fmt.Printf("cache failed: %s\n", val)
	}
}
