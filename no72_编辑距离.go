// 给定两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数 。
// 你可以对一个单词进行如下三种操作：
// 插入一个字符
// 删除一个字符
// 替换一个字符
// 示例 1:
// 输入: word1 = "horse", word2 = "ros"
// 输出: 3
// 解释:
// horse -> rorse (将 'h' 替换为 'r')
// rorse -> rose (删除 'r')
// rose -> ros (删除 'e')
// 示例 2:
// 输入: word1 = "intention", word2 = "execution"
// 输出: 5
// 解释:
// intention -> inention (删除 't')
// inention -> enention (将 'i' 替换为 'e')
// enention -> exention (将 'n' 替换为 'x')
// exention -> exection (将 'n' 替换为 'c')
// exection -> execution (插入 'u')
package main

// 动态规划
// dp[i][j] 存储 s1[0..i] 和 s2[0..j] 的最小编辑距离
// if s1[i] == s2[j]:
//     跳过,不需要步骤,仍是之前的操作数
//     i, j 同时向前移动
// else:
//     三选一：以最小操作数为准, 不管选哪个肯定执行了一次操作, +1
//         插入		dp(i, j - 1)
//         删除		dp(i - 1, j)
//         替换		dp(i - 1, j - 1)

func minDistance(word1 string, word2 string) int {
	n1 := len(word1)
	n2 := len(word2)

	// 多了一个,便于计算
	dp := make([][]int, n1+1)
	for i := 0; i < len(dp); i++ {
		tmp := make([]int, n2+1)
		dp[i] = tmp
	}

	// 初始化列 从空字符串到组成w1最终需要多少操作数
	for j := 1; j <= n2; j++ {
		dp[0][j] = dp[0][j-1] + 1
	}

	// 初始化行 从空字符串到组成w2最终需要多少操作数
	for i := 1; i <= n1; i++ {
		dp[i][0] = dp[i-1][0] + 1
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j-1], dp[i][j-1]), dp[i-1][j]) + 1
			}
		}
	}
	return dp[n1][n2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
