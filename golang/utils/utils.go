package utils

import "leetcode/golang/base"

func SliceToListNode(input []int) *base.ListNode {
	if len(input) == 0 {
		return nil
	}
	result := &base.ListNode{}
	head := result
	for i := range input {
		head.Val = input[i]
		if i < len(input)-1 {
			head.Next = &base.ListNode{}
			head = head.Next
		}
	}
	return result
}
