// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
// ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
// 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
// 你可以假设数组中不存在重复的元素。
// 你的算法时间复杂度必须是 O(log n) 级别。
// 示例 1:
// 输入: nums = [4,5,6,7,0,1,2], target = 0
// 输出: 4
// 示例 2:
// 输入: nums = [4,5,6,7,0,1,2], target = 3
// 输出: -1

package main

import "fmt"

func main() {
	a := []int{-1, 2}
	target := 2
	fmt.Println(search(a, target))

}

// 直接二分法,根据start,mid,end 指向的值每次判断target落在哪里
// 由于题目说数字了无重复，举个例子：
// 1 2 3 4 5 6 7 可以大致分为两类，
// 第一类 2 3 4 5 6 7 1 这种，也就是 nums[start] <= nums[mid]。此例子中就是 2 <= 5。
// 这种情况下，前半部分有序。因此如果 nums[start] <=target<nums[mid]，则在前半部分找，否则去后半部分找。
// 第二类 6 7 1 2 3 4 5 这种，也就是 nums[start] > nums[mid]。此例子中就是 6 > 2。
// 这种情况下，后半部分有序。因此如果 nums[mid] <target<=nums[end]，则在后半部分找，否则去前半部分找。
func search(nums []int, target int) int {
	size := len(nums)
	if size == 0 {
		return -1
	}
	l := 0
	r := size - 1
	// 左指针位置再走 右指针与左指针间隔的1/2位
	mid := l + (r-l)/2
	// 注意条件, 没有等于会忽略size=1的情况
	for l <= r {
		if nums[mid] == target {
			return mid
		}
		// 左部分是排序的, 《因为已经判断了mid的值,所以指针移动是在mid两侧》
		// 小于等于,在两个数的时候仍能和逻辑匹配
		if nums[l] <= nums[mid] {
			// target在左边, 右指针移到mid指针的左边
			if nums[l] <= target && target <= nums[mid] {
				r = mid - 1
				// target在右边, 左指针移到mid指针的右边
			} else {
				l = mid + 1
			}
			// 右部分是排序的
		} else {
			if nums[mid] <= target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		mid = l + (r-l)/2
	}
	return -1
}
