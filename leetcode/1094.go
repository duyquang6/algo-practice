// https://leetcode.com/problems/car-pooling/
package leetcode

func carPooling(trips [][]int, capacity int) bool {
	stops := [1001]int{}
	for _, trip := range trips {
		stops[trip[1]] += trip[0]
		stops[trip[2]] -= trip[0]
	}
	for _, v := range stops {
		capacity -= v
		if capacity < 0 {
			return false
		}
	}
	return true
}
