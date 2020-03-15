package main

import (
	"sort"
)

//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。
//找出那个只出现了一次的元素。
//说明：
//你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
//示例 1:
//输入: [2,2,3,2]
//输出: 3
//示例 2:
//输入: [0,1,0,1,0,1,99]
//输出: 99

// nlogn
func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sort.Ints(nums)
	pre := nums[0]

	if nums[0] != nums[1] {
		return nums[0]
	}
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] != pre && nums[i] != nums[i+1] {
			return nums[i]
		}
		pre = nums[i]
	}
	return nums[len(nums)-1]
}

// 定义一种运算属于三进制的异或
func singleNumber2(nums []int) int {
	ones, twos, threes := 0, 0, 0
	for _, n := range nums {
		twos |= ones & n
		ones ^= n
		threes = ones & twos
		ones &= ^threes
		twos &= ^threes
	}
	return ones

}

//所有不是 3 的倍数的列写 1，其他列写 0 ，就找到了这个出现 1 次的数
func singleNumber3(nums []int) int {
	ans := 0
	for i := 0; i < 32; i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[j]>>uint(i)&1 == 1 {
				count++
			}
		}
		if count%3 != 0 {
			ans = ans | (1 << uint(i))
		}
	}
	// 考虑负数的情况
	//-4 = 1111111111111111111111111111111111111111111111111111111111111100
	if ans > (1<<31 - 1) {
		return -1 * (1<<32 - ans)
	}
	return ans
}
