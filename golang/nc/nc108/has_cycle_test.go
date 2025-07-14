package solution

import (
	"leetcode/golang/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasCycle(t *testing.T) {
	type args struct {
		s []int
		i int

		except bool
	}

	testData := []args{
		{s: []int{3, 2, 0, -4}, i: 1, except: true},
		{s: []int{1}, i: -1, except: false},
		{s: []int{-1, -7, 7, -4, 19, 6, -9, -5, -2, -5}, i: 6, except: true},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, hasCycle(utils.NewLoopListNode(d.s, d.i)))
	}
}
