package main

import (
	"fmt"
)

//给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
//示例 1:
//输入: num1 = "2", num2 = "3"
//输出: "6"
//示例 2:
//输入: num1 = "123", num2 = "456"
//输出: "56088"
//说明：
//num1 和 num2 的长度小于110。
//num1 和 num2 只包含数字 0-9。
//num1 和 num2 均不以零开头，除非是数字 0 本身。
//不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。

//https://leetcode-cn.com/problems/multiply-strings/solution/you-hua-ban-shu-shi-da-bai-994-by-breezean/
//乘数 num1 位数为 M，被乘数 num2 位数为 N， num1 x num2 结果 res 最大总位数为 M+N
//num1[i] x num2[j] 的结果为 tmp(位数为两位，"0x","xy"的形式)，其第一位位于 res[i+j]，第二位位于 res[i+j+1]。

func multiply(num1 string, num2 string) string {
	m := len(num1)
	n := len(num2)
	pos := make([]byte, m+n)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mul := (num1[i] - '0') * (num2[j] - '0')

			p1, p2 := i+j, i+j+1
			// 乘积加上之前进位的数字
			sum := mul + pos[p2]

			// xy, x在p1位，y在p2位
			pos[p1] += sum / 10
			pos[p2] = sum % 10
		}
	}
	tmp := ""
	for _, p := range pos {
		if !(len(tmp) == 0 && p == 0) {
			tmp = tmp + fmt.Sprintf("%d", p)
		}
	}
	if len(tmp) == 0 {
		return "0"
	}
	return tmp

}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	l1, l2 := len(num1), len(num2)

	carry := make([]byte, l1+l2)

	for i := l1; i > 0; i-- {
		index := l2 + i - 1
		n1 := num1[i-1] - '0'

		for j := l2; j > 0; j-- {
			n := carry[index] + n1*(num2[j-1]-'0')
			carry[index] = n % 10
			index--
			carry[index] += n / 10
		}
	}

	k := -1
	for i := 0; i < l1+l2; i++ {
		if carry[i] != 0 && k == -1 {
			k = i
		}
		carry[i] += '0'
	}

	return string(carry[k:])
}
