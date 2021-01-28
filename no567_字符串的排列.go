package main

import (
	"fmt"
)

//给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
//换句话说，第一个字符串的排列之一是第二个字符串的子串。
//示例1:
//输入: s1 = "ab" s2 = "eidbaooo"
//输出: True
//解释: s2 包含 s1 的排列之一 ("ba").
// 
//示例2:
//输入: s1= "ab" s2 = "eidboaoo"
//输出: False

func main() {
	a := "aba"
	b := "eidboaoo"
	fmt.Println(checkInclusion(a, b))
}

func checkInclusion(s1 string, s2 string) bool {
	need, window := make(map[byte]int), make(map[byte]int)
	// 需要的
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}

	l, r := 0, 0
	// 记录满足的进度，否则需要比较两个hash表
	vaild := 0

	// 右指针开始滑动
	for r < len(s2) {
		c := s2[r]
		r++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				vaild++
			}
		}

		// 当左右间隔一个窗口大小， 左指针可以收缩
		for r-l >= len(s1) {
			if vaild == len(need) {
				return true
			}

			d := s2[l]
			l++
			//左指针指向的是否要剔除
			if need[d] > 0 {
				if window[d] == need[d] {
					vaild--
				}
				window[d] --
			}
		}
	}
	return false
}
