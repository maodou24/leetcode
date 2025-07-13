package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	type data struct {
		ransomNote string
		magazine   string

		except bool
	}

	testData := []data{
		{
			ransomNote: "a",
			magazine:   "b",
			except:     false,
		},
		{
			ransomNote: "aa",
			magazine:   "ab",
			except:     false,
		},
		{
			ransomNote: "aa",
			magazine:   "aab",
			except:     true,
		},
	}

	for _, d := range testData {
		assert.EqualValues(t, d.except, canConstruct(d.ransomNote, d.magazine))
	}
}
