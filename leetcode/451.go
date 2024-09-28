package leetcode

import (
	"strings"

	"github.com/emirpasic/gods/trees/binaryheap"
	"github.com/emirpasic/gods/utils"
)

func frequencySort(s string) string {
	freq := make(map[rune]int, 62)
	for _, c := range s {
		freq[c] += 1
	}

	maxHeap := binaryheap.NewWith(func(a, b interface{}) int {
		ia, ib := a.(rune), b.(rune)
		return utils.IntComparator(freq[ib], freq[ia])
	})

	for k := range freq {
		maxHeap.Push(k)
	}

	res := strings.Builder{}
	for maxHeap.Size() > 0 {
		k, _ := maxHeap.Pop()
		kr := k.(rune)
		for i := 0; i < freq[kr]; i++ {
			res.WriteRune(k.(rune))
		}
	}

	return res.String()
}
