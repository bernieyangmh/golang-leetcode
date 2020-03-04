package main

import (
	"sort"
)

//给定一个整数数组，你需要寻找一个连续的子数组，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
//你找到的子数组应是最短的，请输出它的长度。
//示例 1:
//输入: [2, 6, 4, 8, 10, 9, 15]
//输出: 5
//解释: 你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
//说明 :
//输入的数组长度范围在 [1, 10,000]。
//输入的数组可能包含重复元素 ，所以升序的意思是<=。

// 两套循环, 超时
func findUnsortedSubarray(nums []int) int {
	l := len(nums)
	r := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				r = max(r, j)
				l = min(l, i)
			}
		}
	}
	if r-l < 0 {
		return 0
	}
	return r - l + 1
}

// 排序，如果和排好序的不一致，找到最小和最大的不一致索引
func findUnsortedSubarray2(nums []int) int {
	copyNums := append([]int{}, nums...)

	sort.Ints(nums)

	l := len(nums)
	r := 0
	for i := 0; i < len(copyNums); i++ {
		if nums[i] != copyNums[i] {
			l = min(l, i)
			r = max(r, i)
		}
	}
	if r-l < 0 {
		return 0
	}
	return r - l + 1
}

// 用一个栈
// 第一遍顺序，如果当前数比栈顶要大，入栈，否则不断弹出直到大于栈顶，找到最左边的坐标
// 第二遍逆序，如果当前数比栈顶要小，入栈，否则不断弹出直到小于栈顶，找到最右边的坐标
func findUnsortedSubarray3(nums []int) int {
	stack := []int{}
	l := len(nums)
	r := 0
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			l = min(l, top)
		}
		stack = append(stack, i)
	}
	stack = []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			r = max(r, top)
		}
		stack = append(stack, i)

	}
	if r-l < 0 {
		return 0
	}
	return r - l + 1
}

// 局部最大值最小值
func findUnsortedSubarray4(nums []int) int {
	n := len(nums)

	//数组有序时也能return正确值
	start := -1
	end := -2

	minNum := nums[n-1]
	maxNum := nums[0]

	// 学到了顺序逆序的一种写法
	for i, pos:=0, 0; i<n;i++{
		pos = n-1-i
		maxNum = max(maxNum, nums[i])
		minNum = min(minNum, nums[pos])
		//当前值比最大值下，它是应该修改数组的最左端
		if nums[i] < maxNum {
			end = i
		}
		if nums[pos] > minNum {
			start = pos
		}
	}
	return end - start + 1
}

func max(a, b int) int {

	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
