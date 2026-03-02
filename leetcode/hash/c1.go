package leetcode

func TwoSum(nums []int, target int) []int {
	temp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := temp[target-nums[i]]; ok {
			return []int{i, v}
		}
		temp[nums[i]] = i
	}
	return []int{}
}
