package main

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
//如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//你可以假设数组中无重复元素。
//示例 1:
//输入: [1,3,5,6], 5
//输出: 2
//示例 2:
//输入: [1,3,5,6], 2
//输出: 1
//示例 3:
//输入: [1,3,5,6], 7
//输出: 4
//示例 4:
//输入: [1,3,5,6], 0
//输出: 0

func searchInsert(nums []int, target int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	if nums[size-1] < target {
		return size
	}
	left := 0
	right := size - 1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

//链接：https://leetcode-cn.com/problems/search-insert-position/solution/te-bie-hao-yong-de-er-fen-cha-fa-fa-mo-ban-python-/
func searchInsert2(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	left := 0
	// 有可能插入在最后面，所以right 是length
	right := length

	for left < right {
		mid := left + (right-left)/2;
		//数组分成了两个部分 [left:mid], [left+1:right]
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func searchInsert3(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	left := 0
	right := length
	for left < right {
		mid := left + (right-left)/2;
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			return mid
		} else {
			right = mid
		}
	}
	return left
}
