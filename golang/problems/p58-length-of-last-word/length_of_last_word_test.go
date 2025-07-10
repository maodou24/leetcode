package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLastWord(t *testing.T) {
	type data struct {
		s   string
		except int
	}

	testData := []data{
		{
			s:   "Hello World",
			except: 5,
		},
		{
			s:   "   fly me   to   the moon  ",
			except: 4,
		},
		{
			s:   "luffy is still joyboy",
			except: 6,
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, lengthOfLastWord(d.s))
	}
}
