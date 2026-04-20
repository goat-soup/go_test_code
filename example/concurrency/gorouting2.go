package concurrency

import (
	"context"
	"fmt"
	"sync"
)

func ABC3() string {
	var wg sync.WaitGroup
	wg.Add(3)
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)
	ctx, concel := context.WithCancel(context.Background())
	// 启动一个收集器
	var reswg sync.WaitGroup
	reswg.Add(1)
	res := ""
	resCh := make(chan string, 100)
	go func() {
		defer reswg.Done()
		for r := range resCh {
			res += r
		}
	}()
	worker := func(name string, in <-chan int, out chan<- int) {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case token := <-in:
				{
					if token > 100 {
						return
					}
					fmt.Printf("%s%d ", name, token)
					resCh <- name
					next := token + 1
					if token == 100 {
						concel()
						return
					}
					select {
					case <-ctx.Done():
						return
					case out <- next:
					}
				}
			}
		}
	}

	go worker("A", a, b)
	go worker("B", b, c)
	go worker("C", c, a)

	a <- 1

	wg.Wait()
	close(resCh)
	reswg.Wait()

	fmt.Println()
	return res
}
