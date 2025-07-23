package solution

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var merged [][]int
	for i := range intervals {
		if len(merged) == 0 || merged[len(merged)-1][1] < intervals[i][0] {
			merged = append(merged, intervals[i])
		} else {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], intervals[i][1])
		}

	}
	return merged
}
