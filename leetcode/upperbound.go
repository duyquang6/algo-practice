package leetcode

import "sort"

func upperbound(arr []int, target int) int {
	return sort.Search(len(arr), func(i int) bool {
		return arr[i] > target
	})
}
