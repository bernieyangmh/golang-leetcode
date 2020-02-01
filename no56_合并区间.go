// 给出一个区间的集合，请合并所有重叠的区间。
// 示例 1:
// 输入: [[1,3],[2,6],[8,10],[15,18]]
// 输出: [[1,6],[8,10],[15,18]]
// 解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
// 示例 2:
// 输入: [[1,4],[4,5]]
// 输出: [[1,5]]
// 解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
package main

import (
	"fmt"
	"sort"
)

func main() {
	a := [][]int{}
	a = append(a, []int{1, 4})
	a = append(a, []int{0, 4})
	// a = append(a, []int{3, 5})
	// a = append(a, []int{15, 18})
	fmt.Println(merge2(a))

}

func merge(intervals [][]int) [][]int {
	var res [][]int
	sortDouleDimensionArr(intervals)
	size := len(intervals)
	i := 0
	// 遍历每一个数组
	for i < size {
		l := intervals[i][0]
		r := intervals[i][1]
		// 走到底并且找到 左边比当前右指针大的数停止
		for i < size-1 && intervals[i+1][0] <= r {
			i += 1
			r = max(intervals[i][1], r)
		}
		res = append(res, []int{l, r})
		i += 1
	}
	return res
}

//先排序
// 如果当前数组的头大于结果集最后一个数组的尾,说明是独立的间隔,直接放进结果集
// 如果当前数组的头小于等于结果集最后一个数组的尾,
// 说明有重叠的情况,结果集的尾取决于当前数组的尾和结果集尾的最大值
func merge2(intervals [][]int) [][]int {
	var res [][]int
	if len(intervals) == 0 {
		return res
	}
	sortDouleDimensionArr(intervals)
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if res[len(res)-1][1] < intervals[i][0] {
			res = append(res, intervals[i])
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		}
	}
	return res
}

// 写二维数组的排序算法太慢了
// 用一个hash表保存一个数组的头尾, 一维数组用来排序
func merge3(intervals [][]int) [][]int {
	l := len(intervals)
	h := map[int]int{}
	r := [][]int{}
	starts := []int{}
	if len(intervals) == 0 {
		return r
	}
	// 遍历每个数组,如果左端点在hash表里且右端点比存储的还小,跳过
	// 如果左端点不在hash表里,放进数组,方便后续排序.
	// 更新左端点的右端点的值,该值始终保持最大
	// [0,4],[1,5],[2,3] 这种要排序之后处理
	for i := 0; i < l; i++ {
		a, b := intervals[i][0], intervals[i][1]
		if v, ok := h[a]; ok {
			if b <= v {
				continue
			}
		} else {
			starts = append(starts, a)
		}
		h[a] = b
	}
	sort.Ints(starts)

	// 空结果之前处理过了,这里先放进第一个,避免之后每次循环都要判断
	r = append(r, []int{starts[0], h[starts[0]]})
	for i := 1; i < len(starts); i++ {
		a, b := starts[i], h[starts[i]]
		last := r[len(r)-1]
		// i= [2, 5], last = [1, 3]
		// 结果集的尾取决于当前数组的尾和结果集尾的最大值
		// 左端点在结果集最后一个中,更新结果集尾
		if last[0] <= a && a <= last[1] {
			r[len(r)-1][1] = max(b, last[1])
			// 不在,将数组放进结果集
		} else {
			r = append(r, []int{a, b})
		}
	}

	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 根据数组的第一个数 排序二维数组
func sortDouleDimensionArr(arr [][]int) {

	if len(arr) <= 1 {
		return
	}
	mid := arr[0]
	head := 0
	tail := len(arr) - 1
	for i := 1; i <= tail; {
		fmt.Println(i)
		if arr[i][0] > mid[0] {
			arr[i], arr[tail] = arr[tail], arr[i]
			tail--
		} else {
			arr[i], arr[head] = arr[head], arr[i]
			head++
			i++
		}
	}
	arr[head] = mid
	sortDouleDimensionArr(arr[:head])
	sortDouleDimensionArr(arr[head+1:])

}
