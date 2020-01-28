package main

// 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
// 示例 1：
// 输入: "babad"
// 输出: "bab"
// 注意: "aba" 也是一个有效答案。
// 示例 2：
// 输入: "cbbd"
// 输出: "bb"

import "fmt"

func main() {
	fmt.Println(longestPalindrome2("aaaa"))
	a := []int{0, 1, 2}
	fmt.Println(a[1 : 1+2])

}

// 1.暴力法
// 对每一个子字符串判断是否是回文字符串,记录最长字符串
func longestPalindrome(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if isp(s[i:j]) && j-i > len(res) {
				res = s[i:j]
			}
		}
	}
	return res

}

func isp(a string) bool {
	if a == "" {
		return true
	}
	var s, e int
	e = len(a) - 1
	for s < e {
		if a[s] != a[e] {
			return false
		}
		s += 1
		e -= 1
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 动态规划
// 间隔超过2的字符串的判断,可直接由[i+1][j-1] 即右上角的值判断;
// 一个字符串两端相同,里面的子字符串是回文字符串,它就是回文字符串`
// if (s[i] == s[j]) and dp[i + 1][j - 1] {
// dp[i][j] = true
// }
// s = 'aaaaa'
// j   i | 0 | 1 | 2 | 3 | 4 |
// 0     | t |   |   |   |   |
// 1     |   | t |   |   |   |
// 2     |   |   | t |   |   |
// 3     | t |   |   | t |   |
// 4     | t | t |   |   | t |
// [j:3,i:0] <- [j:2,i:1]
// [j:4,i:0] <- [j:3,i:1], [j:4,i:1] < [j:3,i:2]
func longestPalindrome2(s string) string {
	if len(s) < 2 {
		return s
	}
	size := len(s)
	var dp [][]bool
	var max_len = 1
	var start int

	// 指针重合,单个字符串必定是true,构建一个对角线为true的二维数组
	for i := 0; i < size; i++ {
		tmp := []bool{}
		for j := 0; j < size; j++ {
			if j == i {
				tmp = append(tmp, true)
			} else {
				tmp = append(tmp, false)
			}
		}
		dp = append(dp, tmp)
	}

	// 从第二个字符开始,每次以指向头的指针i和指向当前指针的j为字符串,判断其是否是回文
	for j := 1; j < size; j++ {
		for i := 0; i < j; i++ {
			// 如果字符串左右两端一致
			if s[i] == s[j] {
				// 如果是1,2,3,个字符,必定是回文的
				if j-i < 3 {
					dp[i][j] = true
				} else {
					// 更小解 是否是回文的决定当前字符串是否是回文的
					dp[i][j] = dp[i+1][j-i]
				}
				// 不一致肯定不是
			} else {
				dp[i][j] = false
			}
			// 如果当前字符串是回文的,判断是否是最大字符串
			if dp[i][j] {
				cur_len := j - i + 1
				if cur_len > max_len {
					max_len = cur_len
					// 记录最大长度和起始索引,不记录字符串
					start = i
				}
			}
		}
	}
	return s[start : start+max_len]
}

// 中心扩散
func longestPalindrome3(s string) string {
	if len(s) < 2 {
		return s
	}
	var (
		max_len                 = 1
		res                     = string(s[0])
		oddStr, evenStr, maxStr string
		oddLen, evenLen         int
	)
	for i := 0; i < len(s); i++ {
		// 以单个字符串为中心
		oddStr, oddLen = centerSpread(s, len(s), i, i)
		// 以两个字符串为中心
		evenStr, evenLen = centerSpread(s, len(s), i, i+1)

		// 判断奇偶哪个大
		if oddLen > evenLen {
			maxStr = oddStr
		} else {
			maxStr = evenStr
		}
		// 与当前最大字符串判断
		if len(maxStr) > max_len {
			max_len = len(maxStr)
			res = maxStr
		}
	}
	return res
}

func centerSpread(s string, size, left, right int) (str string, len int) {
	var (
		i = left
		j = right
	)
	// 1.左指针到头 2.右指针到头 3.左右指针的字符串相同
	for i >= 0 && j < size && s[i] == s[j] {
		// 向左移动
		i -= 1
		// 向由移动
		j += 1
	}
	// 不是回文的时候,取指针向内一位的长度,返回该字符串长度
	// abacad i=1, j=5; s[2:5]='aca'; j-i-1=3
	return s[i+1 : j], j - i - 1
}
