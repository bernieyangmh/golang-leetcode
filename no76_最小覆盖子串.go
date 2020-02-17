// 给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。
// 示例：
// 输入: S = "ADOBECODEBANC", T = "ABC"
// 输出: "BANC"
// 说明：
// 如果 S 中不存这样的子串，则返回空字符串 ""。
// 如果 S 中存在这样的子串，我们保证它是唯一的答案。
package main

// 双指针, 右边的指针一直向右知道 window包含属于t的每个字符的次数都符合要求
// 符合要求之后,右指针先停,左指针向右收缩,直到不符合要求
func minWindow(s string, t string) string {
	// 长度取最大数
	minLen := 1 << 32
	start := 0
	left := 0
	right := 0
	match := 0

	// 注意是byte, go里索引某个字符串是byte, range 每个字符串是rune
	// 当前窗口内的字符及包含次数
	window := make(map[byte]int)
	// 所需单词的字符及次数
	needs := make(map[byte]int)

	// 存储t里每个字符出现的次数
	for i := 0; i < len(t); i++ {
		if _, ok := needs[t[i]]; ok {
			needs[t[i]]++
		} else {
			needs[t[i]] = 1
		}
	}

	// 右指针一直走到尽头
	for right < len(s) {

		// 如果右指针指向的字符x是t里的, 窗口里x出现的次数+1, 如果x的次数和t出现的次数相同, 这个字符就符合了
		if _, ok := needs[s[right]]; ok {
			window[s[right]]++
			// 只需相同就说明该字符x符合要求,再多只增加窗口里x的次数,match不再增加
			if window[s[right]] == needs[s[right]] {
				match++
			}
		}
		// 右指针向前, 因为距离=j-i+1
		right++

		// 如果所有字符都符合次数要求,收缩
		for match == len(needs) {
			if right-left < minLen {
				start = left
				minLen = right - left
			}
			// 如果左指针移动中遇到了t包含的字符,窗口里包含该x字符的次数减少,并看看是不是多余的
			if cnt, ok := needs[s[left]]; ok && cnt > 0 {
				window[s[left]]--
				// 如果不多余,当前不匹配
				if window[s[left]] < needs[s[left]] {
					match--
				}
			}
			left++
		}

	}
	if minLen == 1<<32 {
		return ""
	}
	return s[start : start+minLen]
}
