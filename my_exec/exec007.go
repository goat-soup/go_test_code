package my_exec

import (
	"fmt"
	"strconv"
	"strings"
)

func Exec007() {
	var s string
	fmt.Print("please enter a string, eg. 3,5 : ")
	_, err := fmt.Scan(&s)
	if err != nil {
		fmt.Print("input error: ", err)
		return
	}
	x, _ := strconv.Atoi(strings.Split(s, ",")[0])
	y, _ := strconv.Atoi(strings.Split(s, ",")[1])
	result := make([][]int, x)
	for i := range result {
		result[i] = make([]int, y)
	}
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			result[i][j] = i * j
		}
	}
	fmt.Printf("%v", result)
}
