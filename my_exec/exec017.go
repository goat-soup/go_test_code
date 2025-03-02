package my_exec

import (
	"fmt"
	"strconv"
	"strings"

	"example.com/go_practise/log"
	"example.com/go_practise/tool"
)

// Question:
// Write a program that computes the net amount of a bank account based a transaction log from console input. The transaction log format is shown as following:
// D 100
// W 200

// D means deposit while W means withdrawal.
// Suppose the following input is supplied to the program:
// D 300
// D 300
// W 200
// D 100
// Then, the output should be:
// 500
func Exec017() {
	log.Info("===Exec017===")
	str := tool.GetStrings()
	var result int
	for _, v := range str {
		num, _ := strconv.Atoi(strings.Split(v, " ")[1])
		if tool.IsEqual(strings.Split(v, " ")[0], "W") {
			result -= num
		}
		if tool.IsEqual(strings.Split(v, " ")[0], "D") {
			result += num
		}
	}
	fmt.Println(result)
}
