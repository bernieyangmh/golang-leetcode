package main

import (
	"fmt"
)

//找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
//说明：
//所有数字都是正整数。
//解集不能包含重复的组合。 
//示例 1:
//输入: k = 3, n = 7
//输出: [[1,2,4]]
//示例 2:
//输入: k = 3, n = 9
//输出: [[1,2,6], [1,3,5], [2,3,4]]

func main() {
	k := 8
	n := 36
	fmt.Println(combinationSum3(k, n))

}

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}

	// 寻找 n 的上限：[9, 8, ... , (9 - k + 1)]，它们的和为 (19 - k) * k / 2
	// 比上限还大，就不用搜索了：
	if k <= 0 || n <= 0 || k >= n || n > (19-k)*k/2 {
		return res
	}

	path := []int{}
	return dfs(1, n, k, path, res)

}

func dfs(c, residue, k int, tmp []int, res [][]int) [][]int {
	if 10-c < k {
		return res
	}
	if k == 0 {
		if residue == 0 {
			res = append(res, tmp)
			return res
		}
	}
	for i := c; i <= 10-k; i++ {
		if residue-i < 0 {
			break
		}
		cur := append([]int{}, tmp...)
		cur = append(cur, i)
		fmt.Println(i, cur, tmp, k-1, residue-i)
		res = dfs(i+1, residue-i, k-1, cur, res)
		//tmp = tmp[:len(tmp)-1]
	}
	return res
}
