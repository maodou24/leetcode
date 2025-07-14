package solution

func isHappy(n int) bool {
	m := make(map[int]struct{})

	for n != 1 {
		_, ok := m[n]
		if ok {
			return false
		}

		m[n] = struct{}{}
		var sum int
		for v := n; v != 0; v = v / 10 {
			sum += (v % 10) * (v % 10)
		}
		n = sum
	}
	return true
}
