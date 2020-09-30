package leetcode

import "sort"

func lowerbound(arr []int, target int) int {
	return sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})
}
