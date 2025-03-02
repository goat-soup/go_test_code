package my_exec

import (
	"fmt"

	"example.com/go_test_code/tool"
)

func Exec014() {
	s := tool.GetString()
	var up, low int
	for _, v := range s {
		if tool.IsUpLetter(v) {
			up++
		} else if tool.IsLowLetter(v) {
			low++
		}
	}
	fmt.Printf("up: %d\nlow: %d\n", up, low)
}
