package main

import "fmt"

func main() {
	a := []int{1, 2, 7, 4, 3, 1}
	nextPermutation(a)
	fmt.Println(a)

}

// 实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
// 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
// 必须原地修改，只允许使用额外常数空间。
// 以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
// 1,2,3 → 1,3,2
// 3,2,1 → 1,2,3
// 1,1,5 → 1,5,1

// 从右向左扫描,找到第一个破坏递增的数字 x, 如果没有,说明该排列已经是最大值,反转即是最小值 eg:4321->1234          x
// 该数字x右侧的数组从左到右降序排列已经是最大值,加上该数字x的数组一定能找到更大的数字的排列顺序(因为x破坏了排序) eg:12531 -> 13125;  531 ok, 2531 wrong
// 再从右开始扫描,找到比x大的数字y; x右侧的数组从右向左逐渐增加,在x下降,必定有一个y符合要求 y>x, 《y是比x大且尽可能小的值》,y右边的值 <= x
// 	   x     y
// [1, 2, 5, 3, 1]
// [1, 3, 5, 2, 1]
// [1, 3, 1, 2 ,5]
// 交换x,y的位置, 反转x右侧的数组,不包括x.该数组从右向左仍是递增的,反转即是该数组排列的最小值
func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	var (
		first  = -1
		second = -1
	)
	size := len(nums)

	// 注意这里,不是从最右边
	for i := size - 2; i >= 0; i-- {
		// i 比 后面的数字要小,破坏了从右到左的递减顺序
		if nums[i] < nums[i+1] {
			first = i
			break
		}
	}
	// first 仍是-1,代表没找到破坏排序的数字,当前数组已经是最大值, 直接反转
	if first == -1 {
		reverse(nums)
		return
	}
	// 从右到左,找到比x大的数字y, second指针指向y
	for i := size - 1; i >= 0; i-- {
		if nums[i] > nums[first] {
			second = i
			break
		}
	}
	// 交换
	nums[first], nums[second] = nums[second], nums[first]
	// 反转x左边的数组
	reverse(nums[first+1 : size])
}

func reverse(nums []int) {
	s := 0
	e := len(nums) - 1
	for s < e {
		nums[s], nums[e] = nums[e], nums[s]
		s++
		e--
	}
}
