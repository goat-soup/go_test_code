package my_exec

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Exec006() {
	fmt.Printf("please enter a string, eg. 100,150,180: ")
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		fmt.Println("invalid input:", err)
		return
	}
	var nums []string
	for _, num := range strings.Split(s, ",") {
		n, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := math.Sqrt(2.0 * 50.0 * float64(n) / 30.0)
		nums = append(nums, strconv.Itoa(int(math.Round(result))))
	}
	fmt.Println("result:", strings.Join(nums, ","))
}
