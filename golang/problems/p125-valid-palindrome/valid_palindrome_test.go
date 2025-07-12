package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	type data struct {
		s      string
		except bool
	}

	testData := []data{
		{
			s:      "A man, a plan, a canal: Panama",
			except: true,
		},
		{
			s:      "race a car",
			except: false,
		},
		{
			s:      " ",
			except: true,
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, isPalindrome(d.s))
	}
}
