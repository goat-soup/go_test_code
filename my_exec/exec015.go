package my_exec

import (
	"fmt"

	"example.com/go_practise/tool"
)

func Exec015() {
	n := tool.GetInt()
	fmt.Println(n + n*11 + n*111 + n*1111)
}
