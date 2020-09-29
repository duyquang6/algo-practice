package leetcode

func wordBreak(s string, wordDict []string) bool {

	wordmap := map[string]bool{}

	for _, v := range wordDict {
		wordmap[v] = true
	}

	se := 1
	dp := []int{}
	for se <= len(s) {
		for _, v := range dp {
			if wordmap[s[v:se]] {
				dp = append(dp, se)
				break
			}
		}

		if wordmap[s[0:se]] {
			dp = append(dp, se)
		}
		se++
	}

	return dp[len(dp)-1] == len(s)
}
