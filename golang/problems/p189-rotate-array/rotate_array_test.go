package rotatearray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	type data struct {
		nums   []int
		k      int
		except []int
	}

	testData := []data{
		{
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			k:      3,
			except: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			nums: []int{-1, -100, 3, 99},
			k:    2,

			except: []int{3, 99, -1, -100},
		},
	}

	for _, d := range testData {
		rotate(d.nums, d.k)
		assert.EqualValues(t, d.except, d.nums)
	}
}
