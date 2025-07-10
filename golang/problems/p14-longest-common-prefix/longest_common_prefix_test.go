package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	type data struct {
		strs []string

		except string
	}

	testData := []data{
		{
			strs:   []string{"flower", "flow", "flight"},
			except: "fl",
		},
		{
			strs:   []string{"dog", "racecar", "car"},
			except: "",
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, longestCommonPrefix(d.strs))
	}
}
