package main

import (
	"fmt"
)

//题目描述
//请实现有重复数字的有序数组的二分查找。
//输出在数组中第一个大于等于查找值的位置，如果数组中不存在这样的数，则输出数组长度加一。
//示例1
//输入
//复制
//5,4,[1,2,4,4,5]
//返回值
//复制
//3

/**
 * 二分查找
 * @param n int整型 数组长度
 * @param v int整型 查找值
 * @param a int整型一维数组 有序数组
 * @return int整型
 */

func main() {
	a := []int{4,5}
	fmt.Println(binaryFind(2,5,a))
}

func binaryFind(n int, v int, a []int) int {
	// write code here

	if n == 0 {
		return 1
	}
	if a[n-1] < v {
		return n + 1
	}

	left := 0
	right := n - 1

	for left < right {
		mid := left + (right-left)/2
		if v <= a[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	fmt.Println(left, right)
	return left + 1
}
