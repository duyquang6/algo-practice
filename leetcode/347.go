package leetcode

import (
	"github.com/emirpasic/gods/trees/binaryheap"
	"github.com/emirpasic/gods/utils"
)

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int, len(nums))
	for _, n := range nums {
		freq[n] += 1
	}
	minHeap := binaryheap.NewWith(func(a, b interface{}) int {
		ia := a.(int)
		ib := b.(int)
		return utils.IntComparator(freq[ia], freq[ib])
	})
	for num, fq := range freq {
		if minHeap.Size() < k {
			minHeap.Push(num)
		} else {
			peekValInterface, _ := minHeap.Peek()
			peekVal := peekValInterface.(int)
			if fq > freq[peekVal] {
				minHeap.Pop()
				minHeap.Push(num)
			}
		}
	}

	res := make([]int, 0, k)

	for _, peekValInterface := range minHeap.Values() {
		peekVal := peekValInterface.(int)
		res = append(res, peekVal)
	}

	return res
}
