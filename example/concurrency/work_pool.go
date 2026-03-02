package concurrency

import (
	"fmt"
	"sync"
)

func WorkPool() {
	//ctx, concel := context.WithCancel(context.Background())
	//errBuff := make(chan string, 1)

	job := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, job, &wg)
	}
	for i := 0; i < 20; i++ {
		job <- i
	}

	close(job)
	wg.Wait()
}

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Println("worker", id, "finished job ", job)
	}
}
