// https://leetcode.com/problems/heaters/
package leetcode

import (
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	i := 0
	start := 0
	maxRadius := int(math.Max(float64(heaters[0]-houses[0]), 0))
	for i < len(heaters) {
		if i+1 == len(heaters) {
			maxRadius = int(math.Max(float64(houses[len(houses)-1]-heaters[i]), float64(maxRadius)))
			break
		}
		cur := heaters[i]
		next := heaters[i+1]
		idx := sort.SearchInts(houses[start:], next) + start
		if idx == len(houses) {
			idx = len(houses) - 1
			if houses[idx] < cur {
				break
			}
		}
		mid := (cur + next) / 2
		if (cur+next)&1 == 1 {
			mid++
		}
		midIdx := sort.SearchInts(houses[start:idx+1], mid) + start
		if midIdx < len(houses) && houses[midIdx] <= next {
			maxRadius = int(math.Max(math.Abs(float64(houses[midIdx]-next)), float64(maxRadius)))
		}
		midIdx--
		if midIdx >= 0 && houses[midIdx] >= cur {
			maxRadius = int(math.Max(math.Abs(float64(houses[midIdx]-cur)), float64(maxRadius)))
		}
		start = idx
		i++
	}
	return maxRadius
}
