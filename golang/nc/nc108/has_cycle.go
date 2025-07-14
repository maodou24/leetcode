package solution

import . "leetcode/golang/base"

func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]struct{})
	for node := head; node != nil; node = node.Next {
		_, ok := m[node]
		if ok {
			return true
		}

		m[node] = struct{}{}
	}

	return false
}
