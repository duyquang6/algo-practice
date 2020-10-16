// https://leetcode.com/problems/search-a-2d-matrix/
package leetcode

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	fmt.Println(m, n)
	N := m * n
	ss, se := 0, N-1
	for ss <= se {
		mid := (ss + se) / 2
		im, jm := mid/n, mid%n
		if matrix[im][jm] == target {
			return true
		}
		if matrix[im][jm] < target {
			ss = mid + 1
		} else {
			se = mid - 1
		}
	}
	return false
}
