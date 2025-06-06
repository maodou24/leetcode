package addtwonumbers

import (
	"leetcode/golang/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	type data struct {
		input1 []int
		input2 []int
		except []int
	}

	testData := []data{
		{
			input1: []int{2, 4, 3},
			input2: []int{5, 6, 4},
			except: []int{7, 0, 8},
		},
		{
			input1: []int{0},
			input2: []int{0},
			except: []int{0},
		},
		{
			input1: []int{9, 9, 9, 9, 9, 9, 9},
			input2: []int{9, 9, 9, 9},
			except: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, tdata := range testData {
		assert.EqualValues(t, utils.SliceToListNode(tdata.except), addTwoNumbers(utils.SliceToListNode(tdata.input1), utils.SliceToListNode(tdata.input2)))	
	}
}