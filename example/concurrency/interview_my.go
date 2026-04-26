package concurrency

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// ============================================================
// 题目1: 交替打印奇偶数
// 使用两个goroutine交替打印1-100的奇数和偶数
// 考察点: channel通信、goroutine同步
// ============================================================
func AlternatePrintOddEvenM() {
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	done := make(chan struct{})
	//奇数
	go func() {
		for i := 1; i < 100; i += 2 {
			<-c1
			fmt.Println("A: ", i)
			c2 <- struct{}{}
		}
	}()
	// 偶数
	go func() {
		for i := 2; i <= 100; i += 2 {
			<-c2
			fmt.Println("B: ", i)
			if i == 100 {
				done <- struct{}{}
				return
			}
			c1 <- struct{}{}
		}
	}()
	// 启动进程
	c1 <- struct{}{}
	<-done
}

// ============================================================
// 题目3: 实现一个并发安全的计数器
// 考察点: sync.Mutex、atomic操作、竞态条件
// ============================================================

type SafeCounterAtomicM struct {
	count int64
}

func (c *SafeCounterAtomicM) Incre() {
	// ++
	atomic.AddInt64(&c.count, 1)
}

func (c *SafeCounterAtomicM) Get() int64 {
	return atomic.LoadInt64(&c.count)
}

func Counter() {
	c := &SafeCounterAtomicM{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Incre()
		}()
	}
	wg.Wait()
	fmt.Println(c.Get())
}
