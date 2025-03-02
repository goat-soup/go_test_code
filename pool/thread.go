package pool

import (
	"fmt"
	"sync"
	"time"
)

type ThreadPool struct {
	TaskQueue chan func()
}

func NewThreadPool(count int) *ThreadPool {
	pool := &ThreadPool{TaskQueue: make(chan func(), 100)}
	for i := 0; i < count; i++ {
		go pool.Worker(i)
	}
	return pool
}

func (pool *ThreadPool) Worker(id int) {
	for task := range pool.TaskQueue {
		fmt.Printf("Worker %d excuting task\n", id)
		task()
	}
}

func (pool *ThreadPool) Submit(task func()) {
	pool.TaskQueue <- task
}

func ThreadPoolTest() {
	pool := NewThreadPool(3)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		id := i
		wg.Add(1)
		pool.Submit(func() {
			defer wg.Done()
			fmt.Printf("i am task %d\n", id)
			time.Sleep(1 * time.Second)
		})
	}
	wg.Wait()
	// time.Sleep(10 * time.Second)
}
