package leetcode

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	m1 := make(map[string]int)
	ss1 := strings.Split(s1, " ")
	for _, v := range ss1 {
		m1[v]++
	}
	m2 := make(map[string]int)
	ss2 := strings.Split(s2, " ")
	for _, v := range ss2 {
		m2[v]++
	}
	res := make([]string, 0)
	for k, v := range m1 {
		if v > 1 {
			continue
		} else if v == 1 {
			if _, ok := m2[k]; !ok {
				res = append(res, k)
			}
		}
	}
	for k, v := range m2 {
		if v > 1 {
			continue
		} else if v == 1 {
			if _, ok := m1[k]; !ok {
				res = append(res, k)
			}
		}
	}
	return res
}
