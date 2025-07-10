package solution

func lengthOfLastWord(s string) int {
	var ans int
	right := len(s) - 1
	for right >= 0 {
		if s[right] == ' ' {
			if ans > 0 {
				return ans
			}
		} else {
			ans++
		}
		right--
	}

	return ans
}
