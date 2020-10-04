// https://leetcode.com/problems/sort-array-by-parity-ii/
package leetcode

func sortArrayByParityII(A []int) []int {
	dist := make([]int, 1001)
	for _, v := range A {
		dist[v]++
	}
	evenIdx := 0
	oddIdx := 1
	for idx, v := range dist {
		if v == 0 {
			continue
		}
		if idx&1 == 0 {
			for i := 0; i < v; i++ {
				A[evenIdx] = idx
				evenIdx += 2
			}
		} else {
			for i := 0; i < v; i++ {
				A[oddIdx] = idx
				oddIdx += 2
			}
		}
	}
	return A
}
