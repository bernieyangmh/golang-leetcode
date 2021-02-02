package main

import (
	"fmt"
)

//最长公共子序列

// 二维数组， 表示两个字符串
// 一开始第一行均是0， 表示空集
// 不同时取左边和上面 最大值， 相同时左上数值+1

// 返回长度
func LCS(str1 string, str2 string) int {
	// write code here
	m := len(str1)
	n := len(str2)

	dp := make([][]int, m+1)

	for j := 0; j < m+1; j++ {
		dp[j] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[m][n]
}

//返回字符串
func LCS2(str1 string, str2 string) string {
	// write code here
	m := len(str1)
	n := len(str2)

	if m == 0 || n == 0 {
		return "-1"
	}
	dp := make([][]int, m+1)

	for j := 0; j < m+1; j++ {
		dp[j] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	res := []string{}
	res = trackBack(m, n, "", str1, str2, dp)
	return res[0]
}

func trackBack(i, j int, lcsS, str1, str2 string, dp [][]int) (res []string) {
	for i > 0 && j > 0 {
		if str1[i-1] == str2[j-1] {
			lcsS += string(str1[i-1])
			i--
			j--
		} else {
			if dp[i-1][j] > dp[i][j-1] {
				i--
			} else if dp[i-1][j] < dp[i][j-1] {
				j--
			} else {
				res = trackBack(i-1, j, lcsS, str1, str2, dp)
				res = trackBack(i, j-1, lcsS, str1, str2, dp)
			}
		}
	}
	res = append(res, reverse(lcsS))
	return res
}

func reverse(str string) string {
	var result string
	strLen := len(str)
	for i := 0; i < strLen; i++ {
		result = result + fmt.Sprintf("%c", str[strLen-i-1])
	}
	return result
}

func LCSLength(x, y string) int {
	if len(x) == 0 || len(y) == 0 {
		return 0
	}

	if x[len(x)-1] == y[len(y)-1] {
		return 1 + LCSLength(x[0:len(x)-1], y[0:len(y)-1])
	} else {
		return max(LCSLength(x, y[0:len(y)-1]), LCSLength(x[0:len(x)-1], y))
	}
}

func LCS1(a, b string) (int, string) {
	arunes := []rune(a)
	brunes := []rune(b)
	aLen := len(arunes)
	bLen := len(brunes)
	lengths := make([][]int, aLen+1)
	for i := 0; i <= aLen; i++ {
		lengths[i] = make([]int, bLen+1)
	}

	// row 0 and column 0 are initialized to 0 already
	for i := 0; i < aLen; i++ {
		for j := 0; j < bLen; j++ {
			if arunes[i] == brunes[j] {
				lengths[i+1][j+1] = lengths[i][j] + 1
			} else if lengths[i+1][j] > lengths[i][j+1] {
				lengths[i+1][j+1] = lengths[i+1][j]
			} else {
				lengths[i+1][j+1] = lengths[i][j+1]
			}
		}
	}

	// read the substring out from the matrix
	s := make([]rune, 0, lengths[aLen][bLen])
	for x, y := aLen, bLen; x != 0 && y != 0; {
		if lengths[x][y] == lengths[x-1][y] {
			x--
		} else if lengths[x][y] == lengths[x][y-1] {
			y--
		} else {
			s = append(s, arunes[x-1])
			x--
			y--
		}
	}
	// reverse string
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return len(s), string(s)
}

func LCS3(str1 string, str2 string) string {
	m, n := len(str1), len(str2)
	dp := make([][]int, m+1)

	for j := 0; j < m+1; j++ {
		dp[j] = make([]int, n+1)
	}

	max, index := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if str1[i] == str2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
				if max < dp[i+1][j+1] {
					max = dp[i+1][j+1]
					index = i + 1
				}
			}
		}
	}
	return str1[index-max : index]

}
