package concurrency

import "sync"

type shard struct {
	mu sync.RWMutex
	m  map[string]int
}

type ConcurrentMap struct {
	shards []shard
}

// 创建
func NewConcurrentMap(n int) *ConcurrentMap {
	cm := &ConcurrentMap{
		shards: make([]shard, n),
	}
	for i := 0; i < n; i++ {
		cm.shards[i].m = make(map[string]int)
	}
	return cm
}

// hash函数（简单版）
func (cm *ConcurrentMap) getShard(key string) *shard {
	h := fnv32(key)
	return &cm.shards[h%uint32(len(cm.shards))]
}

func fnv32(key string) uint32 {
	var h uint32 = 2166136261
	for i := 0; i < len(key); i++ {
		h ^= uint32(key[i])
		h *= 16777619
	}
	return h
}

// 写
func (cm *ConcurrentMap) Set(key string, value int) {
	s := cm.getShard(key)
	s.mu.Lock()
	s.m[key] = value
	s.mu.Unlock()
}

// 读
func (cm *ConcurrentMap) Get(key string) (int, bool) {
	s := cm.getShard(key)
	s.mu.RLock()
	v, ok := s.m[key]
	s.mu.RUnlock()
	return v, ok
}
