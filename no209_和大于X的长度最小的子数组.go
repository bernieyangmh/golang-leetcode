// 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。如果不存在符合条件的连续子数组，返回 0。
// 示例:
// 输入: s = 7, nums = [2,3,1,2,4,3]
// 输出: 2
// 解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
package main

import (
	"fmt"
)

func main() {
	a := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen3(7, a))
}

// 暴力遍历,
// 2 3 1 2
// 3 1 2 4
// 1 2 4
// 2 4 3
// 4 3
func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := helper(nums, s, 0, 0)
	for i := 1; i < len(nums); i++ {
		tmpMin := helper(nums, s, i, 0)
		if tmpMin != 0 {
			res = min(res, tmpMin)
		}
	}
	return res

}

func helper(nums []int, s, start, sum int) int {
	for i := start; i < len(nums); i++ {
		sum += nums[i]
		if sum >= s {
			return i - start + 1
		}
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 滑动窗口,维持一个和小于7的数组,
// 如果和小于目标,向右扩张;如果大于目标,左边收缩直至小于目标
// 这个数组在大于目标时的最小长度就是所求
// 特殊边界 1.windos大于目标时长度为1,就不用继续判断 2.int默认值是0, 为了避免s小于0的情况,初始化windows是第一个数字
func minSubArrayLen2(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if nums[0] >= s {
		return 1
	}
	minLength := 0
	sum := nums[0]
	windows := []int{nums[0]}
	flag := false
	for i := 1; i < len(nums); i++ {
		windows = append(windows, nums[i])
		sum += nums[i]
		fmt.Println(i, windows, sum)
		// windows之和小于s, 继续向右扩张
		if sum < s {
			continue
			// windows目前已大于s, 左侧尽可能收缩
		} else {
			// 因为minlength初始化,我们需要一个flag在第一次的时候将第一个满足的长度值赋给minLength
			if !flag {
				minLength = len(windows)
				flag = true
			}
			for sum >= s {
				fmt.Println(windows)
				minLength = min(minLength, len(windows))
				sum -= windows[0]
				windows = windows[1:]
			}
		}
	}
	return minLength
}

// 原理同2, 换成指针
func minSubArrayLen3(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if nums[0] >= s {
		return 1
	}
	minLength := 1
	sum := nums[0]
	windowsStart := 0
	windowsEnd := 1
	flag := false
	for i := 1; i < len(nums); i++ {
		windowsEnd++
		sum += nums[i]
		// windows之和小于s, 继续向右扩张
		if sum < s {
			continue
			// windows目前已大于s, 左侧尽可能收缩
		} else {
			if !flag {
				minLength = windowsEnd - windowsStart
				flag = true
			}
			for sum >= s {
				sum -= nums[windowsStart]
				windowsStart++
			}
			// 这时指针内的数组小于s,左侧指针向左移动一格就是大于等于s,判断左移后长度是否是最小值
			minLength = min(minLength, windowsEnd-windowsStart+1)
		}
	}
	if windowsStart == 0 && windowsEnd == len(nums) {
		return 0
	}
	return minLength
}
