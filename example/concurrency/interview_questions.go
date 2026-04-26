package concurrency

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// ============================================================
// Go 并发面试题集锦
// ============================================================

// ============================================================
// 题目1: 交替打印奇偶数
// 使用两个goroutine交替打印1-100的奇数和偶数
// 考察点: channel通信、goroutine同步
// ============================================================
func AlternatePrintOddEven() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	done := make(chan struct{})

	// 打印奇数
	go func() {
		for i := 1; i <= 100; i += 2 {
			<-ch1
			fmt.Println("奇数goroutine:", i)
			ch2 <- struct{}{}
		}
	}()

	// 打印偶数
	go func() {
		for i := 2; i <= 100; i += 2 {
			<-ch2
			fmt.Println("偶数goroutine:", i)
			if i == 100 {
				done <- struct{}{}
				return
			}
			ch1 <- struct{}{}
		}
	}()

	ch1 <- struct{}{} // 启动
	<-done
}

// ============================================================
// 题目2: 使用goroutine和channel实现生产者-消费者模型
// 考察点: channel缓冲、goroutine生命周期管理
// ============================================================
func ProducerConsumer() {
	ch := make(chan int, 10) // 带缓冲的channel
	var wg sync.WaitGroup

	// 生产者
	producer := func(id int) {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			item := id*100 + i
			ch <- item
			fmt.Printf("生产者%d 生产了: %d\n", id, item)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		}
	}

	// 消费者
	consumer := func(id int) {
		defer wg.Done()
		for item := range ch {
			fmt.Printf("消费者%d 消费了: %d\n", id, item)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(150)))
		}
	}

	// 启动3个生产者
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go producer(i)
	}

	// 启动2个消费者
	wg.Add(2)
	go consumer(0)
	go consumer(1)

	// 等待所有生产者完成后关闭channel
	go func() {
		// 需要单独等待生产者完成
		// 这里简化处理，实际应该用单独的WaitGroup
		time.Sleep(2 * time.Second)
		close(ch)
	}()

	wg.Wait()
}

// ============================================================
// 题目3: 实现一个并发安全的计数器
// 考察点: sync.Mutex、atomic操作、竞态条件
// ============================================================

// 方法1: 使用Mutex
type SafeCounterMutex struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounterMutex) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *SafeCounterMutex) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// 方法2: 使用atomic
type SafeCounterAtomic struct {
	count int64
}

func (c *SafeCounterAtomic) Increment() {
	atomic.AddInt64(&c.count, 1)
}

func (c *SafeCounterAtomic) Get() int64 {
	return atomic.LoadInt64(&c.count)
}

func DemoConcurrentCounter() {
	counter := &SafeCounterAtomic{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Println("最终计数:", counter.Get()) // 应该是1000
}

// ============================================================
// 题目4: 使用context控制goroutine的取消
// 考察点: context.WithCancel、context.WithTimeout、优雅退出
// ============================================================
func ContextCancelDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	ch := make(chan string)

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("worker收到取消信号:", ctx.Err())
				return
			case ch <- "工作结果":
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("主goroutine: 超时退出")
			return
		case result := <-ch:
			fmt.Println("收到:", result)
		}
	}
}

// ============================================================
// 题目5: 实现一个简单的工作池(Worker Pool)
// 考察点: goroutine池化、channel分发任务
// ============================================================
type WPJob struct {
	ID      int
	Payload int
}

type WPResult struct {
	Job    WPJob
	Output int
}

func WorkerPoolDemo(numWorkers int, jobs []WPJob) []WPResult {
	jobCh := make(chan WPJob, len(jobs))
	resultCh := make(chan WPResult, len(jobs))

	// 启动worker
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobCh {
				// 模拟处理任务
				output := job.Payload * job.Payload
				fmt.Printf("Worker %d 处理了 Job %d\n", workerID, job.ID)
				resultCh <- WPResult{Job: job, Output: output}
			}
		}(i)
	}

	// 发送任务
	for _, job := range jobs {
		jobCh <- job
	}
	close(jobCh)

	// 等待所有worker完成
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// 收集结果
	var results []WPResult
	for result := range resultCh {
		results = append(results, result)
	}
	return results
}

