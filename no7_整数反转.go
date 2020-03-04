package main

import (
	"strconv"
	"strings"
)

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//示例 1:
//输入: 123
//输出: 321
//示例 2:
//输入: -123
//输出: -321
//示例 3:
//输入: 120
//输出: 21
//注意:
//假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

// 1.leetcode的32有正负号，所以是1<<31
// 2.负数时x转化为字符串的处理
func reverse(x int) int {
	var (
		flag   bool
		resStr string
	)
	if x > 0 {
		flag = true
	} else {
		x = -1 * x
	}
	s := strconv.Itoa(x)
	for i := len(s) - 1; i >= 0; i-- {
		resStr += string(s[i])
	}
	res, _ := strconv.Atoi(resStr)
	if res >= 1<<31 || res <= (1<<31*-1) {
		return 0
	}
	if !flag {
		return res * -1
	}

	return res

}

// 除10取余， 注意边界
func reverse(x int) int {
	strings.TrimSpace()
	res := 0
	for x != 0 {
		pop := x % 10
		//1<<31-1 = 2147483647  1<<31*-1 = -2147483648
		if res > (1<<31/10) || res == 1<<31/10 && pop > 7 {
			return 0
		}
		if res < (1<<31*-1)/10 || res == (1<<31*-1)/10 && pop < -8 {
			return 0
		}
		res = res*10 + pop
		x /= 10
	}
	return res
}
