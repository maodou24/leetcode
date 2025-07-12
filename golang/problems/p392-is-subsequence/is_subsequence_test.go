package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLastWord(t *testing.T) {
	type data struct {
		s      string
		t      string
		except bool
	}

	testData := []data{
		{
			s:      "abc",
			t:      "ahbgdc",
			except: true,
		},
		{
			s:      "axc",
			t:      "ahbgdc",
			except: false,
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, isSubsequence(d.s, d.t))
	}
}
