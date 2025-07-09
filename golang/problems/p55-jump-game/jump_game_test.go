package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanJump(t *testing.T) {
	type data struct {
		nums   []int
		except bool
	}

	testData := []data{
		{
			nums:   []int{2, 3, 1, 1, 4},
			except: true,
		},
		{
			nums:   []int{3, 2, 1, 0, 4},
			except: false,
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, canJump(d.nums))
	}
}
