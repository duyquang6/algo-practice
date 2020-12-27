// https://leetcode.com/problems/count-of-matches-in-tournament/
package leetcode

func numberOfMatches(n int) int {
	res := 0
	for n > 1 {
		if n&1 == 0 {
			n /= 2
			res += n
			continue
		}
		n = (n-1)/2 + 1
		res += n - 1
	}
	return res
}
