package solution

import . "leetcode/golang/base"

func getDecimalValue(head *ListNode) int {
	var ans int

	var s []int
	for node := head; node != nil; node = node.Next {
		s = append(s, node.Val)
	}

	n := len(s)
	for i := range s {
		ans += s[n-i-1] << i
	}

	return ans
}

func getDecimalValue2(head *ListNode) int {
	var ans int

	for node := head; node != nil; node = node.Next {
		ans = ans<<1 | node.Val
	}

	return ans
}
