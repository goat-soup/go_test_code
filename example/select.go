package example

import (
	"fmt"
	"time"

	"example.com/go_test_code/tool"
)

func SelectTest() {
	tool.GetFuncName()
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()
	select {
	case r := <-c1:
		fmt.Println(r)
	case r := <-c2:
		fmt.Println(r)
	}
}
