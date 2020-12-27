// https://leetcode.com/problems/maximum-repeating-substring/
package leetcode

import (
	"strings"
)

func maxRepeating(sequence string, word string) int {
	best := 0
	temp := word
	for true {
		if !strings.Contains(sequence, temp) {
			break
		}
		best++
		temp += word
	}
	return best
}
