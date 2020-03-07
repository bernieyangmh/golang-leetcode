package main

import (
	"fmt"
)

//给定一个非负整数数组，你最初位于数组的第一个位置。
//数组中的每个元素代表你在该位置可以跳跃的最大长度。
//你的目标是使用最少的跳跃次数到达数组的最后一个位置。
//示例:
//输入: [2,3,1,1,4]
//输出: 2
//解释: 跳到最后一个位置的最小跳跃数是 2。
// 从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
//说明:
//假设你总是可以到达数组的最后一个位置。

func main()  {
	a := []int{2,3,1,1,4}
	fmt.Println(jump(a))

}




// 贪心算法
// 注意边界条件
func jump(nums []int) int {
	res := 0
	end := 0
	maxSpan := 0
	for i := 0; i < len(nums)-1; i++ {
		//记录在i-end之间下次起跳的最大位置
		maxSpan = max(maxSpan, nums[i]+i)
		// 没到最后，我还需要从上一次的maxspan即end再次起跳，到了最后，就不用跳了
		if i == end {
			fmt.Println(i, end)
			end = maxSpan
			res++
		}
	}
	return res
}

//从左向右看，找到第一个能跳到该位置的位置
func jump2(nums []int) int {
	pos := len(nums)-1
	res := 0
	for pos != 0 {
		for i:=0;i<pos;i++{
			if nums[i]+i >= pos {
				pos = i
				res ++
				break
			}
		}
	}
	return res
}



func max(a, b int) int {

	if a > b {
		return a
	}
	return b
}
