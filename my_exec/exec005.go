package my_exec

import (
	"bufio"
	"fmt"
	"os"

	"example.com/go_practise/log"
	"example.com/go_practise/my_const"
)

type Stringer interface {
	ReadString()
	PrintString()
}
type myString struct {
	s string
}

func (s *myString) ReadString() {
	scanner := bufio.NewScanner(os.Stdin)
	log.Info(my_const.INPUT_STRING)
	if scanner.Scan() {
		s.s = scanner.Text()
	} else {
		fmt.Println("Error reading input:", scanner.Err())
	}
}
func (s *myString) PrintString() {
	fmt.Printf("myPrint : = %s\n", s.s)
}

func Exec005() {
	var s Stringer = new(myString)
	s.ReadString()
	s.PrintString()
}
