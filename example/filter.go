package example

import "fmt"

type testInt func(int) bool

func isOdd(a int) bool {
	if a == 0 {
		return false
	}
	if a%2 == 0 {
		return false
	}
	return true
}

func isEven(a int) bool {
	if a == 0 {
		return true
	}
	if a%2 == 0 {
		return true
	}
	return false
}

func filter(s []int, f testInt) []int {
	res := []int{}
	for _, v := range s {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}
func MyFilter() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	odd := filter(s, isOdd)
	fmt.Println("odd :", odd)
	even := filter(s, isEven)
	fmt.Println("even :", even)
}
