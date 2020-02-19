// 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
// 说明：解集不能包含重复的子集。
// 示例:
// 输入: nums = [1,2,3]
// 输出:
// [
//   [3],
//   [1],
//   [2],
//   [1,2,3],
//   [1,3],
//   [2,3],
//   [1,2],
//   []
// ]

package main

import (
	"fmt"
)

func main() {
	a := []int{9, 0, 3, 5, 7}
	fmt.Println(subsets(a))

}

// 每次从前一个拿结果
func subsets(nums []int) [][]int {
	res := [][]int{}
	res = append(res, []int{})

	for _, i := range nums {
		// 这里不拷贝也可以
		tmp := res
		for _, a := range tmp {
			// 这里一定要copy
			arr := append([]int{}, a...)
			res = append(res, append(arr, i))
		}
		// fmt.Println(res, i, res)
	}
	return res
}
