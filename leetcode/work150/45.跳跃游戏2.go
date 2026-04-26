package work150

func jump(nums []int) int {
	// 下一跳结束的位置
	end := 0
	// 总步长jump
	step := 0
	// 能跳的最长距离
	max_dst := 0
	for i := 0; i < len(nums)-1; i++ {
		// 在当前这个位置我能跳的最远的位置
		max_dst = max(max_dst, i+nums[i])
		// 如果到达了最远位置，则更新下一个的位置，并视为一跳
		if i == end {
			end = max_dst
			step++
		}
	}
	return step
}
