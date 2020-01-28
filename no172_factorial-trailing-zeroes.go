package main

// 给定一个整数 n，返回 n! 结果尾数中零的数量。
// 示例 1:
// 输入: 3
// 输出: 0
// 解释: 3! = 6, 尾数中没有零。
// 示例 2:
// 输入: 5
// 输出: 1
// 解释: 5! = 120, 尾数中有 1 个零.

// 后面有多少个零意味着有多少个10相乘
// 10 = 2 * 5， 2的个数又远多于5，肯定够用，即10的个数==5的个数
// 问题转化为n里又多少个5
func trailingZeroes(n int) int {
	var result int
	for n > 0 {
		tmp := n / 5
		result += tmp
		n = tmp
	}
	return result
}
