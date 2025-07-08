package solution

func maxProfit(prices []int) int {
    var ans int

    for i := 1; i < len(prices); i++ {
        ans = max(ans, ans + prices[i] - prices[i-1])
    }
    return ans
}
