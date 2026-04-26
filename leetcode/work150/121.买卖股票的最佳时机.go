package work150

// [7,1,5,3,6,4]
func maxProfit(prices []int) int {
	maxPrice := 0
	cur := prices[0]
	for i := 0; i < len(prices); i++ {
		if cur > prices[i] {
			cur = prices[i]
			continue
		}
		if cur < prices[i] {
			maxPrice = max(maxPrice, prices[i]-cur)
		}
	}
	return maxPrice
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
