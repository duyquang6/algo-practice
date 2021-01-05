// https://leetcode.com/problems/consecutive-characters/
package leetcode

func maxPower(s string) int {
	res := 0
	counter := map[rune]int{}
	for i, c := range s {
		counter[c]++
		if counter[c] > res {
			res = counter[c]
		}
		if i < len(s) && s[i] != s[i+1] {
			counter[c] = 0
		}
	}
	return res
}
