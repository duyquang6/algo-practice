package leetcode

import (
	"github.com/emirpasic/gods/trees/binaryheap"
	"github.com/emirpasic/gods/utils"
)

func kClosest(points [][]int, k int) [][]int {
	maxHeap := binaryheap.NewWith(func(a, b interface{}) int {
		ia := a.(int)
		ib := b.(int)
		distA := points[ia][0]*points[ia][0] + points[ia][1]*points[ia][1]
		distB := points[ib][0]*points[ib][0] + points[ib][1]*points[ib][1]
		return -utils.IntComparator(distA, distB)
	})
	for i, point := range points {
		dist := point[0]*point[0] + point[1]*point[1]
		if maxHeap.Size() < k {
			maxHeap.Push(i)
		} else {
			peekVal, ok := maxHeap.Peek()
			if !ok {
				panic("something went wrong")
			}
			if dist < points[peekVal.(int)][0]*points[peekVal.(int)][0]+points[peekVal.(int)][1]*points[peekVal.(int)][1] {
				maxHeap.Pop()
				maxHeap.Push(i)
			}
		}
	}

	res := make([][]int, k)
	for _, elem := range maxHeap.Values() {
		res = append(res, []int{points[elem.(int)][0], points[elem.(int)][1]})
	}
	return res
}