// ============================================================
// 题目6: 使用sync.Once实现单例模式
// 考察点: sync.Once的使用、并发安全的单例
// ============================================================
type Singleton struct {
	Name string
}

var (
	instance *Singleton
	once     sync.Once
)

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{Name: "唯一实例"}
		fmt.Println("单例被创建")
	})
	return instance
}

// ============================================================
// 题目7: 使用sync.WaitGroup等待多个goroutine完成
// 考察点: WaitGroup的正确使用方式
// ============================================================
func WaitGroupDemo() {
	var wg sync.WaitGroup

	urls := []string{
		"https://example.com",
		"https://google.com",
		"https://github.com",
	}

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			// 模拟HTTP请求
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Printf("完成请求: %s\n", u)
		}(url) // 注意：必须传参，否则闭包会捕获同一个变量
	}

	wg.Wait()
	fmt.Println("所有请求完成")
}

// ============================================================
// 题目8: 使用select实现超时控制
// 考察点: select多路复用、time.After超时
// ============================================================
func SelectTimeoutDemo() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second) // 模拟耗时操作
		ch <- "操作完成"
	}()

	select {
	case result := <-ch:
		fmt.Println("收到结果:", result)
	case <-time.After(1 * time.Second):
		fmt.Println("操作超时!")
	}
}

// ============================================================
// 题目9: 使用sync.Map实现并发安全的map
// 考察点: sync.Map vs 普通map+mutex的区别和适用场景
// ============================================================
func SyncMapDemo() {
	var m sync.Map
	var wg sync.WaitGroup

	// 并发写入
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			m.Store(key, key*key)
		}(i)
	}

	wg.Wait()

	// 遍历
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("key=%v, value=%v\n", key, value)
		return true
	})
}

// ============================================================
// 题目10: 实现一个带超时的锁
// 考察点: channel模拟锁、select超时
// ============================================================
type TimeoutMutex struct {
	ch chan struct{}
}

func NewTimeoutMutex() *TimeoutMutex {
	return &TimeoutMutex{
		ch: make(chan struct{}, 1),
	}
}

func (m *TimeoutMutex) Lock() {
	m.ch <- struct{}{}
}

func (m *TimeoutMutex) Unlock() {
	<-m.ch
}

func (m *TimeoutMutex) TryLock(timeout time.Duration) bool {
	select {
	case m.ch <- struct{}{}:
		return true
	case <-time.After(timeout):
		return false
	}
}

// ============================================================
// 题目11: 使用channel实现信号量(Semaphore)
// 考察点: 利用带缓冲channel控制并发数
// ============================================================
type Semaphore struct {
	ch chan struct{}
}

func NewSemaphore(maxConcurrency int) *Semaphore {
	return &Semaphore{
		ch: make(chan struct{}, maxConcurrency),
	}
}

func (s *Semaphore) Acquire() {
	s.ch <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.ch
}

func SemaphoreDemo() {
	sem := NewSemaphore(3) // 最多3个并发
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sem.Acquire()
			defer sem.Release()

			fmt.Printf("Goroutine %d 开始执行\n", id)
			time.Sleep(time.Second)
			fmt.Printf("Goroutine %d 执行完毕\n", id)
		}(i)
	}

	wg.Wait()
}

// ============================================================
// 题目12: 经典面试陷阱 - goroutine泄漏
// 考察点: goroutine泄漏的原因和解决方案
// ============================================================

// 错误示例：goroutine泄漏
func LeakyGoroutine() {
	ch := make(chan int)

	go func() {
		val := <-ch // 如果没有人发送数据，这个goroutine永远阻塞
		fmt.Println(val)
	}()

	// 函数返回，ch没有人发送数据，goroutine泄漏!
}

// 正确示例：使用context取消避免泄漏
func NoLeakGoroutine() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int)

	go func() {
		select {
		case val := <-ch:
			fmt.Println(val)
		case <-ctx.Done():
			fmt.Println("goroutine被正确取消")
			return
		}
	}()

	// 不需要结果时，取消goroutine
	cancel()
	time.Sleep(100 * time.Millisecond) // 等待goroutine退出
}

