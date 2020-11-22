// https://leetcode.com/explore/challenge/card/november-leetcoding-challenge/567/week-4-november-22nd-november-28th/3540/
package leetcode

import (
	"strings"
)

func uniqueMorseRepresentations(words []string) int {
	morseCodes := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	morseMap := map[string]struct{}{}
	for _, word := range words {
		sb := strings.Builder{}
		for _, ch := range word {
			sb.WriteString(morseCodes[ch-'a'])
		}
		morseMap[sb.String()] = struct{}{}
	}
	return len(morseMap)
}
