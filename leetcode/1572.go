// https://leetcode.com/problems/matrix-diagonal-sum/
package leetcode

func diagonalSum(mat [][]int) int {
	ans := 0
	for i := 0; i < len(mat); i++ {
		ans += mat[i][i] + mat[i][len(mat)-i-1]
	}
	if len(mat)&1 == 1 {
		ans -= mat[len(mat)/2][len(mat)/2]
	}
	return ans
}
