package example

import (
	"fmt"
	"time"
)

func TimerTest() {
	c := make(chan bool)

	go func() {
		time.Sleep(time.Second * 3)
		c <- true
	}()

	select {
	case msg := <-c:
		fmt.Println(msg)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}
