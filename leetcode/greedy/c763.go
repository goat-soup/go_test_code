package greedy

func PartitionLabels(s string) []int {
	start, end := 0, 0
	ch := make(map[byte]int)
	// 找到每个字母的最后一个位置
	for i := 0; i < len(s); i++ {
		ch[s[i]] = i
	}
	res := []int{}
	for i := 0; i < len(s); i++ {
		if ch[s[i]] > end {
			end = ch[s[i]]
		}
		if i == end {
			// 到达当前字母的极限，则记录此段
			res = append(res, end-start+1)
			start = i + 1
		}
	}
	return res
}
