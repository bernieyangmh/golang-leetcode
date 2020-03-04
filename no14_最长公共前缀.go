package main

//编写一个函数来查找字符串数组中的最长公共前缀。
//如果不存在公共前缀，返回空字符串 ""。
//示例 1:
//输入: ["flower","flow","flight"]
//输出: "fl"
//示例 2:
//输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//说明:
//所有输入只包含小写字母 a-z 

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minLength := 1<<31 - 1
	for i := 0; i < len(strs); i++ {
		minLength = min(minLength, len(strs[i]))
	}
	prefixStr := ""
	fmt.Println(minLength)
	for i := 0; i < minLength; i++ {
		curL := strs[0][i]
		// fmt.Println(curL, i)
		for j := 0; j < len(strs); j++ {
			if curL != strs[j][i] {
				return prefixStr
			}
		}
		prefixStr += string(curL)
	}
	return prefixStr
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != c {
				return strs[0][0:i]
			}
		}
	}
	return strs[0]

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
