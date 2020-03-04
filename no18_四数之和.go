package main

import (
	"sort"
)

//给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
//注意：
//答案中不可以包含重复的四元组。
//示例：
//给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
//满足要求的四元组集合为：
//[
//[-1,  0, 0, 1],
//[-2, -1, 1, 2],
//[-2,  0, 0, 2]
//]

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	n := len(nums)

	for i := 0; i < n-3; i++ {
		//  重复的不需要再判断
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		//从i之后的4个数之后大于target，后面一定都大于target就不用再判断了
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		//拿i和最大的三个数之后，都小于目标值；肯定达不到target
		if nums[i]+nums[n-1]+nums[n-2]+nums[n-3] < target {
			continue
		}
		//第二个数
		//这里相当于三数之和
		for j := i + 1; j < n-2; j++ {
			//这三个判断同上
			if j-i > 1 && nums[j] == nums[j-1] {
				continue
			}
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			if nums[i]+nums[j]+nums[n-1]+nums[n-2] < target {
				continue
			}
			left := j + 1
			right := n - 1
			for left < right {
				tmp := nums[i] + nums[j] + nums[left] + nums[right]
				if tmp == target {
					tmpList := []int{nums[i], nums[j], nums[left], nums[right]}
					res = append(res, tmpList)
					// 如果left指针右边的数相同，直接跳过避免重复
					for left < right && nums[left] == nums[left+1] {
						left += 1
					}
					// 如果right指针左边的数相同，直接跳过避免重复
					for left < right && nums[right] == nums[right-1] {
						right -= 1
					}
					//如果这次left和right符合目标值，缩小两位也可能是 1234 1+4 = 2+3
					left += 1
					right -= 1
				} else if tmp > target {
					right -= 1
				} else {
					left += 1
				}
			}
		}
	}
	return res
}
