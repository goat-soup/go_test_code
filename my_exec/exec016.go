package my_exec

import (
	"fmt"
	"strconv"
	"strings"

	"example.com/go_test_code/tool"
)

// Question:
// Use a list comprehension to square each odd number in a list. The list is input by a sequence of comma-separated numbers.
// Suppose the following input is supplied to the program:
// 1,2,3,4,5,6,7,8,9
// Then, the output should be:
// 1,3,5,7,9
func Exec016() {
	s := tool.GetString()
	str := strings.Split(s, ",")
	var result []string
	for _, v := range str {
		num, _ := strconv.Atoi(v)
		if !tool.IsEven(num) {
			result = append(result, v)
		}
	}
	strings.Join(result, ",")
	fmt.Println(result)
}
