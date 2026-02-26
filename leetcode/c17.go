package leetcode

// 回溯
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

var digitToLetter = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func LetterCombinations(digits string) []string {
	var res []string

	var backtrack func(int, string)

	backtrack = func(index int, prefix string) {
		if index == len(digits) {
			res = append(res, prefix)
			return
		}
		letterStr := digitToLetter[digits[index]]
		for _, i := range letterStr {
			backtrack(index+1, prefix+string(i))
		}
	}
	backtrack(0, "")
	return res
}
