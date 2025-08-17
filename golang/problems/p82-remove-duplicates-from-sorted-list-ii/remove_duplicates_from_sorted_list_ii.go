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
func deleteDuplicates(head *ListNode) *ListNode {
	ans := &ListNode{Val: -101, Next: head}

	prev := ans
	node := head

	for node != nil {
		next := node.Next
		if next == nil || next.Val != node.Val {
			prev.Next = node
			prev = prev.Next
		}

		for next != nil && next.Val == node.Val {
			next = next.Next
		}
		node = next
	}

	prev.Next = nil // cut off the tail of the list if it has duplicates

	return ans.Next
}
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}

	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}
