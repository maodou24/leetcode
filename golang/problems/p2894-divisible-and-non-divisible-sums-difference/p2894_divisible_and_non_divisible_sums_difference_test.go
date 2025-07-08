package solution

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDifferenceOfSums(t *testing.T) {
	type data struct {
		n      int
		m      int
		except int
	}

	testData := []data{
		{
			n:      10,
			m:      3,
			except: 19,
		},
		{
			n:      5,
			m:      6,
			except: 15,
		},
		{
			n:      5,
			m:      1,
			except: -15,
		},
	}

	for _, tdata := range testData {
		assert.Equal(t, tdata.except, differenceOfSums(tdata.n, tdata.m))
	}
}
