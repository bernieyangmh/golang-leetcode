package main

// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
// 示例 1:
// 输入: "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:

// 输入: "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:

// 输入: "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

import "fmt"

func main() {
	println(lengthOfLongestSubstring("aa"))

}

func lengthOfLongestSubstring(s string) int {
	var res, index int
	uniqueMap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		// 保持最新的索引,当有重复的数字的时候,更新开头的指针从此开始
		if _, ok := uniqueMap[s[i]]; ok {
			index = max(uniqueMap[s[i]], index)
		}

		// [a,b,c] 共2-0+1 个数字
		// 随着i的增加,没有遇到重复数字,index不变,滑动窗口不断增大
		res = max(res, i-index+1)
		// 放入数字的最大索引值
		uniqueMap[s[i]] = i + 1
		fmt.Println(i, index, res, uniqueMap, s[:i]+string(s[i]))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// public class Solution {
//     public int lengthOfLongestSubstring(String s) {
//         int n = s.length(), ans = 0;
//         Map<Character, Integer> map = new HashMap<>(); // current index of character
//         // try to extend the range [i, j]
//         for (int j = 0, i = 0; j < n; j++) {
//             if (map.containsKey(s.charAt(j))) {
//                 i = Math.max(map.get(s.charAt(j)), i);
//             }
//             ans = Math.max(ans, j - i + 1);
//             map.put(s.charAt(j), j + 1);
//         }
//         return ans;
//     }
// }
