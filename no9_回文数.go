package main

import (
	"fmt"
)

//判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//示例 1:
//输入: 121
//输出: true
//示例 2:
//输入: -121
//输出: false
//解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
//示例 3:
//输入: 10
//输出: false
//解释: 从右向左读, 为 01 。因此它不是一个回文数。
//

//注意各种数的处理方法

func main()  {
	a := 1235321
	isPalindrome(a)
}
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	div := 1
	//最大位
	for x/div >= 10 {
		div *= 10
	}
	for x> 0{
		left := x/div
		right := x %10
		if left != right {
			return false
		}
		//去掉第一位%，去掉最后一位/
		x =  (x%div)/10
		//去掉两位所以/100
		div /= 100
	}
	return true
}