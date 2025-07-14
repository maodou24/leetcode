package solution

func isIsomorphic(s string, t string) bool {
	mapping1 := make(map[byte]byte)
	mapping2 := make(map[byte]byte)
	for i := range s {
		v1, ok1 := mapping1[s[i]]
		v2, ok2 := mapping2[t[i]]
		if ok1 || ok2 {
			if v1 != t[i] || v2 != s[i] {
				return false
			}
		}

		mapping1[s[i]] = t[i]
		mapping2[t[i]] = s[i]
	}
	return true
}
