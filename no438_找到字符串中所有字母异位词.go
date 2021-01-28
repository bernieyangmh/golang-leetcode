package main

import (
	"fmt"
)

//给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
//字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。
//说明：
//字母异位词指字母相同，但排列不同的字符串。
//不考虑答案输出的顺序。
//示例 1:
//输入:
//s: "cbaebabacd" p: "abc"
//输出:
//[0, 6]
//解释:
//起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
//起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
// 示例 2:
//输入:
//s: "abab" p: "ab"
//输出:
//[0, 1, 2]
//解释:
//起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
//起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
//起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。

func main() {
	s := "abaacbabc"
	p := "abbaac"
	fmt.Println(findAnagrams2(s, p))

}

// 滑动窗口，两个while循环
func findAnagrams(s string, p string) []int {
	matchNum := 0
	wordMap := make(map[string]int)
	curMap := make(map[string]int)
	res := []int{}
	left, right := 0, 0

	for _, s := range p {
		if _, ok := wordMap[string(s)]; ok {
			wordMap[string(s)]++
			curMap[string(s)] = 0
		} else {
			wordMap[string(s)] = 1
		}
	}
	for right < len(s) {
		c1 := string(s[right])
		if _, ok := wordMap[c1]; ok {
			curMap[c1]++
			if wordMap[c1] == curMap[c1] {
				matchNum++
			}
		}
		right++
		for matchNum == len(wordMap) {
			if right-left == len(p) {
				res = append(res, left)
			}
			c2 := string(s[left])
			if _, ok := wordMap[c2]; ok {
				curMap[c2]--
				if curMap[c2] < wordMap[c2] {
					matchNum--
				}
			}
			left++
		}
	}
	return res

}

//别人写的
func findAnagrams2(s string, p string) []int {
	var (
		res                = make([]int, 0)
		match, left, right int
		win                [26]int
		need               [26]int
		needs              int
	)
	for _, i := range p {
		if need[i-'a'] == 0 {
			needs++
		}
		need[i-'a']++
	}
	fmt.Println(need)
	for right < len(s) {
		c1 := s[right] - 'a'
		if need[c1] > 0 {
			win[c1]++
			if win[c1] == need[c1] {
				match++
			}
		}
		right++
		for match == needs {
			if right-left == len(p) {
				res = append(res, left)
			}
			c2 := s[left] - 'a'
			if need[c2] > 0 {
				win[c2]--
				if win[c2] < need[c2] {
					match--
				}
			}
			left++
		}
	}
	return res
}

func findAnagrams3(s string, p string) []int {
	need, window := make(map[byte]int), make(map[byte]int)
	// 需要的
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	//左指针，右指针，进度
	l, r, valid := 0, 0, 0
	res := []int{}

	//右指针滑动
	for r < len(s) {
		c := s[r]
		r++

		//c是我们需要的， 放进窗口里， 如果窗口数量达到，进度+1
		if need[c] > 0 {
			window[c] ++
			if window[c] == need[c] {
				valid++
			}
		}

		// 超过窗口大小，左指针该收缩了
		for r-l >= len(p) {
			// 先判断是否满足
			if valid == len(need) {
				res = append(res, l)
			}

			// 左指针收缩，窗口数量-1， 进度-1
			d := s[l]
			l++
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d] --
			}

		}
	}
	return res
}
