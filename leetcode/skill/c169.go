package skill

import "sort"

// 多数元素
func MajorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
