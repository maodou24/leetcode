package solution

import "strings"

func isPalindrome(s string) bool {
	isAlpha := func(c byte) bool {
		if c >= 'a' && c <= 'z' {
			return true
		}

		if c >= '0' && c <= '9' {
			return true
		}

		return false
	}
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		if !isAlpha(s[l]) {
			l++
			continue
		} else if !isAlpha(s[r]) {
			r--
			continue
		}

		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
