// https://leetcode.com/problems/buddy-strings/
package leetcode

import "reflect"

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) || len(A) == 0 {
		return false
	}
	countDiff := 0
	idx := []int{}
	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			idx = append(idx, i)
			countDiff++
		}
	}

	if countDiff == 0 {
		m := map[rune]int{}
		for _, v := range A {
			m[v]++
			if m[v] > 1 {
				return true
			}
		}
		return false
	}

	if countDiff == 2 {
		r := []rune(A)
		r[idx[0]], r[idx[1]] = r[idx[1]], r[idx[0]]
		return reflect.DeepEqual(string(r), B)
	}

	return false
}
