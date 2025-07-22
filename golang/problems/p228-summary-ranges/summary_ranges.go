package solution

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	var ans []string
	var prev int

	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] <= 1 {
			continue
		}

		if nums[i-1] == nums[prev] {
			ans = append(ans, fmt.Sprintf("%v", nums[prev]))
		} else {
			ans = append(ans, fmt.Sprintf("%v->%v", nums[prev], nums[i-1]))
		}
		prev = i
	}

	if nums[n-1] == nums[prev] {
		ans = append(ans, fmt.Sprintf("%v", nums[prev]))
	} else {
		ans = append(ans, fmt.Sprintf("%v->%v", nums[prev], nums[n-1]))
	}

	return ans
}
