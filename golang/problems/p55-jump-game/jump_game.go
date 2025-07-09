package solution

func canJump(nums []int) bool {
	var next int

	lastIdx := len(nums) - 1
	for i := range nums {
		if next < i {
			break
		}
		next = max(next, i+nums[i])
		if next >= lastIdx {
			return true
		}
	}
	return false
}