// ============================================================
// 题目13: 使用channel实现Fan-in/Fan-out模式
// 考察点: 多个goroutine的结果合并
// ============================================================

// Fan-out: 将一个channel的数据分发给多个goroutine处理
// Fan-in: 将多个channel的结果合并到一个channel
func FanInFanOut() {
	// Fan-out
	input := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			input <- i
		}
		close(input)
	}()

	// 创建多个worker channel (Fan-out)
	numWorkers := 3
	workers := make([]chan int, numWorkers)
	for i := 0; i < numWorkers; i++ {
		workers[i] = make(chan int, 10)
		go func(workerCh chan int, id int) {
			for val := range workerCh {
				fmt.Printf("Worker %d 处理: %d, 结果: %d\n", id, val, val*val)
			}
		}(workers[i], i)
	}

	// 分发任务
	idx := 0
	for val := range input {
		workers[idx%numWorkers] <- val
		idx++
	}

	// 关闭所有worker channel
	for _, ch := range workers {
		close(ch)
	}

	time.Sleep(time.Second) // 等待处理完成
}

// Fan-in: 合并多个channel
func FanIn(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	merged := make(chan int)

	// 为每个channel启动一个goroutine
	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				merged <- val
			}
		}(ch)
	}

	// 所有输入channel完成后关闭合并channel
	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

// ============================================================
// 题目14: 使用sync.Cond实现条件变量
// 考察点: sync.Cond的使用场景和注意事项
// ============================================================
func CondDemo() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	ready := false

	// 等待者
	go func() {
		mu.Lock()
		for !ready {
			cond.Wait() // 释放锁并等待通知，被唤醒后重新获取锁
		}
		fmt.Println("条件满足，继续执行")
		mu.Unlock()
	}()

	// 通知者
	time.Sleep(time.Second)
	mu.Lock()
	ready = true
	cond.Signal() // 唤醒一个等待的goroutine
	mu.Unlock()

	time.Sleep(100 * time.Millisecond)
}

// ============================================================
// 题目15: 使用errgroup处理并发错误
// 考察点: 并发任务的错误处理策略
// 注意: 需要 golang.org/x/sync/errgroup 包
// 这里用手动实现展示原理
// ============================================================
type ErrGroup struct {
	wg   sync.WaitGroup
	err  error
	once sync.Once
}

func (g *ErrGroup) Go(f func() error) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		if err := f(); err != nil {
			g.once.Do(func() {
				g.err = err // 只记录第一个错误
			})
		}
	}()
}

func (g *ErrGroup) Wait() error {
	g.wg.Wait()
	return g.err
}

func ErrGroupDemo() {
	g := &ErrGroup{}

	urls := []string{"url1", "url2", "url3"}
	for _, url := range urls {
		url := url // 避免闭包问题
		g.Go(func() error {
			// 模拟请求
			fmt.Printf("请求: %s\n", url)
			if url == "url2" {
				return fmt.Errorf("请求 %s 失败", url)
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("发生错误:", err)
	}
}

// ============================================================
// 题目16: 经典面试题 - 3个goroutine按顺序打印ABC
// 考察点: goroutine同步、channel编排
// ============================================================
func PrintABC(times int) {
	chA := make(chan struct{}, 1)
	chB := make(chan struct{}, 1)
	chC := make(chan struct{}, 1)
	done := make(chan struct{})

	go func() {
		for i := 0; i < times; i++ {
			<-chA
			fmt.Print("A")
			chB <- struct{}{}
		}
	}()

	go func() {
		for i := 0; i < times; i++ {
			<-chB
			fmt.Print("B")
			chC <- struct{}{}
		}
	}()

	go func() {
		for i := 0; i < times; i++ {
			<-chC
			fmt.Print("C")
			if i == times-1 {
				done <- struct{}{}
				return
			}
			chA <- struct{}{}
		}
	}()

	chA <- struct{}{} // 启动
	<-done
	fmt.Println()
}

// ============================================================
// 题目17: 实现一个读写锁的使用场景
// 考察点: sync.RWMutex读写锁、读多写少场景优化
// ============================================================
type ConcurrentCache struct {
	mu    sync.RWMutex
	items map[string]interface{}
}

func NewConcurrentCache() *ConcurrentCache {
	return &ConcurrentCache{
		items: make(map[string]interface{}),
	}
}

func (c *ConcurrentCache) Get(key string) (interface{}, bool) {
	c.mu.RLock() // 读锁，允许多个读者并发
	defer c.mu.RUnlock()
	val, ok := c.items[key]
	return val, ok
}

func (c *ConcurrentCache) Set(key string, value interface{}) {
	c.mu.Lock() // 写锁，独占访问
	defer c.mu.Unlock()
	c.items[key] = value
}

func (c *ConcurrentCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

// ============================================================
// 题目18: 使用channel实现Pipeline模式
// 考察点: channel链式处理、数据流水线
// ============================================================
func Pipeline() {
	// 阶段1: 生成数据
	generator := func(nums ...int) <-chan int {
		out := make(chan int)
		go func() {
			for _, n := range nums {
				out <- n
			}
			close(out)
		}()
		return out
	}

	// 阶段2: 平方
	square := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				out <- n * n
			}
			close(out)
		}()
		return out
	}

	// 阶段3: 过滤（只保留大于10的）
	filter := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				if n > 10 {
					out <- n
				}
			}
			close(out)
		}()
		return out
	}

	// 组装Pipeline
	ch := generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	ch = square(ch)
	ch = filter(ch)

	for result := range ch {
		fmt.Println("Pipeline结果:", result)
	}
}

