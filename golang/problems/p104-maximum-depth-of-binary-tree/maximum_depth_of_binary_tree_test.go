package solution

import (
	"leetcode/golang/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDepth(t *testing.T) {
	type data struct {
		nums   []int
		except int
	}

	testData := []data{
		{
			nums:   []int{3, 9, 20, -1, -1, 15, 7},
			except: 3,
		},
		{
			nums:   []int{1, -1, 2},
			except: 2,
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, maxDepth(utils.IntSliceToBinaryTree(d.nums)))
	}
}
