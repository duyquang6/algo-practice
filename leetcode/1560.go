// https://leetcode.com/problems/most-visited-sector-in-a-circular-track/
package leetcode

import "math"

func mostVisited(n int, rounds []int) []int {
	counter := make([]int, n)
	for m := 0; m < len(rounds)-1; m++ {
		start, end := rounds[m], rounds[m+1]
		if m > 0 {
			start++
		}
		for (start-1)%n != end-1 {
			counter[(start-1)%n]++
			start++
		}
		counter[end-1]++
	}

	res := []int{}
	max := math.MinInt32
	for i, v := range counter {
		if max == v {
			res = append(res, i+1)
		}
		if v > max {
			max = v
			res = []int{i + 1}
		}
	}

	return res
}

func optimizedMostVisited(n int, rounds []int) []int {
	start, end := rounds[0], rounds[len(rounds)-1]
	res := []int{}
	if start <= end {
		for i := start; i <= end; i++ {
			res = append(res, i)
		}
		return res
	}
	for i := 1; i <= end; i++ {
		res = append(res, i)
	}
	for i := start; i <= n; i++ {
		res = append(res, i)
	}
	return res
}
