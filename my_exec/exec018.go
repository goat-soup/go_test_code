package my_exec

import (
	"fmt"
	"strings"

	"example.com/go_test_code/tool"
)

// Question:
// A website requires the users to input username and password to register. Write a program to check the validity of password input by users.
// Following are the criteria for checking the password:
// 1. At least 1 letter between [a-z]
// 2. At least 1 number between [0-9]
// 1. At least 1 letter between [A-Z]
// 3. At least 1 character from [$#@]
// 4. Minimum length of transaction password: 6
// 5. Maximum length of transaction password: 12
// Your program should accept a sequence of comma separated passwords and will check them according to the above criteria. Passwords that match the criteria are to be printed, each separated by a comma.
// Example
// If the following passwords are given as input to the program:
// ABd1234@1,a F1#,2w3E*,2We3345
// Then, the output of the program should be:
// ABd1234@1

func Exec018() {
	fmt.Println("func: ", tool.GetFuncName())
	s := tool.GetString()
	var result []string
	str := strings.Split(s, ",")
	for _, v := range str {
		if len(v) < 6 || len(v) > 12 {
			// fmt.Println("please input a password length [6, 12]")
			continue
		}
		if !strings.ContainsAny(v, "$#@") {
			// fmt.Println("please input a password contain $#@")
			continue
		}
		var isdigit, islowletter, isupletter bool
		for _, k := range v {
			if tool.IsDigit(k) {
				isdigit = true
			}
			if tool.IsLowLetter(k) {
				islowletter = true
			}
			if tool.IsUpLetter(k) {
				isupletter = true
			}
		}
		if !isdigit || !islowletter || !isupletter {
			// fmt.Println("please input a password contain at least one letter between [a-z] and [A-Z]")
			continue
		}
		result = append(result, v)
	}
	fmt.Println(result)
}
