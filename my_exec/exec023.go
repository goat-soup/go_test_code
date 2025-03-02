package my_exec

import (
	"fmt"

	"example.com/go_practise/tool"
)

//Write a method which can calculate square value of number

func Exec023() {
	n := tool.GetInt()
	fmt.Println(n * n)
}
