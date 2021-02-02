package main

//最长公共子串

func LCS(str1 string, str2 string) int {
	maxValue := 0

	m, n := len(str1), len(str2)
	if m == 0 || n == 0 {
		return 0
	}
	dp := make([][]int, m+1)

	for j := 0; j < m+1; j++ {
		dp[j] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				// 如果相同，取最大值
				if dp[i][j] > maxValue {
					maxValue = dp[i][j]
				}
				// 如果不相同，重置
			} else {
				dp[i][j] = 0
			}
		}
	}
	return maxValue
}
