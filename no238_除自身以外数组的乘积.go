// 给定长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。
// 示例:
// 输入: [1,2,3,4]
// 输出: [24,12,8,6]
// 说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。
// 进阶：
// 你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）

package main

import (
	"fmt"
)

func main() {

}


// 一个数组保存当前索引左边的乘积之和;一个数组保存当前索引右边的乘积之和;之后遍历两数组依次相乘
// [1,2,3,4]
// [1,1,2,6]
// [24,12,4,1]

// todo 可以只保留一个数组,再第二次遍历的时候直接计算
func productExceptSelf(nums []int) []int {
	
	// 左边数组的第一个是1,右边数组的最后一个也是1
	lArr := []int{1}
	rArr := []int{1}
	res := []int{1}
	
	for i := 1; i < len(nums); i++ {
		lArr = append(lArr, nums[i-1]*lArr[i-1])
		rArr = append(rArr, 1)
		res = append(res, 1)
	}
	for j := len(nums) - 2; j >= 0; j-- {
		rArr[j] = nums[j+1] * rArr[j+1]
	}
	for i := 0; i < len(nums); i++ {
		res[i] = lArr[i] * rArr[i]
	}
	fmt.Println(lArr)
	fmt.Println(rArr)
	return res

}
