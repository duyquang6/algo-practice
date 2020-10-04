// https://leetcode.com/problems/third-maximum-number/
package leetcode

import "math"

func thirdMax(nums []int) int {
	max1, max2, max3 := math.MinInt32, math.MinInt32, math.MinInt32
	isMaxAvail := []bool{false, false, false}
	for _, num := range nums {
		if num > max1 || !isMaxAvail[0] {
			if isMaxAvail[0] {
				if isMaxAvail[1] {
					max3 = max2
					isMaxAvail[2] = true
				}
				max2 = max1
				isMaxAvail[1] = true
			}
			max1 = num
			isMaxAvail[0] = true
		} else if num < max1 && (num > max2 || !isMaxAvail[1]) {
			if isMaxAvail[1] {
				max3 = max2
				isMaxAvail[2] = true
			}
			max2 = num
			isMaxAvail[1] = true
		} else if num < max2 && (num > max3 || !isMaxAvail[2]) {
			max3 = num
			isMaxAvail[2] = true
		}
	}

	if !isMaxAvail[2] {
		return max1
	}

	return max3
}
