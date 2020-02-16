// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 注意空字符串可被认为是有效字符串。

package main

// 用栈保存结果,注意边界条件
func isValid(s string) bool {
	stack := []string{}

	for _, v := range s {
		// fmt.Println(stack, string(v))
		switch string(v) {
		// 左边符号直接添加
		case "(":
			stack = append(stack, string(v))
		case "{":
			stack = append(stack, string(v))
		case "[":
			stack = append(stack, string(v))

			// 1.栈长度为0 --> false
			// 2.栈顶不是对应的左符号 --> false
			// 3.栈顶是对应的左符号 --> 抛出栈顶
		case ")":
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == "(" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case "}":
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == "{" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case "]":
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == "[" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
