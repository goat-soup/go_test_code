package binarysearch

// 右边一定包含最小值
func FindMin(nums []int) int {
	r, l := len(nums)-1, 0
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}
