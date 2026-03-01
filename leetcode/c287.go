package leetcode

func FindDuplicate(nums []int) int {
	temp := make(map[int]bool, len(nums)-1)
	for _, v := range nums {
		if temp[v] {
			return v
		}
		temp[v] = true
	}
	return nums[len(nums)-1]
}
