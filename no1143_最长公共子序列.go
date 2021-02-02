package main

func longestCommonSubsequence(text1 string, text2 string) int {
	// write code here
	m := len(text1)
	n := len(text2)

	dp := make([][]int, m+1)

	for j := 0; j < m+1; j++ {
		dp[j] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
