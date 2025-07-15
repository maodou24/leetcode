package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJump(t *testing.T) {
	type data struct {
		nums   []int
		except int
	}

	testData := []data{
		{
			nums:   []int{2, 3, 1, 1, 4},
			except: 2,
		},
		{
			nums:   []int{2, 3, 0, 1, 4},
			except: 2,
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, jump(d.nums))
		assert.Equal(t, d.except, jump2(d.nums))
	}
}
