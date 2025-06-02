package addtwonumbers

import (
	. "leetcode/golang/base"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	header := &ListNode{}
	node := header
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		node.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		node = node.Next
	}

	if carry > 0 {
		node.Next = &ListNode{Val: carry}
	}

	return header.Next
}
