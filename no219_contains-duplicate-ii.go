package main

// 给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的绝对值最大为 k。
// 示例 1:
// 输入: nums = [1,2,3,1], k = 3
// 输出: true
// 示例 2:
// 输入: nums = [1,0,1,1], k = 1
// 输出: true
// 示例 3:
// 输入: nums = [1,2,3,1,2,3], k = 2
// 输出: false

func containsNearbyDuplicate(nums []int, k int) bool {
	vMap := make(map[int]int)
	for index, value := range nums {
		if oldIndex, ok := vMap[value]; ok {
			if index-oldIndex <= k {
				return true
			} else {
				vMap[value] = index
			}
		} else {
			vMap[value] = index
		}
	}
	return false
}

func containsNearbyDuplicate2(nums []int, k int) bool {
	vMap := make(map[int]int)
	for index, value := range nums {
		for _, mvalue := range vMap {
			if value == mvalue {
				return true
			}
			vMap[index] = value
		}
		if len(vMap) > k {
			vMap[index-k] = 0
		}
	}
	return false
}
