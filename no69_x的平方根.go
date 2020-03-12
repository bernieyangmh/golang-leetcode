package main

import (
	"fmt"
)

//实现 int sqrt(int x) 函数。
//计算并返回 x 的平方根，其中 x 是非负整数。
//由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
//示例 1:
//输入: 4
//输出: 2
//示例 2:
//输入: 8
//输出: 2
//说明: 8 的平方根是 2.82842...,
//由于返回类型是整数，小数部分将被舍去。

func main() {
	mySqrt2(17)

}

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	left := 1
	right := x / 2
	for left < right {
		// 取右中位数
		mid := (left + right + 1) >> 1
		square := mid * mid
		//fmt.Println(left, right, mid, square)
		if square > x {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left
}


func mySqrt2(x int) int {
	if x <= 0 {
		return x
	}
	bigx := float64(x)
	cur := 1.0
	for true {
		pre := cur
		cur = (cur+bigx/cur)/2
		fmt.Println(pre, cur)
		// 精确够高了
		if abs(cur-pre) < 0.000000000001 {
			return int(cur)
		}
	}
	return 0
}

func abs(x float64) float64{
	if x <0{
		return -1*x
	}
	return x

}
