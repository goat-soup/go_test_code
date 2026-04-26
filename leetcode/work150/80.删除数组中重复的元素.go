package work150

func removeDuplicates2(nums []int) int {
	m := make(map[int]int)
	left := 0
	for _, n := range nums {
		if v, ok := m[n]; ok && v == 2 {
			continue
		}
		nums[left] = n
		left++
		m[n]++
	}
	return left
}
