package leetcode

import "math"

const MOD = 1_000_000_007

/*
请你帮忙给从 1 到 n 的数设计排列方案，使得所有的「质数」都应该被放在「质数索引」（索引从 1 开始）上；
你需要返回可能的方案总数。
*/
func IsPrime(n int) bool {
	if n == 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func factorial(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res = (res * i) % MOD
	}
	return res
}
func NumPrimeArrangements(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if IsPrime(i) {
			count++
		}
	}
	return (factorial(count) * factorial(n-count)) % MOD
}
