package leetcode

import (
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)
	i, j := 0, len(nums)-1
	for i < len(nums) {
		if nums[i] != sorted[i] {
			break
		}
		i++
	}
	if i == len(nums) {
		return 0
	}

	for j >= 0 {
		if nums[j] != sorted[j] {
			break
		}
		j--
	}

	return j - i + 1
}
