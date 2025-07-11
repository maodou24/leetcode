package solution

func strStr(haystack string, needle string) int {
	n, h := len(needle), len(haystack)

	for i := range haystack {
		if i+n <= h && haystack[i] == needle[0] && haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}
