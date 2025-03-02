package my_exec

import (
	"fmt"
	"strings"

	"example.com/go_practise/log"
	"example.com/go_practise/my_const"
)

func Exec004() {
	log.Info(my_const.INPUT_STRING)
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		fmt.Print("please enter a string :", err)
		return
	}
	str := strings.Split(s, ",")
	fmt.Print(str)
}
