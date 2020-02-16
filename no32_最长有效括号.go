// 给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
// 示例 1:
// 输入: "(()"
// 输出: 2
// 解释: 最长有效括号子串为 "()"
// 示例 2:
// 输入: ")()())"
// 输出: 4
// 解释: 最长有效括号子串为 "()()"

package main

// 动态规划
// 当前是 )
//      1.如果前一个是 ( 长度等于前两个dp的长度                   0 1 2 3 4 5 6 7 8
//      2.如果前一个是) 当前)对应括号的位置在 “i-前一个)有效长度-1” ( ( ) ( ( ( ) ) )  8-dp[7]-1 = 3
//                  dp[3] == '(' 则 dp[i] 是前一个')'的有效长度加上 ‘对应括号前面的有效长度’ 即dp[8] = dp[7] + dp[2] + 2
//                  dp[3] == ')' 说明组不成有效长度,归零,后面来查时相当于 ‘对应括号前面的有效长度’ 为0
// 注意边界条件, -2都不能超过0
func longestValidParentheses(s string) int {
	maxlength := 0

	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if string(s[i]) == ")" {
			if string(s[i-1]) == "(" {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && string(s[i-dp[i-1]-1]) == "(" {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			maxlength = max(maxlength, dp[i])
		}
	}
	return maxlength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 用栈保存索引, 抛出时更新最大长度
func longestValidParentheses2(s string) int {
	maxNum := 0
	stack := []int{}
	stack = append(stack, -1)
	for index, l := range s {
		if string(l) == "(" {
			stack = append(stack, index)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				// 避免(()()算作5, 核心
				stack = append(stack, index)
			} else {
				maxNum = max(maxNum, index-stack[len(stack)-1])
			}
		}
	}
	return maxNum
}

// 双变量保存左右括号的数量, 左到右扫一遍,右到左扫一遍
func longestValidParentheses3(s string) int {
	var left, right, maxlength int

	// 左到右,没有到最后不知道是否还有右多余的右括号
	//       所以当右括号数量大于左括号时, 当前的长度判断已经结束;
	//       而左扩号大于有括号数量时,我们还需要接着走

	// 左到右,右到左的区别在于 前者尽量减少多余的左括号,后者减少多余的右括号
	for _, l := range s {
		if string(l) == "(" {
			left++
		} else {
			right++
		}
		// 判断,但并不清空
		if left == right {
			maxlength = max(maxlength, 2*right)
			// 右的比左多,清空重新计算
		} else if right >= left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == "(" {
			left++
		} else {
			right++
		}
		if left == right {
			maxlength = max(maxlength, 2*left)
		} else if left >= right {
			left, right = 0, 0
		}
	}
	return maxlength
}
