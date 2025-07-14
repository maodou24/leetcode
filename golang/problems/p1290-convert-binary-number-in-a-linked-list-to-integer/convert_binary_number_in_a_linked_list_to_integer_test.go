package solution

import (
	"leetcode/golang/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDecimalValue(t *testing.T) {
	type args struct {
		s []int

		except int
	}

	testData := []args{
		{
			s:      []int{1, 0, 1},
			except: 5,
		},
		{
			s:      []int{0},
			except: 0,
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, getDecimalValue2(utils.SliceToListNode(d.s)))
	}
}
