package main

// 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
// 示例 1:
// 输入: [3,2,3]
// 输出: 3
// 示例 2:
// 输入: [2,2,1,1,1,2,2]
// 输出: 2

// 不同的元素互相抵消，最后剩下来的就是多数的数字
func majorityElement(nums []int) int {
	count := 1
	majority := nums[0]
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			majority = nums[i]
		}
		if majority == nums[i] {
			count += 1
		} else {
			count -= 1
		}
	}
	return majority
}
