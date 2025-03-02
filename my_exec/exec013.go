package my_exec

import (
	"bufio"
	"fmt"
	"os"

	"example.com/go_practise/log"
	"example.com/go_practise/my_const"
	"example.com/go_practise/tool"
)

func Exec013() {
	log.Info(my_const.INPUT_STRING)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	var letter, digit int
	for _, v := range str {
		if tool.IsLetter(v) {
			letter++
		}
		if tool.IsDigit(v) {
			digit++
		}
	}
	fmt.Printf("LETTER: %d\nDIGIT: %d\n", letter, digit)
}
