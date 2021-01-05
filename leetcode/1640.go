package leetcode

func canFormArray(arr []int, pieces [][]int) bool {
	i := 0

	for i < len(arr) {
		cur := i
		for _, piece := range pieces {
			cond := true
			for j := 0; j < len(piece); j++ {
				if arr[i+j] != piece[j] {
					cond = false
					break
				}
			}

			if !cond {
				continue
			}
			i += len(piece)
			break
		}
		if i == cur {
			return false
		}
	}

	return i == len(arr)
}
