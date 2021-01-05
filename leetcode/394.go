// https://leetcode.com/explore/challenge/card/november-leetcoding-challenge/566/week-3-november-15th-november-21st/3536/
package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func decodeString(s string) string {
	res := strings.Builder{}
	i := 0
	for i < len(s) {
		ch := s[i]
		if ch >= '0' && ch <= '9' {
			j := i + 1
			for j < len(s) && isDigit(rune(s[j])) {
				j++
			}
			k, _ := strconv.Atoi(s[i:j])
			fmt.Println(k)
		} else if ch == ']' {

		} else {

		}
	}
	return res.String()
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}
