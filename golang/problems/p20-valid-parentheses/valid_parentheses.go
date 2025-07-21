package solution

func isValid(s string) bool {
	queue := make([]rune, 0)
	for _, v := range s {
		lastIdx := len(queue) - 1
		if lastIdx >= 0 && (queue[lastIdx] == '(' && v == ')' || queue[lastIdx] == '[' && v == ']' ||
			queue[lastIdx] == '{' && v == '}') {
			queue = queue[:lastIdx]
			continue
		}

		queue = append(queue, v)
	}
	return len(queue) == 0
}
