package main

import (
	"fmt"
)

//现在你总共有 n 门课需要选，记为 0 到 n-1。
//在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]
//给定课程总量以及它们的先决条件，判断是否可能完成所有课程的学习？
//示例 1:
//输入: 2, [[1,0]]
//输出: true
//解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
//示例 2:
//输入: 2, [[1,0],[0,1]]
//输出: false
//解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。
//
//说明:
//输入的先决条件是由边缘列表表示的图形，而不是邻接矩阵。详情请参见图的表示法。
//你可以假定输入的先决条件中没有重复的边。
//
//提示:
//这个问题相当于查找一个循环是否存在于有向图中。如果存在循环，则不存在拓扑排序，因此不可能选取所有课程进行学习。
//通过 DFS 进行拓扑排序 - 一个关于Coursera的精彩视频教程（21分钟），介绍拓扑排序的基本概念。
//拓扑排序也可以通过 BFS 完成。

//链接：https://leetcode-cn.com/problems/course-schedule/solution/course-schedule-tuo-bu-pai-xu-bfsdfsliang-chong-fa/

func main() {
	pre := [][]int{}
	pre = append(pre, []int{2,1})
	pre = append(pre, []int{4,1})
	pre = append(pre, []int{4,2})
	pre = append(pre, []int{3,2})
	pre = append(pre, []int{5,4})
	pre = append(pre, []int{5,3})
	fmt.Println(pre)
	println(canFinish(6, pre))
}


func canFinish(numCourses int, prerequisites [][]int) bool {

	in_degrees := make([]int, numCourses)
	fmt.Println(in_degrees)
	queue := []int{}

	//每个课程需要前置课程的数量
	for _, cp := range prerequisites {
		in_degrees[cp[0]]++
	}

	// 所有不需要前置课程的课程放进队列， 有环的课程则都不在
	for i := 0; i < numCourses; i++ {
		if in_degrees[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) != 0 {
		pre := queue[0]
		queue = queue[1:]
		numCourses--
		for _, arr := range prerequisites {
			// 找到该课程为前置条件的条件
			if arr[1] != pre {
				continue
			}
			// 把arr[1]课程为前置课程的课程arr[0]所需课程数量减1
			// 如[1,3] 找到了1课程的一个前置课程3， 就把1需要de数量-1
			in_degrees[arr[0]]--

			// 如果0，说明arr[0]课程的前置条件我们都已找到
			if in_degrees[arr[0]] == 0 {
				queue = append(queue, arr[0])
			}
		}
	}
	//如果无环，最终每个数都进队列里一遍
	return numCourses == 0
}

func canFinish2(numCourses int, prerequisites [][]int) bool {
	adj := [][]int{}

	// 为每个课程构建一个邻接表
	for i := 0; i < numCourses; i++ {
		tmp := make([]int, numCourses)
		adj = append(adj, tmp)
	}

	// 1前置课程，0非前置课程，-1不在环里
	flags := make([]int, numCourses)

	for _, cp := range prerequisites {
		adj[cp[1]][cp[0]] = 1
	}
	for i := 0; i < numCourses; i++ {
		if !dfs(adj, flags, i) {
			return false
		}
	}
	return true

}


func dfs(adj [][]int, flags []int, i int) bool {
	// 因为递归，当我从一个前置课程出发，前面的课程应该要么是0（非前置课程）,要么是-1(已经判断过的)
	// 如果遇到的是1，说明是一条路走到底仍有前置课程，即首尾相连
	if flags[i] == 1 {
		return false
	}
	if flags[i] == -1 {
		return true
	}
	flags[i] = 1
	for j := 0; j < len(adj); j++ {
		// 是前置课程，且他在环里
		// 相反的是1.不是前置课程 2.是前置课程但不在环里
		if adj[i][j] == 1 && !dfs(adj, flags, j) {
			return false
		}
	}
	//
	flags[i] = -1
	return true
}
