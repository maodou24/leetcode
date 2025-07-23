package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSummaryRanges(t *testing.T) {
	type data struct {
		nums [][]int

		except [][]int
	}

	testData := []data{
		{
			nums:   [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			except: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			nums:   [][]int{{1, 4}, {4, 5}},
			except: [][]int{{1, 5}},
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, merge(d.nums))
	}
}
