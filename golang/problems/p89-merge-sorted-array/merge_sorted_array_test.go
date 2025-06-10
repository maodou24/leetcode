package mergesortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestMerge(t *testing.T) {
	testDatas := []struct {
		nums1   []int
        m int
		nums2 []int
        n int

        except []int
	}{
		{
			nums1:   []int{1,2,3,0,0,0},
			m: 3,
			nums2: []int{2,5,6},
            n: 3,
            except: []int{1,2,2, 3,5,6},
		},
		{
			nums1:   []int{},
			m: 0,
			nums2: []int{},
            n: 0,
            except: []int{},
		},
		{
			nums1:   []int{0},
			m: 0,
			nums2: []int{1},
            n: 1,
            except: []int{1},
		},
	}


	for _, d := range testDatas {
        merge(d.nums1, d.m, d.nums2, d.n)
		assert.EqualValues(t, d.except, d.nums1)	
	}
}