// ============================================================
// 题目19: 限速器(Rate Limiter)实现
// 考察点: time.Ticker、channel控制请求速率
// ============================================================
type RateLimiter struct {
	ticker *time.Ticker
	done   chan struct{}
}

func NewRateLimiter(rate int) *RateLimiter {
	rl := &RateLimiter{
		ticker: time.NewTicker(time.Second / time.Duration(rate)),
		done:   make(chan struct{}),
	}
	return rl
}

func (rl *RateLimiter) Allow() {
	<-rl.ticker.C // 等待下一个令牌
}

func (rl *RateLimiter) Stop() {
	rl.ticker.Stop()
}

func RateLimiterDemo() {
	limiter := NewRateLimiter(5) // 每秒5个请求
	defer limiter.Stop()

	for i := 0; i < 10; i++ {
		limiter.Allow()
		fmt.Printf("请求 %d 在 %v\n", i, time.Now().Format("15:04:05.000"))
	}
}

// ============================================================
// 题目20: 经典死锁场景及解决方案
// 考察点: 死锁的识别和避免
// ============================================================

// 死锁场景1: 向无缓冲channel发送数据但没有接收者
func DeadlockExample1() {
	// ch := make(chan int)
	// ch <- 1 // 死锁! 没有接收者
	// fmt.Println(<-ch)

	// 解决方案: 使用带缓冲的channel或在另一个goroutine中发送
	ch := make(chan int, 1) // 缓冲大小为1
	ch <- 1
	fmt.Println(<-ch)
}

// 死锁场景2: 两个goroutine互相等待对方的锁
func DeadlockExample2() {
	// var mu1, mu2 sync.Mutex
	// 错误示例:
	// go func() {
	//     mu1.Lock()
	//     time.Sleep(time.Millisecond)
	//     mu2.Lock() // 等待mu2
	//     mu1.Unlock()
	//     mu2.Unlock()
	// }()
	// go func() {
	//     mu2.Lock()
	//     time.Sleep(time.Millisecond)
	//     mu1.Lock() // 等待mu1 -> 死锁!
	//     mu2.Unlock()
	//     mu1.Unlock()
	// }()

	// 解决方案: 保证锁的获取顺序一致
	var mu1, mu2 sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		mu1.Lock()
		time.Sleep(time.Millisecond)
		mu2.Lock()
		fmt.Println("Goroutine 1 完成")
		mu2.Unlock()
		mu1.Unlock()
	}()
	go func() {
		defer wg.Done()
		mu1.Lock() // 保证顺序: 先mu1后mu2
		time.Sleep(time.Millisecond)
		mu2.Lock()
		fmt.Println("Goroutine 2 完成")
		mu2.Unlock()
		mu1.Unlock()
	}()
	wg.Wait()
}
