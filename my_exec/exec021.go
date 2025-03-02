package my_exec

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"example.com/go_practise/tool"
)

// Question
// A robot moves in a plane starting from the original point (0,0). The robot can move toward UP, DOWN, LEFT and RIGHT with a given steps. The trace of robot movement is shown as the following:
// UP 5
// DOWN 3
// LEFT 3
// RIGHT 2
// ¡­
// The numbers after the direction are steps. Please write a program to compute the distance from current position after a sequence of movement and original point. If the distance is a float, then just print the nearest integer.
// Example:
// If the following tuples are given as input to the program:
// UP 5
// DOWN 3
// LEFT 3
// RIGHT 2
// Then, the output of the program should be:
// 2

func Exec021() {
	strs := tool.GetStrings()
	var dist [4]int
	for k, str := range strs {
		num, _ := strconv.Atoi(strings.Split(str, " ")[1])
		dist[k] = num
	}
	x, y := tool.AbsInt(dist[0]-dist[1]), tool.AbsInt(dist[2]-dist[3])
	fmt.Println(math.Round(math.Sqrt(float64(x*x + y*y))))
}
