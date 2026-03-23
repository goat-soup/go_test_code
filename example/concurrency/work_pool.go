package concurrency

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
任务输入 → jobs chan → worker pool

	      ↓
	  并发执行
	      ↓
	  results chan（无序）
	      ↓
	reorder（排序输出）
	      ↓
	    输出
*/
type Task struct {
	ID int
}

type Result struct {
	ID int
}

func worker(ctx context.Context, id int, jobs <-chan Task, results chan<- Result, rate <-chan time.Time, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case task, ok := <-jobs:
			if !ok {
				return
			}

			// 限流
			select {
			case <-rate:
			case <-ctx.Done():
				return
			}

			// 模拟处理
			time.Sleep(100 * time.Millisecond)

			select {
			case results <- Result{ID: task.ID}:
			case <-ctx.Done():
				return
			}
		}
	}
}

func Work_Pool() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := make(chan Task)
	results := make(chan Result)

	// 限流：每 50ms 处理一个
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	var wg sync.WaitGroup

	// 启动 worker pool
	workerNum := 3
	wg.Add(workerNum)
	for i := 0; i < workerNum; i++ {
		go worker(ctx, i, jobs, results, ticker.C, &wg)
	}

	// 生产任务
	go func() {
		for i := 0; i < 10; i++ {
			select {
			case jobs <- Task{ID: i}:
			case <-ctx.Done():
				return
			}
		}
		close(jobs)
	}()

	// 收集结果（有序输出）
	done := make(chan struct{})
	go func() {
		defer close(done)

		buffer := make(map[int]Result)
		next := 0

		for r := range results {
			buffer[r.ID] = r

			for {
				if val, ok := buffer[next]; ok {
					fmt.Printf("result: %d\n", val.ID)
					delete(buffer, next)
					next++
				} else {
					break
				}
			}
		}
	}()

	// 等待 worker 结束
	wg.Wait()
	close(results)

	<-done
}
