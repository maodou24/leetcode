package solution

import (
	"leetcode/golang/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartitionList(t *testing.T) {
	type data struct {
		head   []int
		x      int
		except []int
	}

	testData := []data{
		{
			head:   []int{1, 4, 3, 2, 5, 2},
			x:      3,
			except: []int{1, 2, 2, 4, 3, 5},
		},
		{
			head:   []int{2, 1},
			x:      2,
			except: []int{1, 2},
		},
	}

	for _, tdata := range testData {
		assert.EqualValues(t, utils.SliceToListNode(tdata.except), partition(utils.SliceToListNode(tdata.head), tdata.x))
	}
}
