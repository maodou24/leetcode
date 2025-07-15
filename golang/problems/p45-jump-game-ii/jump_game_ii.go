package solution

func jump(nums []int) int {
	ans, farthest, curr := 0, 0, 0

	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])

		if i == curr {
			ans++
			curr = farthest
		}
	}
	return ans
}

func jump2(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	n := len(nums)
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i]+i)
	}

	ans := 1
	for i := 0; i < len(dp); {
		if dp[i] >= n-1 {
			return ans
		}
		ans++
		i = dp[i]
	}

	return ans
}
