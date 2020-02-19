// 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。
// 示例:
// 输入: [2,1,5,6,2,3]
// 输出: 10
package main

import "fmt"

func largestRectangleArea(heights []int) int {

	if len(heights) == 1 {
		return heights[0]
	}

	maxArea := 0
	for i := 0; i < len(heights); i++ {
		for j := i; j < len(heights); j++ {
			minH := 1 << 32
			for k := i; k <= j; k++ {
				minH = min(minH, heights[k])
			}
			maxArea = max(maxArea, (j-i+1)*minH)
		}
	}
	return maxArea
}

// 分治法
func largestRectangleArea2(heights []int) int {
	return divCalArea(heights, 0, len(heights)-1)
}

func divCalArea(heights []int, start, end int) int {
	if start > end {
		return 0
	}
	// 找到从start到end之间的最小值
	minIndex := start
	for i := start; i <= end; i++ {
		if heights[minIndex] > heights[i] {
			minIndex = i
		}
	}
	// 最大值在
	// 1.给定的start-end范围内,最小值*(start到end的距离)
	// 2.最小值左边他的最大值
	// 3.最小值右边它的最大值
	return max(heights[minIndex]*(end-start+1), max(divCalArea(heights, start, minIndex-1), divCalArea(heights, minIndex+1, end)))

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 找到比当前高度左边第一个小于该高度的下标和右边第一个小于该高度的下标
func largestRectangleArea3(heights []int) int {
	maxAera := 0
	stack := []int{}
	p := 0



	//如果空,放进去
	for p < len(heights) {
		if len(stack) == 0 {
			stack = append(stack, p)
			p++
			// 如果比栈顶大,也放进去
		} else {
			top := stack[len(stack)-1]
			if heights[p] >= heights[top] {
				stack = append(stack, p)
				p++

				// 比栈顶低,这个就是栈顶右边第一个小于栈顶高度的下标
			} else {
				h := heights[stack[len(stack)-1]]
				stack = stack[:len(stack)-1]
				var leftMin, rightMin int
				// 此时的栈顶是左边第一个小于栈顶的元素
				if len(stack) == 0 {
					leftMin = -1
				} else {
					leftMin = stack[len(stack)-1]
				}

				rightMin = p
				// 两个最矮高度中间的区域,不包括两个最矮,所以-1
				maxAera = max(maxAera, h*(rightMin-leftMin-1))
				// 没有p++, 右边元素保持不变,
			}

		}
	}

	for len(stack) != 0 {
		var leftMin, rightMin int

		h := heights[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			leftMin = -1
		} else {
			leftMin = stack[len(stack)-1]
		}

		rightMin = len(heights)
		maxAera = max(maxAera, h*(rightMin-leftMin-1))
	}
	return maxAera
}

// 加两个0进来, 因为距离是相对的,不改变结果. 在最后遍历到0时, 自然会把栈内所有元素弹出
func largestRectangleArea4(heights []int) int {
	stack := []int{}
	dummyH := []int{0}
	dummyH = append(dummyH, heights...)
	dummyH = append(dummyH, 0)
	maxAera := 0

	for i := 0; i < len(dummyH); i++ {
		for len(stack) != 0 && dummyH[stack[len(stack)-1]] > dummyH[i] {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			maxAera = max(maxAera, (i-stack[len(stack)-1]-1)*dummyH[tmp])
		}
		stack = append(stack, i)
	}
	return maxAera

}
