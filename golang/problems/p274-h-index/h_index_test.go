package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHIndex(t *testing.T) {
	type data struct {
		citations []int
		except    int
	}

	testData := []data{
		{
			citations: []int{3, 0, 6, 1, 5},
			except:    3,
		},
		{
			citations: []int{1, 3, 1},
			except:    1,
		},
		{
			citations: []int{0},
			except:    0,
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, hIndex(d.citations))
	}
}
