package problem

import (
	"math"
	"slices"
)

func minElement(nums []int) int {
	_min := math.MaxInt
	for _, num := range nums {
		tmp := 0
		for num > 0 {
			tmp += num % 10
			num /= 10
		}
		_min = min(tmp, _min)
	}
	return _min
}

func maximumTotalSum(maximumHeight []int) int64 {
	slices.Sort(maximumHeight)
	lastMaxHeight := maximumHeight[len(maximumHeight)-1]
	_sum := int64(maximumHeight[len(maximumHeight)-1])
	for i := len(maximumHeight) - 2; i >= 0; i-- {
		if maximumHeight[i] >= lastMaxHeight {
			lastMaxHeight -= 1
		} else {
			lastMaxHeight = maximumHeight[i]
		}
		_sum += int64(lastMaxHeight)
	}

	return _sum
}

func validSequence(word1 string, word2 string) []int {
	paths := []int{}
	if !recur(word1, word2, 0, 0, &paths, false) {
		return nil
	}
	return paths
}

func recur(word1 string, word2 string, w1i, w2i int, paths *[]int, hasReplace bool) bool {
	if w2i == len(word2) {
		return true
	}
	if len(word1)-w1i < len(word2)-w2i {
		return false
	}

	pathLen := len(*paths)
	if word1[w1i] == word2[w2i] {
		// pick
		*paths = append(*paths, w1i)
		return recur(word1, word2, w1i+1, w2i+1, paths, hasReplace)
	}

	// try to choose
	if !hasReplace {
		*paths = append(*paths, w1i)
		if recur(word1, word2, w1i+1, w2i+1, paths, true) {
			return true
		}
		*paths = (*paths)[:pathLen]
	}

	// skip to match word
	for w1i < len(word1) && word1[w1i] != word2[w2i] {
		w1i++
	}
	return recur(word1, word2, w1i, w2i, paths, hasReplace)
}
