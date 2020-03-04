package main

import (
	"sort"
)

//给定一个包括 n 个整数的数组 nums 和 一个目标值 target。
//找出 nums 中的三个整数，使得它们的和与 target 最接近。
//返回这三个数的和。假定每组输入只存在唯一答案。
//例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
//与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).





func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]
	//对每一位都进行判断
	for i := 0; i < len(nums); i++ {
		start, end := i+1, len(nums)-1
		//这里小于，最后一两位就bu需判断
		for start < end {
			sum := nums[start] + nums[end] + nums[i]
			// 更接近，就更新答案
			if abs(target-sum) < abs(target-ans) {
				ans = sum
			}
			// 比目标值大，end指针向左移减少
			if sum > target {
				end--
				//比目标值小，end指针向右移增大
			} else if sum < target {
				start++
			} else {
				return ans
			}
		}
	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}
