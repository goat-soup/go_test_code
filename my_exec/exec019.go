package my_exec

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"example.com/go_practise/tool"
)

// Question:
// You are required to write a program to sort the (name, age, height) tuples by ascending order where name is string, age and height are numbers. The tuples are input by console. The sort criteria is:
// 1: Sort based on name;
// 2: Then sort based on age;
// 3: Then sort by score.
// The priority is that name > age > score.
// If the following tuples are given as input to the program:
// Tom,19,80
// John,20,90
// Jony,17,91
// Jony,17,93
// Json,21,85
// Then, the output of the program should be:
// [('John', '20', '90'), ('Jony', '17', '91'), ('Jony', '17', '93'), ('Json', '21', '85'), ('Tom', '19', '80')]

type Tuple struct {
	Name   string
	Age    int
	Height int
}

type Tuples []Tuple

func (t Tuples) Len() int {
	return len(t)
}

func (t Tuples) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Tuples) Less(i, j int) bool {
	if t[i].Name != t[j].Name {
		return t[i].Name < t[j].Name
	}
	if t[i].Age != t[j].Age {
		return t[i].Age < t[j].Age
	}
	return t[i].Height < t[j].Height
}

func Exec019() {
	str := tool.GetStrings()
	var tuples Tuples
	for _, v := range str {
		strs := strings.Split(v, ",")
		age, _ := strconv.Atoi(strs[1])
		height, _ := strconv.Atoi(strs[2])
		t := Tuple{Name: strs[0], Age: age, Height: height}
		tuples = append(tuples, t)
	}
	// sort
	sort.Sort(tuples)
	fmt.Println(tuples)
}
