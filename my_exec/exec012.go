package my_exec

import (
	"fmt"
	"strconv"
	"strings"

	"example.com/go_practise/tool"
)

func Exec012() {
	var result []string
	for i := 100; i <= 300; i++ {
		if tool.IsEven(i) {
			result = append(result, strconv.Itoa(i))
		}
	}
	str := strings.Join(result, ",")
	fmt.Print(str)
}
