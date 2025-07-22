package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSummaryRanges(t *testing.T) {
	type data struct {
		nums   []int

		except []string
	}

	testData := []data{
		{
			nums:   []int{0,1,2,4,5,7},
			except: []string{"0->2","4->5","7"},
		},
		{
			nums:   []int{0,2,3,4,6,8,9},
			except: []string{"0","2->4","6","8->9"},
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, summaryRanges(d.nums))
	}
}
