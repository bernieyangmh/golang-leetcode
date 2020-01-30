package main

// 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的每个数字在每个组合中只能使用一次。
// 说明：
// 所有数字（包括目标数）都是正整数。
// 解集不能包含重复的组合。
// 示例 1:
// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
// 所求解集为:
// [
//   [1, 7],
//   [1, 2, 5],
//   [2, 6],
//   [1, 1, 6]
// ]
// 示例 2:
// 输入: candidates = [2,5,2,1,2], target = 5,
// 所求解集为:
// [
//   [1,2,2],
//   [5]
// ]

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{10, 1, 1, 2, 7, 6, 1, 5}
	fmt.Println(combinationSum2(a, 8))

}

func combinationSum2(candidates []int, target int) [][]int {
	var (
		res  [][]int
		path []int
	)
	sort.Ints(candidates)
	dfs(candidates, 0, target, path, &res)
	return res
}

func dfs(candidates []int, begin, target int, path []int, res *[][]int) {
	// 符合条件的数组,放入res
	if target == 0 {
		// path浅拷贝, 深拷贝重新生成tmp, 避免后面的回溯影响到最终的结果
		// 深拷贝
		tmp := append([]int(nil), path...) 
		*res = append(*res, tmp)

	}
	// 从当前指针开始
	for i := begin; i < len(candidates); i++ {
		// 同一层, 不应出现和当前字符相同的字符, 取第一个
		// [0,0,1,1,1,...] 同一层只要index=1 的0和index=2 的1
		if begin < i && candidates[i-1] == candidates[i] {
			continue
		}
		// 计算差值
		residue := target - candidates[i]
		if residue < 0 {
			break
		}

		path = append(path, candidates[i])
		// i+1 跳过当前字符,避免2+2+2+2 = 8 的重复字符情况
		dfs(candidates, i+1, residue, path, res)
		// 每一次递归前的“状态”应该是一样的 即 arr+[i]与 arr+[i+1] arr应该是一致的
		// 去掉最后一个, 回溯,恢复到append之前的状态
		if len(path) > 0 {
			path = path[0 : len(path)-1]
		}
	}
}
