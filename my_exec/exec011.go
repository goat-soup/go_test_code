package my_exec

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/go_test_code/log"
	"example.com/go_test_code/my_const"
	"example.com/go_test_code/tool"
)

func Exec011() {
	log.Info(my_const.INPUT_STRING)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	str := strings.Split(s, ",")
	for _, v := range str {
		num, _ := tool.TwoTOTen(v)
		if num%5 == 0 {
			fmt.Println(v)
		}
	}
}
