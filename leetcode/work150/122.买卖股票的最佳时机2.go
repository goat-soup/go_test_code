package work150

func maxProfit2(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		ans += max(prices[i]-prices[i-1], 0)
	}
	return ans
}
