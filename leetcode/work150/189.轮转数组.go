package work150

func rotate(nums []int, k int) {
	temp := make([]int, len(nums))

	for i, v := range nums {
		temp[(i+k)%len(nums)] = v
	}
	copy(nums, temp)
}
