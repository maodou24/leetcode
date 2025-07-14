package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsHappy(t *testing.T) {
	type data struct {
		n      int
		except bool
	}

	testData := []data{
		{
			n:      19,
			except: true,
		},
		{
			n:      2,
			except: false,
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, isHappy(d.n))
	}
}
