package leetcode

func CanJump(nums []int) bool {
	reach := 0
	for i := range nums {
		if i > reach {
			return false
		}
		reach = max(reach, i+nums[i])
	}
	return true
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
