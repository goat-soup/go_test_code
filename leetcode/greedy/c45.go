package greedy

func Jump(nums []int) int {
	step := 0
	end := 0
	max_dst := 0
	for i := 0; i < len(nums)-1; i++ {
		max_dst = max(max_dst, i+nums[i])
		if i == end {
			end = max_dst
			step++
		}
	}
	return step
}
