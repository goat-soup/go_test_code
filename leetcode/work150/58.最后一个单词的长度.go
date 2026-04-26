package work150

import "strings"

func lengthOfLastWord(s string) int {
	ss := strings.Split(s, " ")
	for i := len(ss) - 1; i >= 0; i-- {
		if len(ss[i]) == 0 {
			continue
		}
		return len(ss[i])
	}
	return 0
}
