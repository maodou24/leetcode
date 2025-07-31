package solution

import (
	"leetcode/golang/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoSortedLists(t *testing.T) {
	type data struct {
		input1 []int
		input2 []int
		except []int
	}

	testData := []data{
		{
			input1: []int{1, 2, 4},
			input2: []int{1, 3, 4},
			except: []int{1, 1, 2, 3, 4, 4},
		},
		{
			input1: []int{},
			input2: []int{},
			except: []int{},
		},
		{
			input1: []int{},
			input2: []int{0},
			except: []int{0},
		},
	}

	for _, tdata := range testData {
		assert.EqualValues(t, utils.SliceToListNode(tdata.except), mergeTwoLists(utils.SliceToListNode(tdata.input1), utils.SliceToListNode(tdata.input2)))
	}
}
