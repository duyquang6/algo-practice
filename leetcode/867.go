// https://leetcode.com/problems/transpose-matrix/
package leetcode

func transpose(A [][]int) [][]int {
	m, n := len(A), len(A[0])
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, m)
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			res[j][i] = A[i][j]
		}
	}

	return res
}
