package work150

func removeDuplicates(nums []int) int {
	m := make(map[int]struct{})
	left := 0
	for _, n := range nums {
		if _, ok := m[n]; ok {
			continue
		}
		nums[left] = n
		m[n] = struct{}{}
		left++
	}
	return left
}
