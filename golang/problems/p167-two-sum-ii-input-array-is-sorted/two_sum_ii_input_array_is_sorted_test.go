package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	type data struct {
		numbers []int
		target  int
		except  []int
	}

	testData := []data{
		{
			numbers: []int{2, 7, 11, 15},
			target:  9,
			except:  []int{1, 2},
		},
		{
			numbers: []int{2, 3, 4},
			target:  6,
			except:  []int{1, 3},
		},
		{
			numbers: []int{-1, 0},
			target:  -1,
			except:  []int{1, 2},
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, twoSum(d.numbers, d.target))
	}
}
