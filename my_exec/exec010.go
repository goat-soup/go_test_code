package my_exec

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"example.com/go_practise/set"
)

func Exec010() {
	fmt.Print("please enter a string,Ctrl + D end:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	myset := set.NewSet[string]()
	str := strings.Split(s, " ")
	for _, v := range str {
		myset.Add(v)
	}
	result := myset.ToSlice()
	sort.Strings(result)
	fmt.Println(result)
}
