//给定一个未排序的整数数组，找出最长连续序列的长度。
//要求算法的时间复杂度为 O(n)。
//示例:
//输入: [100, 4, 200, 1, 3, 2]
//输出: 4
//解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。

package main



var maxLength int


// 遍历一遍，放进map里
// 对每一个数组里的数字，看他的左右边是否在map里存在
// 如果存在，map标记为已使用，继续递归；如果不存在，返回当前长度
func longestConsecutive(nums []int) int {
	numMap := make(map[int]bool)

	for i:=0;i<len(nums);i++{
		numMap[nums[i]] = true
	}
	maxLength = 0
	curLength := 0
	for i:=0;i<len(nums);i++{
		curLength = Helper(nums[i], 0, numMap)
	}
	maxLength = max(maxLength, curLength)
	return maxLength

}


func Helper(num, curLength int, numMap map[int]bool) int {
	if unused, ok := numMap[num]; ok && unused {
		curLength+=1
		// 这里标记为false，表明已使用
		numMap[num] = false
		maxLength = max(maxLength, curLength)
		// 向左向右看
		curLength = Helper(num-1, curLength, numMap)
		curLength = Helper(num+1, curLength, numMap)
	}
	return curLength
}


// 动态规划的思想，从该数的左边取连续长度，然后+1，更新最大值，更新3个数的连续长度
func longestConsecutive2(nums []int) int {
	numMap := make(map[int]int)
	maxLength := 0
	for _, n := range nums {
		if _, ok:=numMap[n];!ok{
			var left, right int
			if l, ok := numMap[n-1];ok{
				left = l
			} else {
				left = 0
			}

			if r, ok := numMap[n+1];ok{
				right = r
			} else {
				right = 0
			}

			curLength := 1+left+right
			if curLength > maxLength{
				maxLength = curLength
			}
			numMap[n] = curLength
			numMap[n-left] = curLength
			numMap[n+right] = curLength

		}
	}
	return  maxLength

}



func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

