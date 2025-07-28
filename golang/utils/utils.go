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

func NewLoopListNode(input []int, idx int) *base.ListNode {
	if len(input) == 0 {
		return nil
	}
	result := &base.ListNode{}

	var loop *base.ListNode
	head := result
	for i := range input {
		head.Val = input[i]
		if i == idx {
			loop = head
		}
		if i < len(input)-1 {
			head.Next = &base.ListNode{}
			head = head.Next
		}
	}

	head.Next = loop
	return result
}

func IntSliceToBinaryTree(input []int) *base.TreeNode {
	s := make([]int, len(input)+1)
	copy(s[1:], input)
	return binaryTree(s, 1)
}

func binaryTree(input []int, index int) *base.TreeNode {
	if index < len(input) {
		val := input[index]
		if val >= 0 {
			return &base.TreeNode{
				Val:   input[index],
				Left:  binaryTree(input, index*2),
				Right: binaryTree(input, index*2+1),
			}
		}
	}
	return nil
}