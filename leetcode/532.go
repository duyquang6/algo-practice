// https://leetcode.com/problems/k-diff-pairs-in-an-array/
package leetcode

import (
	"sort"
)

func findPairs(nums []int, k int) int {
	sort.Ints(nums)
	i, j := 0, 1
	res := 0
	for j < len(nums) {
		if i >= j {
			j = i + 1
			continue
		}
		delta := nums[j] - nums[i]
		if delta == k {
			res++
		}

		if delta <= k {
			for j < len(nums)-1 && nums[j] == nums[j+1] {
				j++
			}
			j++
		}

		if delta > k {
			for i < len(nums)-1 && nums[i] == nums[i+1] {
				i++
			}
			i++
		}
	}
	return res
}
