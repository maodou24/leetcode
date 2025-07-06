package solution

import . "leetcode/golang/base"

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	prev, next := head, head.Next

	for next != nil {
		temp := next.Next

		next.Next = prev
		prev = next
		next = temp
	}

	head.Next = nil
	return prev
}
