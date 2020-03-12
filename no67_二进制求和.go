package main

import (
	"fmt"
	"math/big"
	"strconv"
)

//给定两个二进制字符串，返回他们的和（用二进制表示）。
//输入为非空字符串且只包含数字 1 和 0。
//示例 1:
//输入: a = "11", b = "1"
//输出: "100"
//示例 2:
//输入: a = "1010", b = "1011"
//输出: "10101"

// 直觉的写法，偏慢, 不改a,b字符串
func addBinary(a string, b string) string {
	aPoint := len(a) - 1
	bPoint := len(b) - 1
	anum, bnum, carry := 0, 0, 0
	res := make([]int, max(len(a), len(b))+1)
	resPoint := len(res) - 1
	for aPoint >= 0 && bPoint >= 0 {
		if string(a[aPoint]) == "0" {
			anum = 0
		} else {
			anum = 1
		}
		if string(b[bPoint]) == "0" {
			bnum = 0
		} else {
			bnum = 1
		}

		res[resPoint] = (carry + anum + bnum) % 2
		carry = (carry + anum + bnum) / 2
		aPoint--
		bPoint--
		resPoint--
	}

	for aPoint >= 0 {
		if string(a[aPoint]) == "0" {
			anum = 0
		} else {
			anum = 1
		}
		res[resPoint] = (carry + anum) % 2
		carry = (carry + anum) / 2
		resPoint--
		aPoint--
	}

	for bPoint >= 0 {
		if string(b[bPoint]) == "0" {
			bnum = 0
		} else {
			bnum = 1
		}
		res[resPoint] = (carry + bnum) % 2
		carry = (carry + bnum) / 2
		resPoint--
		bPoint--
	}
	if carry == 1 {
		res[0] = 1
	} else {
		res = res[1:]
	}
	ans := ""
	for _, n := range res {
		if n == 1 {
			ans += "1"
		} else {
			ans += "0"
		}
	}
	return ans
}

func addBinary2(a string, b string) string {
	var A, _ = new(big.Int).SetString(a, 2)
	var B, _ = new(big.Int).SetString(b, 2)
	return fmt.Sprintf("%b", new(big.Int).Add(A, B))
}

//别人的写法
func addBinary3(a string, b string) string {
	const base = 2
	m, n := len(a), len(b)
	// 这里始终保证i比j大，就不用再两次for循环
	if m > n {
		return addBinary(b, a)
	}
	buf := make([]byte, n+1)
	carry := 0
	for i, j := n-1, m-1; i >= 0; i-- {
		if j >= 0 {
			carry += int(a[j] - '0')
			j--
		}
		carry += int(b[i] - '0')
		buf[i+1] = byte(carry%base + '0')
		carry /= base
	}
	if carry == 0 {
		return string(buf[1:])
	}
	buf[0] = '1'
	return string(buf)
}




func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
