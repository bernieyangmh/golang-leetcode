package main

//给定一个数组arr，返回子数组的最大累加和
//例如，arr = [1, -2, 3, 5, -2, 6, -1]，所有子数组中，[3, 5, -2, 6]可以累加出最大的和12，所以返回12.
//题目保证没有全为负数的数据
//[要求]
//时间复杂度为O(n)O(n)，空间复杂度为O(1)O(1)

/**
 * max sum of the subarray
 * @param arr int整型一维数组 the array
 * @return int整型
 */
func maxsumofSubarray(arr []int) int {
	// write code here
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	maxValue := 0
	// 直接在原数组上计算
	for i := 1; i < len(arr); i++ {
		arr[i] = max(arr[i], arr[i-1]+arr[i])
		maxValue = max(maxValue, arr[i])
	}
	return maxValue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
