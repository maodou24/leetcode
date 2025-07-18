package solution

import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)

	var ans int
	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] <= ans {
			break
		}
		ans++
	}
	return ans
}
