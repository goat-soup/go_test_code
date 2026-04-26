package work150

/*
从左到右遍历罗马数字

	每个字符和它右边字符比较：
		如果左 < 右 → 减去左
		否则 → 加上左
		最后把最后一个字符加上即可（因为没有右边可比较）
*/
func romanToInt(s string) int {
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res := 0
	for i := 0; i < len(s)-1; i++ {
		if roman[s[i]] < roman[s[i+1]] {
			res -= roman[s[i]]
		} else {
			res += roman[s[i]]
		}
	}
	res += roman[s[len(s)-1]]
	return res
}
