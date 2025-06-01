package twosum

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	testDatas := []struct {
		nums   []int
		target int
		except []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			except: []int{0, 1},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			except: []int{1, 2},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			except: []int{0, 1},
		},
	}

	for _, d := range testDatas {
		tmp := twoSum(d.nums, d.target)
		sort.Ints(d.except)
		sort.Ints(tmp)
		assert.Equal(t, d.except, tmp)
	}
}
