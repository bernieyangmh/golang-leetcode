package main

import (
	"sort"
)

//给出方程式 A / B = k, 其中 A 和 B 均为代表字符串的变量， k 是一个浮点型数字。根据已知方程式求解问题，并返回计算结果。如果结果不存在，则返回 -1.0。
//示例 :
//给定 a / b = 2.0, b / c = 3.0
//问题: a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ? 
//返回 [6.0, 0.5, -1.0, 1.0, -1.0 ]
//输入为: vector<pair<string, string>> equations, vector<double>& values, vector<pair<string, string>> queries(方程式，方程式结果，问题方程式)， 其中 equations.size() == values.size()，即方程式的长度与方程式结果长度相等（程式与结果一一对应），并且结果值均为正数。以上为方程式的描述。 返回vector<double>类型。
//基于上述例子，输入如下：
//equations(方程式) = [ ["a", "b"], ["b", "c"] ],
//values(方程式结果) = [2.0, 3.0],
//queries(问题方程式) = [ ["a", "c"], ["b", "a"], ["a", "e"], ["a", "a"], ["x", "x"] ].
//输入总是有效的。你可以假设除法运算中不会出现除数为0的情况，且不存在任何矛盾的结果。

type Value struct {
	Status int
	Val    float64
}

//一个图，map构成的二维表格，表示x-y,y-x的比值
// 注意的是构建图的过程，要写的简单点
// 对于每一个x,y的边，深度递归
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)
	sort.Ints()
	for i := 0; i < len(equations); i++ {
		x := equations[i][0]
		y := equations[i][1]
		k := values[i]
		if graph[x] == nil {
			graph[x] = make(map[string]float64)
		}
		if graph[y] == nil {
			graph[y] = make(map[string]float64)
		}
		graph[x][y] = k
		graph[y][x] = 1.0 / k
	}
	res := make([]float64, len(queries))
	for i := 0; i < len(queries); i++ {
		x := queries[i][0]
		y := queries[i][1]
		var hasX, hasY bool
		_, hasX = graph[x]
		_, hasY = graph[y]
		if !hasX || !hasY {
			res[i] = -1.0
		} else {
			// 这里的visited是对每一对xy重新走一次
			res[i] = divide(x, y, make(map[string]bool), graph)
		}
	}
	return res

}

func divide(x, y string, visited map[string]bool, graph map[string]map[string]float64) float64 {
	if x == y {
		return 1.0
	}
	visited[x] = true
	// 核心 从x出发的地方，如果有两条路任走其一均可
	// 不存在路径，如x-e e-x 返回-1，不影响d
	for n, _ := range graph[x] {
		if visited[n] {
			continue
		}
		visited[n] = true
		// 核心， x-a-y x-y的值等于各个路径的乘积
		// x=2b b=1/2a a= 1/3c c=4y x=4/3y
		d := divide(n, y, visited, graph)
		if d > 0 {
			return d * graph[x][n]
		}
	}
	return -1.0
}


//并查集
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	res := make([]float64, len(queries))
	root := make(map[string]string)
	dist := make(map[string]float64)

	for i := 0; i < len(equations); i++ {
		r1 := find(root, dist, equations[i][0])
		r2 := find(root, dist, equations[i][1])
		root[r1] = r2
		dist[r1] = dist[equations[i][1]] * values[i] / dist[equations[i][0]]
	}

	for i := 0; i < len(queries); i++ {
		var hasX, hasY bool
		_, hasX = root[queries[i][0]]
		_, hasY = root[queries[i][1]]
		if !hasX || !hasY {
			res[i] = -1.0
			continue
		}
		r1 := find(root, dist, queries[i][0])
		r2 := find(root, dist, queries[i][1])
		if r1 != r2 {
			res[i] = -1.0
			continue
		}
		res[i] = dist[queries[i][0]] / dist[queries[i][1]]
	}
	return res
}

func find(root map[string]string, dist map[string]float64, s string) string {
	if _, ok := root[s]; !ok {
		root[s] = s
		dist[s] = 1.0
		return s
	}
	if root[s] == s {
		return s
	}
	lastP := root[s]
	P := find(root, dist, lastP)
	root[s] = P
	dist[s] = dist[s] * dist[lastP]
	return P
}
