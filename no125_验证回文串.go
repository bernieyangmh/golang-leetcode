package main

import (
	"fmt"
	"strings"
	"unicode"
)

//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//说明：本题中，我们将空字符串定义为有效的回文串。
//示例 1:
//输入: "A man, a plan, a canal: Panama"
//输出: true
//示例 2:
//输入: "race a car"
//输出: false

func main() {
	a := "race a car"
	fmt.Println(isPalindrome(a))

	fmt.Println(fmt.Sprintf("%b", -4))

}

// 注意string和rune的转换
// 判断是否是字母或数字的前置关系
func isPalindrome(s string) bool {
	r := []rune(s)
	var start, end int
	start = 0
	end = len(r) - 1
	for start < end {
		// fmt.Println(strings.ToLower(string(r[start])), strings.ToLower(string(r[end])), start, end)
		if !unicode.IsLetter(r[start]) && !unicode.IsDigit(r[start]) {
			start++
			continue
		}
		if !unicode.IsLetter(r[end]) && !unicode.IsDigit(r[end]) {
			end--
			continue
		}
		if strings.ToLower(string(r[start])) != strings.ToLower(string(r[end])) {
			fmt.Println(start, end, string(r[start]), string(r[end]))
			return false
		} else {
			start++
			end--
		}
	}
	return true
}
