// 给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
// 此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
// 注意:
// 不能使用代码库中的排序函数来解决这道题。
// 示例:
// 输入: [2,0,2,1,1,0]
// 输出: [0,0,1,1,2,2]
// 进阶：
// 一个直观的解决方案是使用计数排序的两趟扫描算法。
// 首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
// 你能想出一个仅使用常数空间的一趟扫描算法吗？
package main

import (
	"fmt"
)

func main() {
	a := []int{2, 0, 2, 1, 1, 0}
	sortColors2(a)
	fmt.Println(a)
}

// 一个数组,第一位记录0的个数,第二位记录1的个数,第三位记录2的个数
func sortColors(nums []int) {
	var cntArr = [3]int{0, 0, 0}
	for _, value := range nums {
		cntArr[value]++
	}
	// 依次修改数组,注意边界条件就好
	for i := 0; i < cntArr[0]; i++ {
		nums[i] = 0
	}
	for i := cntArr[0]; i < cntArr[0]+cntArr[1]; i++ {
		nums[i] = 1
	}
	for i := cntArr[1] + cntArr[0]; i < cntArr[0]+cntArr[1]+cntArr[2]; i++ {
		nums[i] = 2
	}
}

// 3个指针,一个指向0的最右边,一个指向2的最左边,一个是当前指针
// 不停交换,保持上诉条件,直至当前指针与p2指针交汇
func sortColors2(nums []int) {
	var p0, curr, tmp int
	var p2 = len(nums) - 1
	for curr <= p2 {
		fmt.Println(p0, curr, p2, nums)
		if nums[curr] == 0 {
			tmp = nums[p0]
			nums[p0] = nums[curr]
			nums[curr] = tmp
			p0++
			curr++
		} else if nums[curr] == 2 {
			tmp = nums[curr]
			nums[curr] = nums[p2]
			nums[p2] = tmp
			p2--
			// 这里当前指针不能前进, p2指针在之前不一定指的是2,需要cur再判断一次
		} else {
			curr++
		}
	}
}
