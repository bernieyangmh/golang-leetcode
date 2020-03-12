package main

//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
//示例:
//输入: n = 4, k = 2
//输出:
//[
//[2,4],
//[3,4],
//[2,3],
//[1,2],
//[1,3],
//[1,4],
//]

func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return [][]int{}
	}
	res := [][]int{}
	dfs(1, k, n, []int{}, &res)
	return res
}

func dfs(start, k, n int, pre []int, res *[][]int) {
	if len(pre) == k {
		*res = append(*res, append([]int{}, pre...))
		return
	}
	//(k - len(pre)) 空的位置； n-(k-len(pre))+1之后的数字不需要再放进去
	//1234567 n=7, k = 6, len(pre)=3
	//[2,3,4,0,0,0] =>  7-6+3+1=5   5可以放进去，但6，7就不行了
	for i := start; i <= n-(k-len(pre))+1; i++ {
		dfs(i+1, k, n, append(pre, i), res)
	}

}
