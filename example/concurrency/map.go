package concurrency

import (
	"sync"
	"time"
)

type MyConcurrency struct {
	sync.Mutex
}

func (m *MyConcurrency) Put(k, v int) {

}

func (m *MyConcurrency) Get(k int, maxWaitTime time.Duration) (int, error) {

}
