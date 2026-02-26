package leetcode

func IsPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	if n%3 == 0 {
		return IsPowerOfThree(n / 3)
	}
	// math.Max()
	return false
}
