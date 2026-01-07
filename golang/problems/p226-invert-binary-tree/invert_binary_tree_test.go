package solution

import (
	"leetcode/golang/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSummaryRanges(t *testing.T) {
	type data struct {
		nums   []int

		except []int
	}

	testData := []data{
		{
			nums:   []int{4,2,7,1,3,6,9},
			except: []int{4,7,2,9,6,3,1},
		},
		{
			nums:   []int{2,1,3},
			except: []int{2,3,1},
		},
		{
			nums:   []int{},
			except: []int{},
		},
	}

	for _, d := range testData {
		assert.Equal(t, utils.IntSliceToBinaryTree(d.except), invertTree(utils.IntSliceToBinaryTree(d.nums)))
	}
}
