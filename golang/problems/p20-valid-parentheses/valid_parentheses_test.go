package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	type data struct {
		s string

		except bool
	}

	testData := []data{
		{
			s:      "()",
			except: true,
		},
		{
			s:      "()[]{}",
			except: true,
		},
		{
			s:      "(]",
			except: false,
		},
		{
			s:      "([])",
			except: true,
		},
		{
			s:      "([)]",
			except: false,
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, isValid(d.s))
	}
}
