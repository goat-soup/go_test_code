package work150

/*
核心思路：从大到小贪心匹配：
从最大的罗马值开始，如果当前整数 >= 这个值
就减去这个值，并把对应字符加到结果中
重复直到整数变为 0
*/
func intToRoman(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	sym := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res := ""
	for i := 0; i < len(val); i++ {
		for val[i] <= num {
			num -= val[i]
			res += sym[i]
		}
	}
	return res
}
