package main

//给定一个无序的整数数组，找到其中最长上升子序列的长度。
//示例:
//输入: [10,9,2,5,3,7,101,18]
//输出: 4
//解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
//说明:
//可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
//你算法的时间复杂度应该为 O(n2) 。
//进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?

//https://leetcode-cn.com/problems/longest-increasing-subsequence/solution/dong-tai-gui-hua-er-fen-cha-zhao-tan-xin-suan-fa-p/

//	nums	[10,9,2,5,3,7,101,18]
//	dp		[1, 1,1,2,2,3, 4, 4]
func lengthOfLIS(nums []int) int {
	size := len(nums)

	if size < 2 {
		return size
	}

	dp := []int{}
	for i:=0;i<size;i++{
		dp = append(dp, 1)
	}

	//0-i能组成的最大子序
	for i:=1;i<size;i++{
		for j:=0;j<i;j++{
			if nums[j]<nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	res := 0
	for i:=0;i<size;i++{
		res = max(res, dp[i])
	}
	return res


}

// 维护一个数组，1.它是升序 2.最尾端尽可能的小
//[10,9,2,5,3,7,101,18]
//10
//9
//2
//2,5
//2,3
//2,3,7
//2,3,7,101
//2,3,7,18

// 正常是需要二维数组来表示不同长度的上升子序列
// 用二分法查找并更新的话只需要一维
func lengthOfLIS2(nums []int) int {
	tails := make([]int, len(nums))
	res := 0

	// 如果n在数组里,i最后就是n的索引
	// 如果n在区间里,i就是应该更新的索引
	// 如果n在区间左边,i就是0
	// 如果n不在区间右,i最后比res多1,且j仍是res
	for _, n := range nums {
		i, j:=0,res
		for i<j{
			mid := (i+j)/2
			if tails[mid] < n{
				i =  mid + 1
			}else {
				j=mid
			}
		}
		tails[i] =n
		if res == j {
			res++
		}
	}
	return res
}








func max(a, b int) int {
	if a >b {
		return a
	}
	return b

}