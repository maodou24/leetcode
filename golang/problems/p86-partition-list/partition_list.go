package solution

import (
	. "leetcode/golang/base"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	dummy := &ListNode{Next: head}

	right := &ListNode{Next: head}

	dummyR := right

	cur := dummy
	for cur != nil {
		n := cur

		for n.Next != nil && n.Next.Val >= x {
			right.Next = n.Next
			right = right.Next
			n = n.Next
		}

		cur.Next = n.Next
		if n.Next == nil {
			right.Next = nil
			cur.Next = dummyR.Next
			break
		}
		cur = cur.Next
	}

	return dummy.Next
}
