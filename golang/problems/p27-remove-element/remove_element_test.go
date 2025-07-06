package p27removeelement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	type data struct {
		nums   []int
		val    int
		except int
	}

	testData := []data{
		{
			nums: []int{3, 2, 2, 3},
			val:  3,

			except: 2,
		},
		{
			nums:   []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:    2,
			except: 5,
		},
	}

	for _, d := range testData {
		assert.EqualValues(t, d.except, removeElement(d.nums, d.val))
	}
}
