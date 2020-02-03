// 给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。返回所有这些可能的句子。
// 说明：
// 分隔时可以重复使用字典中的单词。
// 你可以假设字典中没有重复的单词。
// 示例 1：
// 输入:
// s = "catsanddog"
// wordDict = ["cat", "cats", "and", "sand", "dog"]
// 输出:
// [
//   "cats and dog",
//   "cat sand dog"
// ]
// 示例 2：
// 输入:
// s = "pineapplepenapple"
// wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
// 输出:
// [
//   "pine apple pen apple",
//   "pineapple pen apple",
//   "pine applepen apple"
// ]
// 解释: 注意你可以重复使用字典中的单词。
// 示例 3：
// 输入:
// s = "catsandog"
// wordDict = ["cats", "dog", "sand", "and", "cat"]
// 输出:
// []

package main

import "fmt"

func main() {
	for _, s := range wordBreak("pineapplepenapple", []string{"pen", "applepen", "apple", "pine", "pineapple"}) {
		fmt.Println(s)
	}
}

// https://leetcode-cn.com/problems/word-break-ii/solution/dong-tai-gui-hua-hui-su-qiu-jie-ju-ti-zhi-python-d/
func wordBreak(s string, wordDict []string) []string {
	dp := []bool{}
	for i := 0; i <= len(s); i++ {
		dp = append(dp, false)
	}
	hashMap := make(map[string]bool)
	for _, w := range wordDict {
		hashMap[w] = true
	}
	dp[0] = true
	for r := 1; r < len(s)+1; r++ {
		for l := 0; l < r; l++ {
			if dp[l] {
				if _, ok := hashMap[s[l:r]]; ok {
					dp[r] = true
					break
				}
			}
		}
	}
	// 生成dp参考139
	var res []string
	if dp[len(s)] {
		res = dfs(s, len(s), hashMap, res, []string{}, dp)
		return res
	}
	return res

}

func dfs(s string, end int, hashMap map[string]bool, res, path []string, dp []bool) []string {
	// 一种路径走到底,加进结果里
	if end == 0 {
		tmpStr := ""
		for _, value := range path {
			tmpStr += value
			tmpStr += " "
		}
		res = append(res, tmpStr[:len(tmpStr)-1])
		return res
	}
	// e   i
	// apple
	//     t

	// 以每个单词作为根结点,组成二叉树
	// a b c d e f g

	for i := 0; i < end; i++ {
		if dp[i] {
			if _, ok := hashMap[s[i:end]]; ok {
				// path 不变,回溯
				// path = [pine apple]
				// dfs...
				// path back -> [pine], i继续前进到下一个单词
				// path = [pine applepen]
				tmp := []string{s[i:end]}
				tmp = append(tmp, path...)
				res = dfs(s, i, hashMap, res, tmp, dp)
			}
		}
	}
	return res
}
