package main

import (
	"fmt"
)

//给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
//在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。
//注意:
//假设字符串的长度不会超过 1010。
//示例 1:
//输入:
//"abccccdd"
//输出:
//7
//解释:
//我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。

func main() {
	a := "dccaccd"
	fmt.Println(longestPalindrome(a))

}

func longestPalindrome(s string) int {
	cMap := make(map[rune]int)
	for _, c := range s {
		if _, ok := cMap[c]; ok {
			cMap[c]++
		} else {
			cMap[c] = 1
		}
	}
	//如果有奇数个，中间可以是单个字符
	midSingle := false
	length := 0
	for _, cnt := range cMap {
		// 去掉多余的1
		length += (cnt / 2) * 2
		if cnt%2 == 1 {
			midSingle = true
		}
	}
	if midSingle {
		length++
	}
	return length
}
