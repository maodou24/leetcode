package solution

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	var ans int
	i := 0
	for i < len(points) {
		j := i + 1
		r := points[i][1]
		for j < len(points) && r >= points[j][0] {
			r = min(r, points[j][1])
			j++
		}
		i = j

		ans++
	}

	return ans
}
