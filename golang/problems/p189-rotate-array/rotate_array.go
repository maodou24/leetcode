package rotatearray

func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}

	n := len(nums)

	temp := make([]int, k)
	copy(temp, nums[n-k:])

	copy(nums[k:], nums[:n-k+1])
	copy(nums[:k], temp)
}
