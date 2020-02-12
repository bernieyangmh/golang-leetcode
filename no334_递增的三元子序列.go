// 给定一个未排序的数组，判断这个数组中是否存在长度为 3 的递增子序列。
// 数学表达式如下:
// 如果存在这样的 i, j, k,  且满足 0 ≤ i < j < k ≤ n-1，
// 使得 arr[i] < arr[j] < arr[k] ，返回 true ; 否则返回 false 。
// 说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1) 。
// 示例 1:
// 输入: [1,2,3,4,5]
// 输出: true
// 示例 2:
// 输入: [5,4,3,2,1]
// 输出: false
package main

import (
	"fmt"
)

func main() {
	a := []int{5, 1, 5, 5, 2, 5, 4}
	fmt.Println(increasingTriplet(a))
}

// 数组的求值使用动态规划
// 一个数组表示每个位置左边的最小值,一个数组表示每个位置右边的最大值
// 最后每个位置拿出左边最小值,当前值,右边最大值进行比较
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	leftArr := make([]int, len(nums))
	leftArr[0] = nums[0]
	// minIndex := 0
	for i := 1; i < len(nums); i++ {
		leftArr[i] = min(nums[i], leftArr[i-1])
	}

	rightArr := make([]int, len(nums))
	rightArr[len(nums)-1] = nums[len(nums)-1]

	for i := len(nums) - 2; i >= 0; i-- {
		rightArr[i] = max(nums[i], rightArr[i+1])
	}

	for i := 1; i <= len(nums)-2; i++ {
		if leftArr[i] < nums[i] && nums[i] < rightArr[i] {
			return true
		}
	}

	return false
}

// 为什么不会出现更新最小值之后漏掉的情况?,
// 可以想象数组的每个一个位置代表一座山,位置的数值代表山顶
// 我们的标准就是一条向左下倾斜的一条直线,无论倾斜多少角度,只要能划一条倾斜的线,就说明满足
// 这条线从天空慢慢降下来,因为后面的山比前面的山低才重新判断,所以线降到后面山之前,肯定已经判断前面的山是否满足升高的情况

//
// |                                        |  |
// |                                   /   |||||
// ||                             |  / | | |||||
// || |   /|                     ||| | | | |||||
// ||||  /||                     ||| | | | |||||
// |||| /||||                    |||/||| |||||||
// ||||||||||||                  |||||||||||||||
// ||||||||||||                  |||||||||||||||
//       线                            线

func increasingTriplet2(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	// min, mid初始化需要是最大值,我们赋值的条件是当前值比最小值和中间值小

	minNum := 1 << 32
	midNum := 1 << 32
	for i := 0; i < len(nums); i++ {
		// 当前值比中间值大
		if nums[i] > midNum {
			return true
			// 当前值比最小值小,重置最小值
		} else if nums[i] < minNum {
			minNum = nums[i]
			// 中间值比最小值大的基础上尽可能小
		} else if minNum < nums[i] && nums[i] <= midNum {
			midNum = nums[i]
		}
	}
	return false
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
