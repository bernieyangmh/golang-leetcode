// 给定一个字符串s，找到其中最长的回文子序列。可以假设s的最大长度为1000。
// 示例 1:
// 输入:
// "bbbab"
// 输出:4
// 一个可能的最长回文子序列为 "bbbb"。
// 示例 2:
// 输入:
// "cbbd"
// 输出:2
package main

import (
	"fmt"
)

func main() {
	s := `zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz`
	fmt.Println(longestPalindromeSubseq2(s))

}

// f[i][j] 表示s的第i个字符到第j个字符组成的子串中，最长的回文序列长度是多少
func longestPalindromeSubseq(s string) int {
	var n int
	size := len(s)
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		tmp := make([]int, size)
		// 每个单个字符串都是长度1的回文
		tmp[i] = 1
		dp[i] = tmp
	}
	fmt.Println(dp)
	// 指针从后向前, 遍历指针后面的字符串
	for i := size - 1; i >= 0; i-- {
		for j := i + 1; j < size; j++ {
			n++
			// 如果回文,+2
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				// 不是回文,取左缩小或右缩小的最大值
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	fmt.Println(dp)
	fmt.Println(n)
	return dp[0][size-1]
}

// 从前向后的超时...
func longestPalindromeSubseq2(s string) int {
	var n int
	size := len(s)
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		tmp := make([]int, size)
		tmp[i] = 1
		dp[i] = tmp
	}
	fmt.Println(dp)
	for j := 0; j < size; j++ {
		for i := j - 1; i >= 0; i-- {
			n++
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				// fmt.Println(dp, i, j)
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	fmt.Println(dp)
	fmt.Println(n)
	return dp[0][size-1]
}

// "bbbab"
// [[1 0 0 0 0] [0 1 0 0 0] [0 0 1 0 0] [0 0 0 1 0] [0 0 0 0 1]]
// [[1 2 3 3 2] [0 1 2 2 2] [0 0 1 1 3] [0 0 0 1 1] [0 0 0 0 1]]

// [[1 0 0 0 0] [0 1 0 0 0] [0 0 1 0 0] [0 0 0 1 0] [0 0 0 0 1]]
// [[1 2 3 3 4] [0 1 2 2 3] [0 0 1 1 3] [0 0 0 1 1] [0 0 0 0 1]]

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
