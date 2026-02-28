package concurrency

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func ABC1() {
	var wg sync.WaitGroup
	wg.Add(3)
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)
	done := make(chan struct{})
	worker := func(name string, in <-chan int, out chan<- int) {
		defer wg.Done()

		for {
			select {
			case <-done:
				return
			case token := <-in:
				{
					if token > 100 {
						return
					}
					fmt.Printf("%s%d ", name, token)
					next := token + 1
					if token == 100 {
						close(done)
						return
					}
					select {
					case <-done:
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
	fmt.Println()

}
func ABC() {
	var wg sync.WaitGroup
	wg.Add(3)
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)
	ctx, concel := context.WithCancel(context.Background())
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
	fmt.Println()

}

func ABC2() {
	var wg sync.WaitGroup
	buffer := make(map[int]bool)
	out := make(chan int)
	done := make(chan struct{})
	ctx, concel := context.WithCancel(context.Background())
	worker := func(id int) {
		//
		defer wg.Done()
		time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
		doSometing := func() error {
			return nil
		}
		if doSometing() != nil {
			concel()
		}
		//
		// fmt.Printf("task %d done\n", id)
		select {
		case out <- id:
		case <-ctx.Done():
			return
		}
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i)
	}
	go func() {
		next := 0
		for i := range out {
			select {
			case <-ctx.Done():
				break
			default:
				buffer[i] = true
				for {
					if _, ok := buffer[next]; ok {
						fmt.Printf("task %d done\n", next)
						delete(buffer, next)
						next++
					} else {
						break
					}
				}
			}

		}
		done <- struct{}{}
	}()
	wg.Wait()
	close(out)
	<-done
}
