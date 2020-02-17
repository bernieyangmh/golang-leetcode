package main

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
// 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。
// 示例:
// 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出: 6

// https://leetcode-cn.com/problems/trapping-rain-water/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-w-8/

import (
	"fmt"
)

func main() {
	a := []int{4, 1, 1, 2, 1, 1, 6}
	println(trap3(a))

}

// 对每个元素,找当前位置的左边右边最大值,取两者最小值相差
func trap(height []int) int {
	volume := 0
	size := len(height)
	for i := 1; i < size-1; i++ {
		var maxLeft, maxRight int
		for j := i; j >= 0; i-- {
			maxLeft = max(maxLeft, height[j])
		}
		for j := i; j < size; j++ {
			maxRight = max(maxRight, height[j])
		}
		volume += min(maxLeft, maxRight) - height[i]
	}
	return volume
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

func trap2(height []int) int {
	size := len(height)

	if size == 0 {
		return 0
	}
	volume := 0

	leftMax := make([]int, size)
	rightMax := make([]int, size)

	// 从左到右遍历，找到每个位置的左边最大值
	leftMax[0] = height[0]
	for i := 1; i < size; i++ {
		leftMax[i] = max(height[i], leftMax[i-1])
	}

	// 从右到左遍历，找到每个位置的右边最大值
	rightMax[size-1] = height[size-1]
	for i := size - 2; i >= 0; i-- {
		rightMax[i] = max(height[i], rightMax[i+1])
	}

	// 同1暴力递归，两个最大值的最小值减当前值
	// todo 数组内的每个最大值只用一次,可以优化
	for i := 1; i < size-1; i++ {
		volume += min(leftMax[i], rightMax[i]) - height[i]
	}
	return volume

}

// 遍历高度
// 如果当前高度小于栈顶的墙高度，说明会有积水，将当前墙的下标入栈
// 如果当前高度大于栈顶的墙高度，说明之前的积水到这形成一个“湖”，计算体积; 计算完，就把当前的墙继续入栈，作为新的积水的墙。
func trap3(height []int) int {
	volume := 0
	current := 0
	stack := []int{}
	for current < len(height) {
		// fmt.Println(stack)
		// 大于栈顶且栈非空就pop (右边的墙大于左边的墙)
		for len(stack) > 0 && height[current] > height[stack[len(stack)-1]] {
			// 栈顶值
			top := stack[len(stack)-1]
			// pop
			stack = stack[:len(stack)-1]

			// 栈空,左边没可匹配的
			if len(stack) == 0 {
				break
			}
			// 两墙距离, 即抛出元素的宽
			distance := current - stack[len(stack)-1] - 1
			// 高度差是 左边墙或当前墙的高度 减去 抛出的元素的高度
			bound_height := min(height[current], height[stack[len(stack)-1]]) - height[top]
			volume = volume + distance*bound_height
			fmt.Println(volume, stack[len(stack)-1], current, distance)
		}
		// 小于等于栈顶就push
		stack = append(stack, current)
		current++
	}
	return volume
}

func trap4(height []int) int {
	maxLeft, maxRight, volume := 0, 0, 0
	left := 1
	right := len(height) - 2
	for i := 1; i < len(height)-1; i++ {

		// 动态规划时我们要每个元素左右最大值的最小值
		// leftMax[i] = max(height[i], leftMax[i-1]) 
		// rightMax[i] = max(height[i], rightMax[i+1])

		// height[left-1] < height[right+1] ==> max_left < max_right。
		if height[left-1] < height[right+1] {
			maxLeft = max(maxLeft, height[left-1])
			minNum := maxLeft
			if minNum > height[left] {
				volume = volume + (minNum - height[left])
			}
			left++
		} else {
			maxRight = max(maxRight, height[right+1])
			minNum := maxRight
			if minNum > height[right] {
				volume = volume + (minNum - height[right])
			}
			right--
		}
	}
	return volume
}
