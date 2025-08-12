package solution

import (
	"leetcode/golang/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	type data struct {
		nums []string

		except [][]string
	}

	testData := []data{
		{
			nums:   []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			except: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			nums:   []string{},
			except: [][]string{},
		},
		{
			nums:   []string{"a"},
			except: [][]string{{"a"}},
		},
	}

	for _, d := range testData {
		ans := groupAnagrams(d.nums)
		utils.SortStringMatrix(ans)

		utils.SortStringMatrix(d.except)

		assert.EqualValues(t, d.except, ans)
	}

}
