package tool

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"

	"example.com/go_practise/log"
	"example.com/go_practise/my_const"
)

/**
* 2to10
* author: gcsun
* date: 2024/09/21
* version: 1.0.0
 */
func TwoTOTen(binary string) (int64, error) {
	return strconv.ParseInt(binary, 2, 64)
}

func IsEven(num int) bool {
	return num%2 == 0
}

func IsLetter(s rune) bool {
	return IsUpLetter(s) || IsLowLetter(s)
}

func IsUpLetter(s rune) bool {
	return s >= 'A' && s <= 'Z'
}
func IsLowLetter(s rune) bool {
	return s >= 'a' && s <= 'z'
}
func IsDigit(s rune) bool {
	if s >= '0' && s <= '9' {
		return true
	}
	return false
}

func GetString() string {
	log.Info(my_const.INPUT_STRING)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func GetInt() int {
	log.Info(my_const.INPUT_NUMBER)
	var num int
	fmt.Scan(&num)
	return num
}

func GetStrings() []string {
	log.Info(my_const.INPUT_STRINGS)
	var result []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}

func IsEqual(a, b string) bool {
	return a == b
}

func GetFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

// AbsInt return the number of int
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
