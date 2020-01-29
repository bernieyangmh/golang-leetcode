// 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的数字可以无限制重复被选取。
// 说明：
// 所有数字（包括 target）都是正整数。
// 解集不能包含重复的组合。
// 示例 1:
// 输入: candidates = [2,3,6,7], target = 7,
// 所求解集为:
// [
//   [7],
//   [2,2,3]
// ]
// 示例 2:
// 输入: candidates = [2,3,5], target = 8,
// 所求解集为:
// [
//   [2,2,2,2],
//   [2,3,3],
//   [3,5]
// ]
package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{2, 3, 5}
	// fmt.Println(a[len(a)-1])
	fmt.Println(combinationSum(a, 8))
}

// 回溯+根dsf遍历
// candidates = [2, 3, 6, 7]，target = 7
//  4个树,根结点是每个元素,每一层的叶子结点是每个元素,
//	2
//  2		3		6		7
//  2367    2367    2367    2367
//  xyxx    fxxx    ffxx    fffx
//  y代表符合条件的路径, x代表走到底也不符合的路径,f代表叶子结点的值大于父结点的值,不符合要求
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int

	if target == 0 || len(candidates) == 0 {
		return res
	}
	// 排序,利于是否判断重复
	// 1、前面用过的数后面不能再用；2、下一层边上的数不能小于上一层边上的数
	sort.Ints(candidates)
	var path []int
	// 传res的指针,不然res在函数内的改动不生效
	dfs(candidates, 0, target, path, &res)
	return res
}

func dfs(candidates []int, begin, target int, path []int, res *[][]int) {
	// 符合条件的数组,放入res
	if target == 0 {
		var tmp []int
		// path浅拷贝, 深拷贝重新生成tmp, 避免后面的回溯影响到最终的结果
		for _, v := range path {
			tmp = append(tmp, v)
		}
		*res = append(*res, tmp)
	}
	// 从当前指针开始
	for i := begin; i < len(candidates); i++ {
		// 计算差值
		residue := target - candidates[i]
		if residue < 0 {
			break
		}

		path = append(path, candidates[i])
		dfs(candidates, i, residue, path, res)
		// 每一次递归前的“状态”应该是一样的 即 arr+[i]与 arr+[i+1] arr应该是一致的
		// 去掉最后一个, 回溯,恢复到append之前的状态
		if len(path) > 0 {
			fmt.Println(path, i, residue)
			path = path[0 : len(path)-1]
		}
	}
}
