package solution

func romanToInt(s string) int {
	var m = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var ans int
	for i := 0; i < len(s); i++ {
		// 减法规则：当左侧的符号值小于右侧时，需用右侧值减去左侧值
		if i + 1 < len(s) && m[s[i]] < m[s[i+1]] {
			ans -= m[s[i]]
		} else {
			ans += m[s[i]]
		}
	}
	return ans
}
