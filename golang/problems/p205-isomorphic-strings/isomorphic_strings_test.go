package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIsomorphic(t *testing.T) {
	type data struct {
		s string
		t string

		except bool
	}

	testData := []data{
		{
			s:      "egg",
			t:      "add",
			except: true,
		},
		{
			s:      "foo",
			t:      "bar",
			except: false,
		},
		{
			s:      "paper",
			t:      "title",
			except: true,
		},
		{
			s:      "badc",
			t:      "baba",
			except: false,
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, isIsomorphic(d.s, d.t))
	}
}
