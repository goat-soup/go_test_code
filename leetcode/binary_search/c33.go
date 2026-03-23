package binarysearch

func Search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[l] <= nums[mid] { // 这个代表左边有序
			if nums[l] <= target && nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[r] >= target && nums[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
