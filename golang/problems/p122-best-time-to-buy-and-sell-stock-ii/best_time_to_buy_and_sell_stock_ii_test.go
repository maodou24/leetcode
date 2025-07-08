package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	type data struct {
		nums   []int
		except int
	}

	testData := []data{
		{
			nums:   []int{7, 1, 5, 3, 6, 4},
			except: 7,
		},
		{
			nums:   []int{1, 2, 3, 4, 5},
			except: 4,
		},
		{
			nums:   []int{7, 6, 4, 3, 1},
			except: 0,
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, maxProfit(d.nums))
	}
}
