package solution

import (
	"leetcode/golang/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	head := utils.SliceToListNode([]int{1, 2, 3, 4, 5})

	r := reverseList(head)

	assert.Equal(t, utils.SliceToListNode([]int{5, 4, 3, 2, 1}), r)
}
