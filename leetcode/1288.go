// https://leetcode.com/explore/challenge/card/october-leetcoding-challenge/559/week-1-october-1st-october-7th/3483/
package leetcode

import (
	"sort"
)

func removeCoveredIntervals(intervals [][]int) int {
	isCover := func(a, b []int) bool {
		return a[0] <= b[0] && a[1] >= b[1]
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	skipMap := map[int]struct{}{}
	for i := 0; i < len(intervals); i++ {
		if _, ok := skipMap[i]; ok {
			continue
		}
		for j := i + 1; j < len(intervals); j++ {
			if _, ok := skipMap[j]; ok {
				continue
			}
			if isCover(intervals[i], intervals[j]) {
				skipMap[j] = struct{}{}
			}
		}
	}

	return len(intervals) - len(skipMap)
}
