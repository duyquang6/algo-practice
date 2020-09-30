package leetcode

import "sort"

func binarySearch(arr []int, target int) int {
	return sort.SearchInts(arr, target)
}
