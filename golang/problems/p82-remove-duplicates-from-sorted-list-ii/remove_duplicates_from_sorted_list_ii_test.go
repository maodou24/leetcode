package solution

import (
	"leetcode/golang/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicatesFromSortedListIi(t *testing.T) {
	type data struct {
		head   []int
		except []int
	}

	testData := []data{
		{
			head:   []int{1, 2, 3, 3, 4, 4, 5},
			except: []int{1, 2, 5},
		},
		{
			head:   []int{1, 1, 1, 2, 3},
			except: []int{2, 3},
		},
		{
			head:   []int{1, 1},
			except: []int{},
		},
	}

	for _, tdata := range testData {
		assert.EqualValues(t, utils.SliceToListNode(tdata.except), deleteDuplicates(utils.SliceToListNode(tdata.head)))

		assert.EqualValues(t, utils.SliceToListNode(tdata.except), deleteDuplicates2(utils.SliceToListNode(tdata.head)))
	}
}
