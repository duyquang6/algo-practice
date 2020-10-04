// https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/
package leetcode

import (
	"container/list"
	"strings"
)

func MinRemoveToMakeValid(s string) string {
	stack := list.New()
	buffer := strings.Builder{}
	temp := []int{}
	for i, v := range s {
		if v == ')' {
			if stack.Len() == 0 {
				temp = append(temp, i)
				// Delete this char
			} else {
				// pop stack
				stack.Remove(stack.Back())
			}
		} else if v == '(' {
			stack.PushBack(i)
		}
	}

	p := 0
	q := stack.Front()
	for i := range s {
		if len(temp) > 0 && p < len(temp) && temp[p] == i {
			p++
			continue
		}

		if len(temp) == p && q != nil && q.Value.(int) == i {
			q = q.Next()
			continue
		}

		buffer.WriteByte(s[i])
	}

	return buffer.String()
}
