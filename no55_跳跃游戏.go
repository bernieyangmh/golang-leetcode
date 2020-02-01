// 给定一个非负整数数组，你最初位于数组的第一个位置。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 判断你是否能够到达最后一个位置。
// 示例 1:
// 输入: [2,3,1,1,4]
// 输出: true
// 解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
// 示例 2:
// 输入: [3,2,1,0,4]
// 输出: false
// 解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。

package main

func main() {
	a := []int{1}
	println(canJump2(a))

}

// 自底向上的动态规划
func canJump(nums []int) bool {
	return back(nums, len(nums)-1, false)
}

// 一个指针x从后向前遍历,如果x走到头,说明可以
// 指针y从当前指针x的前一位开始走到头,如果y指向的数字可以跳到x,则递归y指向的数字,看前面是否可以跳到y指向数字的数字.
func back(nums []int, index int, flag bool) bool {
	// 如果有一条路径可以走到头,就不需再判断
	// 如果去掉,相当于很多路径重复计算
	if flag {
		return flag
	}
	// x走到头, 这条路径ok
	if index == 0 {
		flag = true
		return flag
	}
	// y从x前一个开始向前走
	for i := index - 1; i >= 0; i-- {
		// y指向的位置i跳nums[i]个距离,一定能到index x指针上
		if i+nums[i] >= index {
			// 判断y指向的数字
			flag = back(nums, i, flag)
			return flag
			// 否则跳过,y前移
		} else {
			continue
		}
	}
	return flag
}

// 贪心算法
//      g g g x
// [2,3,1,1,1,4]
// 在上面的情况,x指向4时,第3,4,5位均可以,这时没有必要进入4,5位再判断一次前面的3,4是否能否走到.
// 所以我们保存一个指针y,y永远指向可以走到x的最左端位置,在上面例子中y指向第三个位置, 避免重复计算
func canJump2(nums []int) bool {
	if len(nums) < 1 {
		return true
	}
	var lastPos = len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		// nums数组非负,lastpos右边的位置一定满足
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}
	return lastPos == 0
}
