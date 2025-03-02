package my_exec

import (
	"fmt"

	"example.com/go_practise/log"
	"example.com/go_practise/my_const"
)

func Exec003() {
	log.Info(my_const.INPUT_STRING)
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Print("error invalid number:", err)
		return
	}
	m := make(map[int]int, n)
	for i := 1; i <= n; i++ {
		m[i] = i * i
	}
	fmt.Printf("%v", m)
}
