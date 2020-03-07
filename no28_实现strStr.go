package main

import (
	"fmt"
)

//实现 strStr() 函数。
//给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
//示例 1:
//输入: haystack = "hello", needle = "ll"
//输出: 2
//示例 2:
//输入: haystack = "aaaaa", needle = "bba"
//输出: -1
//说明:
//当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
//对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。

func main() {
	a := "mississippi"
	b := "issip"
	fmt.Println(strStr(a, b))

}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 {
		return -1
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		has := true
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				has = false
				break
			}
		}
		if has {
			return i
		}
	}
	return -1
}

func strStr2(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 {
		return -1
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if j == len(needle)-1 && haystack[i+j] == needle[j] {
				return i
			}
		}
	}
	return -1
}

//kmp
func strStr3(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 {
		return -1
	}

	x := 0
	dp := [][]int{}
	m := len(needle)

	for i := 0; i < m; i++ {
		tmp := make([]int, 256)
		dp = append(dp, tmp)
	}
	dp[0][needle[0]] = 1
	for i := 1; i < m; i++ {
		for j := 0; j < 256; j++ {
			dp[i][j] = dp[x][j]
		}
		dp[i][needle[i]] = i + 1
		x = dp[x][needle[i]]
	}

	s := 0
	for i := 0; i < len(haystack); i++ {
		s = dp[s][haystack[i]]
		if s == m {
			return i - m + 1
		}
	}
	return -1
}
