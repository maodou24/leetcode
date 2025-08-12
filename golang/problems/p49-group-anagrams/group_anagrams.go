package solution

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string, 0)
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		m[string(s)] = append(m[string(s)], str)
	}

	ans := make([][]string, 0, len(m))
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}
