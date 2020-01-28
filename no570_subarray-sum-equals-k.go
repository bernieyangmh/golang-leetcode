// 给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
// 示例 1 :
// 输入:nums = [1,1,1], k = 2
// 输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
package main

// 暴力计算,对每个元素，计算该元素与之后元素的所有组合的结果
func subarraySum(nums []int, k int) int {
	var count int
	for index, _ := range nums {
		var sum int
		sum += nums[index]
		// 单个元素符合k
		if sum == k {
			count += 1
		}
		for j := index + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count += 1
			}
		}
	}
	return count
}

// k是7, 第2位累加是4, 第5位累加是4, 在第8位;累加是11
// 则从2-8是符合要求的子数组,从5-8也是符合要求的子数组
func subarraySum2(nums []int, k int) int {
	var sum, count int
	var HashMap = make(map[int]int)
	for _, value := range nums {
		sum += value
		if sum == k {
			count += 1
		}
		if mv, ok := HashMap[sum-k]; ok {
			count += mv
		}
		if mv2, ok := HashMap[sum]; ok {
			HashMap[sum] = mv2 + 1
		} else {
			HashMap[sum] = 1
		}

	}
	return count
}
