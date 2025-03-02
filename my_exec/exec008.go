package my_exec

import (
	"fmt"
	"sort"
	"strings"
)

func Exec008() {
	fmt.Printf("please enter a string:")
	var s string
	fmt.Scan(&s)
	str := strings.Split(s, ",")
	sort.Strings(str)
	fmt.Println(strings.Join(str, ","))
}
