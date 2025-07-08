package solution

func maxProfit(prices []int) int {
	var ans int

	prev := prices[0]
	for i := 1; i < len(prices); i++ {
		if prev > prices[i] {
			prev = prices[i]
		}
		ans = max(ans, prices[i]-prev)
	}

	return ans
}
