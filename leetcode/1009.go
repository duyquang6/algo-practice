// https://leetcode.com/problems/complement-of-base-10-integer/
package leetcode

func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}
	i, cur := 0, 1
	ans := 0
	for N > 0 {
		if N&1 == 0 {
			ans += cur
		}
		N >>= 1
		i++
		cur *= 2
	}

	return ans
}
