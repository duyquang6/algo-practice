package leetcode

func champagneTower(poured int, query_row int, query_glass int) float64 {
	dp := [][]float64{
		{float64(poured)},
	}
	i := 0
	isFinalRow := false
	for !isFinalRow && i < 100 {
		isFinalRow = true
		dp = append(dp, make([]float64, i+2))
		for j := 0; j <= i; j++ {
			if dp[i][j] > 1 {
				halfRemain := (dp[i][j] - 1) / 2
				dp[i][j] = 1
				dp[i+1][j] += halfRemain
				dp[i+1][j+1] += halfRemain
				isFinalRow = false
			}
		}
		if i == query_row {
			return dp[query_row][query_glass]
		}
		i += 1
	}
	if query_row < len(dp) {
		return dp[query_row][query_glass]
	}
	return 0
}
