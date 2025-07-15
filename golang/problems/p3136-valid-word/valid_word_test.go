package solution

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	type data struct {
		word string

		except bool
	}

	fmt.Println('a', 'A')

	testData := []data{
		{
			word:   "234Adas",
			except: true,
		},
		{
			word:   "b3",
			except: false,
		},
		{
			word:   "a3$e",
			except: false,
		},
		{
			word:   "AhI",
			except: true,
		},
	}

	for _, d := range testData {
		assert.Equal(t, d.except, isValid(d.word))
	}
}
