package main

//给定一个经过编码的字符串，返回它解码后的字符串。
//编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
//你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
//此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
//示例:
//s = "3[a]2[bc]", 返回 "aaabcbc".
//s = "3[a2[c]]", 返回 "accaccacc".
//s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".

import (
	"strconv"
)

// 栈保存之前的结果
func decodeString(s string) string {
	stack := []string{}
	for _, l := range s {
		if string(l) != "]" {
			stack = append(stack, string(l))
		} else {
			word := ""
			left := len(stack) - 1
			for left >= 0 {
				if stack[left] == "[" {
					numPointLeft := left - 1
					for numPointLeft >= 0 {
						if !isNum(stack[numPointLeft]) {
							break
						}
						numPointLeft--
					}
					repeatStr := ""
					for i := numPointLeft + 1; i < left; i++ {
						repeatStr += stack[i]
					}

					repeatNum, _ := strconv.Atoi(repeatStr)
					singWord := ""
					for i := left + 1; i < len(stack); i++ {
						singWord += stack[i]
					}

					for i := 0; i < repeatNum; i++ {
						word += singWord
					}
					// fmt.Println(stack, left, word, singWord)
					stack = stack[:numPointLeft+1]
					stack = append(stack, word)
					// fmt.Println(stack)
					break
				}
				left--
			}
		}
	}
	// fmt.Println(stack)
	res := ""
	for _, w := range stack {
		res += w
	}
	return res
}

func isNum(l string) bool {
	if l == "0" || l == "1" || l == "2" || l == "3" || l == "4" || l == "5" || l == "6" || l == "7" || l == "8" || l == "9" {
		return true
	}
	return false
}

//递归方式
func decodeString(s string) string {
	return dfs(s, 0)[0]
}

func dfs(s string, i int) []string {
	res := ""
	mutil := 0
	for i < len(s) {
		if isNum(string(s[i])) {
			n, _ := strconv.Atoi(string(s[i]))
			mutil = mutil*10 + n
		} else if string(s[i]) == "[" {
			tmp := dfs(s, i+1)
			i, _ = strconv.Atoi(tmp[0])
			for mutil > 0 {
				res += tmp[1]
				mutil--
			}
		} else if string(s[i]) == "]" {
			return []string{strconv.Itoa(i), res}
		} else {
			res += string(s[i])
		}
		i++
	}
	return []string{res}
}
