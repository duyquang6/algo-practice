package leetcode

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	out := merge([][]int{
		// [[1,3],[2,6],[8,10],[15,18]]
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	})
	fmt.Println(out)
}
