package solution

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i := range nums {
		idx, ok := m[nums[i]]
		if ok && i-idx <= k {
			return true
		}
		m[nums[i]] = i
	}

	return false
}
