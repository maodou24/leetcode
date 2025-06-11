package majorityelement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	type data struct {
		nums   []int
		except int
	}

	testData := []data{
		{
			nums:   []int{3, 2, 3},
			except: 3,
		},
		{
			nums:   []int{2, 2, 1, 1, 1, 2, 2},
			except: 2,
		},
	}

	for _, d := range testData {
		assert.EqualValues(t, d.except, majorityElement(d.nums))
	}
}
