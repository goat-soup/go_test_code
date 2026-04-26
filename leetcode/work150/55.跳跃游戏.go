package work150

func canJump(nums []int) bool {
	step := 0
	for i := range nums {
		if step < i {
			return false
		}
		step = max(step, i+nums[i])
	}
	return true
}
