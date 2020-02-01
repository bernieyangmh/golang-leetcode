// 给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素最多出现两次，返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
// 示例 1:
// 给定 nums = [1,1,1,2,2,3],
// 函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。
// 你不需要考虑数组中超出新长度后面的元素。
// 示例 2:
// 给定 nums = [0,0,1,1,1,1,2,3,3],
// 函数应返回新长度 length = 7, 并且原数组的前五个元素被修改为 0, 0, 1, 1, 2, 3, 3 。

package main

import (
	"fmt"
)

func main() {
	a := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	size := removeDuplicates(a)
	res := []int{}
	for i := 0; i < size; i++ {
		res = append(res, a[i])
	}
	fmt.Println(res)

}

// 两个指针,第一个指针x依次遍历数组,第二个指针y维护符合要求的数组结果, 临时数组维护当前数字重复出现的最大次数
// x指针指向的数字与前一位相同,根据临时数组的长度判断是否放进数组
// 不同,根据临时数据和y指针的位置修改原数组, 之后更新临时数组的值
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	var point = 0
	var tmpArr = []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		fmt.Println(i, nums[i], point, tmpArr, nums)
		if nums[i] == nums[i-1] {
			if len(tmpArr) < 2 {
				tmpArr = append(tmpArr, nums[i])
			}
		} else {
			for _, v := range tmpArr {
				nums[point] = v
				point++
			}
			tmpArr = []int{nums[i]}
		}
	}
	// 记得把最后一个临时数组的值修改好
	for _, v := range tmpArr {
		nums[point] = v
		point++
	}
	return point
}
