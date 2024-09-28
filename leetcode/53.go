package leetcode

import (
	"slices"
)

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	res := make([][]int, 0, len(intervals))

	for _, v := range intervals {
		if len(res) == 0 {
			res = append(res, v)
			continue
		}

		last := res[len(res)-1]
		if v[0] <= last[1] {
			if v[1] > last[1] {
				res[len(res)-1][1] = v[1]
			}
			continue
		} else {
			res = append(res, v)
		}
	}

	return res
}
