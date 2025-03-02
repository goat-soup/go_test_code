package my_exec

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Exec009() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("input some lines")
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		fmt.Println(strings.ToUpper(line))
	}
}
