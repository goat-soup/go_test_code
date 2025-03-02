package my_exec

import (
	"fmt"
	"sort"
	"strings"

	"example.com/go_test_code/tool"
)

// Question:
// Write a program to compute the frequency of the words from the input. The output should output after sorting the key alphanumerically.
// Suppose the following input is supplied to the program:
// New to Python or choosing between Python 2 and Python 3? Read Python 2 or Python 3.
// Then, the output should be:
// 2:2
// 3.:1
// 3?:1
// New:1
// Python:5
// Read:1
// and:1
// between:1
// choosing:1
// or:2
// to:1

func Exec022() {
	str := tool.GetString()
	words := strings.Split(str, " ")
	results := make(map[string]int, len(words))
	for _, word := range words {
		results[word]++
	}
	var keys []string
	for k := range results {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, v := range keys {
		fmt.Println(v, ":", results[v])
	}
}
