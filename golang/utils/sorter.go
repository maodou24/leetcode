package utils

import "sort"

type stringMatrixSorter [][]string

func (s stringMatrixSorter) Len() int      { return len(s) }
func (s stringMatrixSorter) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s stringMatrixSorter) Less(i, j int) bool {
	m := min(len(s[i]), len(s[j]))
	// first compare by fisrt character, if they are equal, compare by the rest
	// next compare length
	for k := 0; k < m; k++ {
		if s[i][k] < s[j][k] {
			return true
		} else if s[i][k] > s[j][k] {
			return false
		}
	}
	return len(s[i]) < len(s[j])
}

func SortStringMatrix(matrix [][]string) {
	for i := range matrix {
		sort.Strings(matrix[i]) // sort each row
	}
	sort.Sort(stringMatrixSorter(matrix))
}
