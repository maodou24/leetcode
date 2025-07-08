package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	type data struct {
		nums string

		except int
	}

	testData := []data{
        {
			nums:   " ",
			except: 1,
		},
		{
			nums:   "abcabcbb",
			except: 3,
		},
		{
			nums:   "bbbbb",
			except: 1,
		},
		{
			nums:   "pwwkew",
			except: 3,
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, lengthOfLongestSubstring(d.nums))
	}
}
