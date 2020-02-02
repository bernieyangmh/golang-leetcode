// 给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
// 说明：解集不能包含重复的子集。
// 示例:
// 输入: [1,2,2]
// 输出:
// [
//   [2],
//   [1],
//   [1,2,2],
//   [2,2],
//   [1,2],
//   []
// ]
package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(a))

}

// 回溯算法,剪枝的逻辑需要注意
// 先排序数组,避免重复的情况
// 当前结点之前的数字不进入(进入下一层前判断 深度), 该结点和前一个相同跳过(在同一层遍历时判断 广度), 每进入一个结点,新生成一个数组放进结果集
// nums = [1,2,2,4]
//   1     2    4
//  /|     |
// 2 4     4
// |
// 4

// ps:注意go的深拷贝,和二维数组传指针
func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	dfs(0, nums, []int{}, &res)
	return res
}

func dfs(s int, nums, path []int, res *[][]int) {
	// 走到数组最尾
	if s == len(nums) {
		return
	}
	for i := s; i < len(nums); i++ {

		//不是第一个并且和前一个数字相同则跳过,不考虑该结点
		if i != s && nums[i] == nums[i-1] {
			continue
		}
		// 深拷贝之前走过的路径, 将当前的结点加入到路径中,放进结果集
		tmp := append([]int(nil), path...)
		tmp = append(tmp, nums[i])
		*res = append(*res, tmp)
		dfs(i+1, nums, tmp, res)
	}
}

// 用hashmap 存储每个数字和出现的次数,然后组合 eg:[1,2,2,4]
// []

// []
// [1]

// []
// [1]
// [2] [1,2]

// []
// [1]
// [2] [1,2]
// [2] [1,2,2]

// []
// [1]
// [2] [1,2]
// [2,2] [1,2,2]
// [4]
// [1,4]
// [2,4] [1,2,4]
// [2,2,4] [1,2,2,4]
func subsetsWithDup2(nums []int) [][]int {
	var cntMap = make(map[int]int)
	var res [][]int
	for _, value := range nums {
		if _, ok := cntMap[value]; ok {
			cntMap[value]++
		} else {
			cntMap[value] = 1
		}
	}
	res = append(res, []int{})
	// 拿到数字和次数
	for k, v := range cntMap {
		// 从尾到头遍历结果 ,看上面逻辑
		for i := len(res) - 1; i >= 0; i-- {
			// 深拷贝
			next := append([]int(nil), res[i]...)
			// 用尽次数
			for j := 0; j < v; j++ {
				// 加入数字
				next = append(next, k)
				res = append(res, next)
			}
		}
	}
	return res
}
