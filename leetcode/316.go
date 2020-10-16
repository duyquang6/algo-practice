package leetcode

func removeDuplicateLetters(s string) string {
	counter := map[rune]uint{}
	for _, v := range s {
		counter[v]++
	}
	for k, v := range counter {
		if v < 2 {
			continue
		}
		prev := -1
		for i, c := range s {
			if k == c {
				if prev == -1 {
					prev = i
					continue
				}

			}

		}
	}
	return ""
}
