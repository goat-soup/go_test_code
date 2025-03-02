package example

import (
	"fmt"
	"time"
)

func ChanTest() {
	c := make(chan bool, 1)
	go func(c chan bool) {
		fmt.Println("start...")
		time.Sleep(time.Second)
		fmt.Println("done...")
		c <- true
	}(c)

	<-c
}

// close chan
func ChanTest1() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			i, ok := <-jobs
			if ok {
				fmt.Println("received jobs: ", i)
			} else {
				fmt.Println("jobs done")
				done <- true
				return

			}
		}
	}()
	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("the jobs :", i)
	}
	close(jobs)
	<-done
	_, ok := <-jobs
	if !ok {
		fmt.Println("jobs done!!!")
	}
}
