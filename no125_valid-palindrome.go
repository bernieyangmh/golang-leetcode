package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	a := "123321"
	println(isPalindrome(a))
}

// 注意string和rune的转换
// 判断是否是字母或数字的前置关系

func isPalindrome(s string) bool {
	r := []rune(s)
	var start, end int
	start = 0
	end = len(r) - 1
	for start < end {
		fmt.Println(strings.ToLower(string(r[start])), strings.ToLower(string(r[end])), start, end)
		if !unicode.IsLetter(r[start]) && !unicode.IsDigit(r[start]) {
			start++
			continue
		}
		if !unicode.IsLetter(r[end]) && !unicode.IsDigit(r[end]) {
			end++
			continue
		}
		if strings.ToLower(string(r[start])) != strings.ToLower(string(r[end])) {
			return false
		} else {
			start++
			end--
		}
	}
	return true
}
