package concurrency

import (
	"context"
	"sync"
	"time"
)

/*
要求实现一个map：
1、面向高并发。
2、查询时，若key存在，则直接返回，否则阻塞等待，知道有kv被写入，等待一定时间超时报错。
3、读写都是O{1}
*/
type MyConcurrencyMap struct {
	sync.Mutex
	mp map[int]int
	kc map[int]chan struct{}
}

func NewMyConcurrencyMap() *MyConcurrencyMap {
	return &MyConcurrencyMap{
		mp: make(map[int]int),
		kc: make(map[int]chan struct{}),
	}
}
func (m *MyConcurrencyMap) Put(k, v int) {
	m.Lock()
	defer m.Unlock()
	m.mp[k] = v

	select {
	// 防止重复关闭
	case <-m.kc[k]:
		return
	default:
		// 关闭管道，会出发读操作。
		close(m.kc[k])
	}
}

func (m *MyConcurrencyMap) Get(k int, maxWaitTime time.Duration) (int, error) {
	m.Lock()
	v, ok := m.mp[k]
	if ok {
		m.Unlock()
		return v, nil
	}
	ch, ok := m.kc[k]
	if !ok {
		ch = make(chan struct{})
		m.kc[k] = ch
	}
	// 阻塞之前开锁
	m.Unlock()
	// 阻塞等待写入
	ctx, concel := context.WithTimeout(context.Background(), maxWaitTime)
	defer concel()
	// 等待唤醒或者超时
	select {
	case <-ch:
	case <-ctx.Done():
		return -1, ctx.Err()
	}
	// 如果被唤醒，说明由kv写入，直接读取就行
	m.Lock()
	v = m.mp[k]
	m.Unlock()
	return v, nil
}
