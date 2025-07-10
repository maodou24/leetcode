package solution

func longestCommonPrefix(strs []string) string {
	for i := range strs[0] {
		for j := 1; j < len(strs); j++ {
			if i > len(strs[j]) || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
