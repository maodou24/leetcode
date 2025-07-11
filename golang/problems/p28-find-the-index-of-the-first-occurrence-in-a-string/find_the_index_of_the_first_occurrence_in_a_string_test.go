package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	type data struct {
		haystack string
		needle   string
		except   int
	}

	testData := []data{
		{
			haystack: "sadbutsad",
			needle:   "sad",
			except:   0,
		},
		{
			haystack: "leetcode",
			needle:   "leeto",
			except:   -1,
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, strStr(d.haystack, d.needle))
	}
}
