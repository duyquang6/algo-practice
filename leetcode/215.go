// https://leetcode.com/problems/kth-largest-element-in-an-array/
package leetcode

import (
	"fmt"
	"math/rand"
)

func findKthLargest(nums []int, k int) int {
	partition(nums, 0, len(nums)-1)
	return 0
}

func partition(list []int, lo, hi int) {
	pivotIdx := rand.Intn(hi-lo+1) + lo
	// Swap pivot to first elem
	list[lo], list[pivotIdx] = list[pivotIdx], list[lo]

	a, b := lo+1, hi

	pivot := list[lo]
	for a < b {
		for a < b && list[a] <= pivot {
			a++
		}
		for a < b && list[b] > pivot {
			b--
		}
		list[a], list[b] = list[b], list[a]
	}

	a--
	list[lo], list[a] = list[a], list[lo]

	fmt.Println("pivot:", pivot, "arr:", list, "a:", a)
}
