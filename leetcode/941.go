// https://leetcode.com/problems/valid-mountain-array/
package leetcode

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	if A[0] > A[1] {
		return false
	}
	isDecrease := false
	for i := 1; i < len(A); i++ {
		if A[i] == A[i-1] {
			return false
		}
		if A[i-1] < A[i] {
			if isDecrease {
				return false
			}
		} else {
			isDecrease = true
		}
	}
	return isDecrease
}
