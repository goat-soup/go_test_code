package work150

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	n, m := len(haystack), len(needle)
	if n < m {
		return -1
	}
	for i := 0; i <= n-m; i++ {
		if haystack[i:i+m] == needle {
			return i
		}
	}
	return -1
}
