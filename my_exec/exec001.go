package my_exec

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	strconv.Itoa
	strings.Join

*
*/
func Exec001() {
	var str []string
	for i := 2000; i < 3201; i++ {
		if i%7 == 0 && i%5 != 0 {
			// fmt.Println(i)
			str = append(str, strconv.Itoa(i))
		}
	}
	fmt.Println(strings.Join(str, ","))
}
