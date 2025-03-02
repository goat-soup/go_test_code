package my_exec

import (
	"fmt"

	"example.com/go_practise/log"
	"example.com/go_practise/tool"
)

// Question 20
// Level 3

// Question:
// Define a class with a generator which can iterate the numbers, which are divisible by 7, between a given range 0 and n.

func Exec020() {
	log.Info(tool.GetFuncName())
	n := tool.GetInt()
	for i := 0; i < n; i++ {
		if i%7 == 0 {
			fmt.Println(i)
		}
	}
}
