// 给定一个整数数组 nums ，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）。
// 示例 1:
// 输入: [2,3,-2,4]
// 输出: 6
// 解释: 子数组 [2,3] 有最大乘积 6。
// 示例 2:
// 输入: [-2,0,-1]
// 输出: 0
// 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
package main

import (
	"fmt"
)

func main() {
	a := []int{1, 0, -1, 2, 3, -5, -2}
	println(maxProduct4(a))

}

// 暴力算一遍
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	var maxNum int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			maxNum = max(maxNum, 0)
			continue
		}
		// 这里tmp初始化当前值的时候要判断一次大小,针对一个数字就是最大的情况 [0, 2]
		var tmp = nums[i]
		maxNum = max(maxNum, tmp)
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == 0 {
				maxNum = max(maxNum, 0)
				break
			}
			tmp = tmp * nums[j]
			maxNum = max(maxNum, tmp)
		}
	}
	return maxNum
}

// 用一个数组保存上一次的结果,每次再相乘
// eg:2,3,-2,4
// [2]					max=2
// [6, 3]				max=6
// [-12,-6,-2]			max=6
// [-48, -24, -8, 4]	max=6
func maxProduct2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var maxNum = nums[0]
	// res 第一个初始化[1],每次循环从res的当前位置拿到上一次循环处理的结果
	// 优化1.res可以只保留(上一次的结果),o(n^2)->o(n)
	res := []int{nums[0]}
	for i := 1; i < len(nums); i++ {

		// 优化2.直接在原有数据上改,不新建数组
		for index, value := range res {
			mutil := value * nums[i]
			maxNum = max(maxNum, mutil)
			res[index] = mutil
		}
		// 包括单个数字是最大值的情况
		maxNum = max(maxNum, nums[i])
		res = append(res, nums[i])
	}
	return maxNum
}

// [1,0,-1,2,3,-5,-2]
func maxProduct3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var maxNum = nums[0]
	// res只保留每次结果的最大数和最小数
	res := []int{nums[0], nums[0]}
	for i := 1; i < len(nums); i++ {
		fmt.Println(i, nums[i], maxNum, res)
		maxNum = max(maxNum, nums[i])
		if nums[i] == 0 {
			i++
			// 0是最后一个,不用走了
			if i == len(nums) {
				break
			} else {
				// 如果是0,重置状态 直接跳过
				// [0, 2, 3] res=[2,2] 判断3
				maxNum = max(maxNum, nums[i])
				res = []int{nums[i], nums[i]}
				continue
			}
		}
		maxNum = max(maxNum, nums[i])
		n1 := res[0] * nums[i]
		n2 := res[1] * nums[i]
		res[0] = min(min(n1, n2), nums[i])
		res[1] = max(max(n1, n2), nums[i])
		maxNum = max(maxNum, res[1])
	}
	return maxNum
}

// res 换成两个数字, 优化性能
func maxProduct4(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var maxNum, imin, imax = nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		fmt.Println(i, nums[i], maxNum, imin, imax)
		maxNum = max(maxNum, nums[i])
		if nums[i] == 0 {
			i++
			// 0是最后一个,不用走了
			if i == len(nums) {
				break
			} else {
				// 如果是0,重置状态 直接跳过
				// [0, 2, 3] res=[2,2] 判断3
				maxNum = max(maxNum, nums[i])
				imin, imax = nums[i], nums[i]
				continue
			}
		}
		n1 := imin * nums[i]
		n2 := imax * nums[i]
		imin = min(min(n1, n2), nums[i])
		imax = max(max(n1, n2), nums[i])
		maxNum = max(maxNum, imax)
	}
	return maxNum
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
