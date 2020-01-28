// Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Note that an empty string is also considered valid.

// Example 1:

// Input: "()"
// Output: true
// Example 2:

// Input: "()[]{}"
// Output: true
// Example 3:

// Input: "(]"
// Output: false
// Example 4:

// Input: "([)]"
// Output: false
// Example 5:

// Input: "{[]}"
// Output: true

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValid("(()"))
}

func isValid(s string) bool {
	if s == "" {
		return true
	}
	var stackArr []string

	r := strings.Split(s, "")
	fmt.Println(r, len(r))
	for _, i := range r {
		// 进栈 无需判断
		if i == "{" || i == "[" || i == "(" {
			stackArr = append(stackArr, i)
		}
		// 出栈需要判断1.栈是否会空，无匹配字符串
		// 栈顶的元素能否匹配
		// 匹配后出栈

		if i == "}" {
			if len(stackArr) == 0 || stackArr[len(stackArr)-1] != "{" {
				return false
			} else {
				stackArr = stackArr[:len(stackArr)-1]
			}
		}
		if i == "]" {
			if len(stackArr) == 0 || stackArr[len(stackArr)-1] != "[" {
				return false
			} else {
				stackArr = stackArr[:len(stackArr)-1]
			}
		}
		if i == ")" {
			if len(stackArr) == 0 || stackArr[len(stackArr)-1] != "(" {
				return false
			} else {
				stackArr = stackArr[:len(stackArr)-1]
			}
		}

	}

	// 仍有 留存未匹配字符串
	if len(stackArr) > 0 {
		return false
	}
	return true
}
