package solution

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)

	for _, v := range magazine {
		m[v] = m[v] + 1
	}

	for _, v := range ransomNote {
		if m[v] <= 0 {
			return false
		}
		m[v] = m[v] - 1
	}
	return true
}