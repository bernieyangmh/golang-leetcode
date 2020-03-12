package main

//
//假设按照升序排序的数组在预先未知的某个点上进行了旋转。
//( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。
//编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。
//示例 1:
//输入: nums = [2,5,6,0,0,1,2], target = 0
//输出: true
//示例 2:
//输入: nums = [2,5,6,0,0,1,2], target = 3
//输出: false

func main() {
	a := []int{5, 6, 7, 1, 2, 3, 4}
	search(a, 1)

}

func search(nums []int, target int) bool {
	length := len(nums)
	if length == 0 {
		return false
	}

	left := 0
	right := length - 1

	for left < right {
		mid := (left + right) >> 1
		//左边部分是有序的
		if nums[mid] > nums[left] {
			// 在左边有序数组里
			if nums[left] <= target && target <= nums[mid] {
				right = mid
			} else {
				left = mid + 1
			}
			// 左边部分无序，右边是有序的
		} else if nums[mid] < nums[left] {
			//在右边有序数组里
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid
			}
			// 旋转点在重复的数字上
		} else {
			if nums[left] == target {
				return true
			} else {
				left = left + 1
			}
		}
	}
	return nums[left] == target
}

// 更易理解
func search2(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	left := 0
	right := len(nums) - 1
	mid := 0

	for left <= right {
		mid = left + (right-left)/2

		if nums[mid] == target {
			return true
		}
		//旋转点在重复数字上
		// 不知道左边有序还是右边有序，直接把第一位去掉
		if nums[left] == nums[mid] {
			left++
			continue
		}
		//左边有序
		if nums[left] < nums[mid] {
			//target在左边
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
			//右边有序
		} else {
			// target在右边部分
			if nums[mid] < target && target <= nums[right] {
				// 在区间长度为2时，left+1成为right，不会死循环
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}
