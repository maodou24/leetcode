package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	type data struct {
		nums   []int
		k      int
		except bool
	}

	testData := []data{
		{
			nums:   []int{1, 2, 3, 1},
			k:      3,
			except: true,
		},
		{
			nums:   []int{1, 0, 1, 1},
			k:      1,
			except: true,
		},
		{
			nums:   []int{1, 2, 3, 1, 2, 3},
			k:      2,
			except: false,
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, containsNearbyDuplicate(d.nums, d.k))
	}
}
