package my_exec

import (
	"fmt"

	"example.com/go_test_code/log"
	"example.com/go_test_code/my_const"
)

func factorial(n int) int {
	sum := 1
	for i := 1; i <= n; i++ {
		sum *= i
	}
	return sum
}

func Exec002() {
	var n int
	log.Info(my_const.INPUT_NUMBER)
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Error("invalid number")
	}
	fmt.Println("the factorial of", n, "is", factorial(n))
}
