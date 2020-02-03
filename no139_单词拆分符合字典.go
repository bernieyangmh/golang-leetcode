// 给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
// 说明：
// 拆分时可以重复使用字典中的单词。
// 你可以假设字典中没有重复的单词。
// 示例 1：
// 输入: s = "leetcode", wordDict = ["leet", "code"]
// 输出: true
// 解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
// 示例 2：
// 输入: s = "applepenapple", wordDict = ["apple", "pen"]
// 输出: true
// 解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
//      注意你可以重复使用字典中的单词。
// 示例 3：
// 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
// 输出: false

package main

import "fmt"

func main() {
	fmt.Println(wordBreak3("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	fmt.Println(wordBreak3("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak2("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"aaa", "aaaa", "aaaaaaaaaa"}))

}

// 暴力 最坏n^n 每次前缀都在字典里,超时
func wordBreak(s string, wordDict []string) bool {
	hashMap := make(map[string]bool)
	for _, w := range wordDict {
		hashMap[w] = true
	}
	return helper(s, hashMap, 0)
}

func helper(s string, hashMap map[string]bool, start int) bool {
	if start == len(s) {
		return true
	}
	for end := start + 1; end <= len(s); end++ {
		if _, ok := hashMap[s[start:end]]; ok {
			if helper(s, hashMap, end) {
				return true
			}
		}
	}
	return false
}

// 记录下了某个结点是否作为 单词起始 使用过,
func wordBreak2(s string, wordDict []string) bool {
	hashMap := make(map[string]bool)
	for _, w := range wordDict {
		hashMap[w] = true
	}
	used := []int{}
	for i := 0; i < len(s); i++ {
		used = append(used, 0)
	}
	return 1 == helper2(s, hashMap, 0, used)
}

func helper2(s string, hashMap map[string]bool, start int, used []int) int {
	// 能走到底,符合
	if start == len(s) {
		return 1
	}
	// 该位置已走过, 返回该位置的结果
	if used[start] != 0 {
		return used[start]
	}
	// 向后扩张单词 
	// apple | penapple
	//         |从这里再开始判断
	//  pplepenapple
	//    lepenapple
	//     epenapple
	for end := start + 1; end <= len(s); end++ {
		if _, ok := hashMap[s[start:end]]; ok {
			// 如果在字典里, 判断后面未走过的路径是否符合条件
			//             如果符合,记录当前位置可以
			if helper2(s, hashMap, end, used) == 1 {
				used[start] = 1
				return 1
			}
		}
	}
	// 从该位置开始向后走 没有符合的条件
	used[start] = 2
	return 2
}


// 动态规划
// AAA(OK) BBB(OK) -> AAABBB(OK)

func wordBreak3(s string, wordDict []string) bool {
	hashMap := make(map[string]bool)
	for _, w := range wordDict {
		hashMap[w] = true
	}
	// dp的长度比s的长度多一,要存初始的状态
	dp := []bool{}
	for i:=0;i<=len(s);i++{
		dp = append(dp, false)
	}
	// 第一个代表空字符串,总是在字典里是true
	dp[0] = true

	// 遍历整个字符串,i指针的位置比平常的遍历+1, 是因为取子字符串区间 左闭右开, i始终在想要的字符串的右侧
	for i:=1;i<=len(s);i++{
		// j从头移动到i的左侧,每次把0~i的字符串分成两部分
		for j:=0;j<i;j++{
			// 如果S1 ok, S2 在字典里,i指针就ok
			if dp[j] {
				if _, ok:=hashMap[s[j:i]]; ok {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[len(s)]
}