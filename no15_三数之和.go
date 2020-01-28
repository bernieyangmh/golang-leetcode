// 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。
// 示例：
// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
// 满足要求的三元组集合为：
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]
package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(a))
}

func threeSum(nums []int) [][]int {
	var res [][]int
	var tmp []int
	if len(nums) < 3 {
		return res
	}
	// 排序
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		// 如果当前值就大于0,后面的两个数比当前值还大,不可能等于0
		if nums[0] > 0 {
			break
		}
		// 如果当前值和之前值相等,那么相同的数已经算过一遍,不需要再算
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 两个指针指向 当前值后面的字符串 的两端
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			// 如果求和==0,放进结果集
			if sum == 0 {
				tmp = []int{nums[i], nums[l], nums[r]}
				res = append(res, tmp)
				tmp = []int{}
				// 如果l指针后面的数和l指针指向的数相同,一直后移到不同,避免重复计算
				for l < r && nums[l] == nums[l+1] {
					l += 1
				}
				// 如果r指针前面的数和r指针指向的数相同,一直前移到不同,避免重复计算
				for l < r && nums[r] == nums[r-1] {
					r -= 1
				}
				// 满足之后换下一组,既然和为0, 左右指针都动才有可能为0
				l++
				r--
				// 和小了,左指针后移变大一点
			} else if sum < 0 {
				l += 1
				// 和大了,右指针前移变小一点
			} else if sum > 0 {
				r -= 1
			}
		}
	}
	return res
}

func threeSum2(nums []int) [][]int {
	//先排序
	sort.Ints(nums)

	res := make([][]int, 0)
	f := func(nums []int, begin int, end int, target int) {
		for begin < end { //从两端向中间遍历
			if nums[begin]+nums[end]+target == 0 {
				r := make([]int, 0)
				r = append(r, nums[begin], nums[end], target)
				res = append(res, r)
				//遇到相等的，就快进
				for begin < end && nums[begin] == nums[begin+1] {
					begin++
				}
				for begin < end && nums[end] == nums[end-1] {
					end--
				}
				begin++
				end--
			} else if (target + nums[begin] + nums[end]) < 0 {
				begin++
			} else {
				end--
			}
		}
	}
	l := len(nums)
	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		f(nums, i+1, l-1, nums[i])
	}
	return res
}
