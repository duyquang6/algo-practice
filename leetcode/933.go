// https://leetcode.com/problems/number-of-recent-calls/
package leetcode

import "sort"

type RecentCounter struct {
	arr []int
}

func Constructor() RecentCounter {
	return RecentCounter{arr: make([]int, 0, 3000)}
}

func (this *RecentCounter) Ping(t int) int {
	this.arr = append(this.arr, t)
	idx := sort.SearchInts(this.arr, t-3000)
	this.arr = this.arr[idx:]
	return len(this.arr)
}
