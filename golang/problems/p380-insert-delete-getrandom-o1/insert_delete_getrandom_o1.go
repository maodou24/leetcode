package solution

import (
	"math/rand"
)

type RandomizedSet struct {
	m    map[int]int // val: index
	keys []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m:    make(map[int]int),
		keys: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.m[val]
	if ok {
		return false
	}
	this.keys = append(this.keys, val)
	this.m[val] = len(this.keys) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.m[val]
	if ok {
		lastIdx := len(this.keys) - 1
		if idx < lastIdx {
			lastvalue := this.keys[lastIdx]
			this.keys[idx] = lastvalue
			this.m[lastvalue] = idx
		}

		delete(this.m, val)
		this.keys = this.keys[:lastIdx]
		return true
	}

	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.keys[rand.Intn(len(this.keys))]
}
