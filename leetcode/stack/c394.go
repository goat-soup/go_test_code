package stack

import (
	"strings"
	"unicode"
)

/*
输入：s = "3[a]2[bc]"
输出："aaabcbc"
*/
func DecodeString(s string) string {
	numStack := []int{}
	strStack := []string{}
	k := 0
	currStr := ""
	for _, c := range s {
		if unicode.IsDigit(c) {
			k = k*10 + int(c-'0')
		} else if c == '[' {
			numStack = append(numStack, k)
			strStack = append(strStack, currStr)
			currStr = "" //重新开始
			k = 0        // 重新计数
		} else if c == ']' {
			repeat := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			preStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			currStr = preStr + strings.Repeat(currStr, repeat)
		} else {
			currStr += string(c)
		}
	}
	return currStr
}
