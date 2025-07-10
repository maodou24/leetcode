package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	type data struct {
		s      string
		except int
	}

	testData := []data{
		{
			s:      "III",
			except: 3,
		},
		{
			s:      "LVIII",
			except: 58,
		},
		{
			s:      "MCMXCIV",
			except: 1994,
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, romanToInt(d.s))
	}
}
