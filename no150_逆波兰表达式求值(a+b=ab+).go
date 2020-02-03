// 根据逆波兰表示法，求表达式的值。
// 有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
// 说明：
// 整数除法只保留整数部分。
// 给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
// 示例 1：
// 输入: ["2", "1", "+", "3", "*"]
// 输出: 9
// 解释: ((2 + 1) * 3) = 9
// 示例 2：
// 输入: ["4", "13", "5", "/", "+"]
// 输出: 6
// 解释: (4 + (13 / 5)) = 6
// 示例 3：
// 输入: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
// 输出: 22
// 解释:
//   ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
// = ((10 * (6 / (12 * -11))) + 17) + 5
// = ((10 * (6 / -132)) + 17) + 5
// = ((10 * 0) + 17) + 5
// = (0 + 17) + 5
// = 17 + 5
// = 22
package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN2(a))

}

// 栈思想
// 从前到后遍历,遇到运算符,就把nums数组里的最后两个进行运算,并把结果放到nums数组里
//            遇到数字,就放进nums数组里
func evalRPN(tokens []string) int {
	var nums []string
	// 最后一个一定是运算符
	for i := 0; i < len(tokens); i++ {
		fmt.Println(tokens[i], nums)
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			a, b := nums[len(nums)-2], nums[len(nums)-1]
			nums = nums[:len(nums)-2]
			nums = append(nums, handleOperator(a, b, tokens[i]))
		} else {
			nums = append(nums, tokens[i])
		}
	}
	res, _ := strconv.Atoi(nums[0])
	return res
}

func handleOperator(a, b, oper string) string {
	i1, _ := strconv.Atoi(a)
	i2, _ := strconv.Atoi(b)

	switch oper {
	case "+":
		return strconv.Itoa(i1 + i2)
	case "-":
		return strconv.Itoa(i1 - i2)
	case "*":
		return strconv.Itoa(i1 * i2)
	case "/":
		return strconv.Itoa(i1 / i2)
	}
	return ""
}

// 优化了上面第一次写的代码,不需要数字字符来回转,集中在一个函数里,逻辑更简单
func evalRPN2(tokens []string) int {
	stack := []int{}
	for i := 0; i < len(tokens); i++ {
		fmt.Println(i, tokens[i], stack)
		switch tokens[i] {
		case "+":
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a+b)
		case "-":
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a-b)
		case "*":
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a*b)
		case "/":
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a/b)
		default:
			num, _ := strconv.Atoi(tokens[i])
			stack = append(stack, num)
		}
	}
	return stack[0]
}
