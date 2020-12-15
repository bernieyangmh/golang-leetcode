package main

import (
	"fmt"
)

//给定两个字符串 s 和 t，判断它们是否是同构的。
//如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。
//所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。
//示例 1:
//输入: s = "egg", t = "add"
//输出: true
//示例 2:
//输入: s = "foo", t = "bar"
//输出: false
//示例 3:
//输入: s = "paper", t = "title"
//输出: true

func main() {
	a := "ab"
	b := "aa"
	fmt.Println(isIsomorphic(a, b))

}

// 两个map保存相互之间的映射
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sBit := []rune{}
	c2toc1 := make(map[rune]rune)
	c1toc2 := make(map[rune]rune)

	for _, c1 := range s {
		sBit = append(sBit, c1)
	}
	for index, c2 := range t {
		prevC1, ok := c2toc1[c2]
		// fmt.Println(index, c2, c1, ok, c2toc1, sBit[index])
		if ok {
			if sBit[index] != prevC1 {
				return false
			}
		} else {
			c2toc1[c2] = sBit[index]
		}
		prevC2, ok := c1toc2[sBit[index]]
		// fmt.Println(index, c2, c1, ok, c2toc1, sBit[index])
		if ok {
			if prevC2 != c2 {
				return false
			}
		} else {
			c1toc2[sBit[index]] = c2
		}
	}
	return true
}

// 一个mapping， 调换计算两次
func isIsomorphic2(s string, t string) bool {
	return helper(s, t) && helper(t, s)
}

func helper(s string, t string) bool {
	length := len(s)
	mapping := make(map[byte]byte)

	for i := 0; i < length ; i++{
		c1 := s[i]
		c2 := t[i]
		if old, ok := mapping[c1];ok{
			if old != c2 {
				return false
			}
		} else {
			mapping[c1] = c2
		}
	}
	return true
}

















