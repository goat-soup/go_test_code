package stack

func DailyTemperatures(temperatures []int) []int {
	s := []int{}
	ans := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for len(s) != 0 && temperatures[i] > temperatures[s[len(s)-1]] {
			ans[s[len(s)-1]] = i - s[len(s)-1]
			s = s[:len(s)-1]
		}
		s = append(s, i)
	}
	return ans
}
