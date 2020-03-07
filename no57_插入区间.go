package main

import (
	"fmt"
)

//给出一个无重叠的 ，按照区间起始端点排序的区间列表。
//在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
//示例 1:
//输入: intervals = [[1,3],[6,9]], newInterval = [2,5]
//输出: [[1,5],[6,9]]
//示例 2:
//输入: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
//输出: [[1,2],[3,10],[12,16]]
//解释: 这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。

func main()  {
	a := [][]int{}
	a = append(a, []int{1,3})
	a = append(a, []int{6,9})

	newIntervals := []int{2,5}
	fmt.Println(insert2(a, newIntervals))


}

//贪心算法
// 在范围之前的先加进结果里
// 范围内的合并
// 后面的再加进结果里
func insert(intervals [][]int, newInterval []int) [][]int {
	newStart, newEnd := newInterval[0], newInterval[1]
	index, n := 0, len(intervals)

	output := [][]int{}

	//先遍历一遍，把所有左边界的加进结果里
	for index < n && newStart > intervals[index][0] {
		output = append(output, intervals[index])
		index++
	}

	interval := make([]int, 2)
	if len(output) == 0 || output[len(output)-1][1] < newStart {
		output = append(output, newInterval)
	} else {
		interval = output[len(output)-1]
		output = output[:len(output)-1]
		interval[1] = max(interval[1], newEnd)
		output = append(output, interval)
	}

	for index < n {
		interval = intervals[index]
		index++
		start, end := interval[0], interval[1]
		if output[len(output)-1][1] < start {
			output = append(output, interval)
		} else {
			interval = output[len(output)-1]
			output = output[:len(output)-1]
			interval[1] = max(interval[1], end)
			output = append(output, interval)

		}
	}
	return output
}

// 清晰
func insert2(intervals [][]int, newInterval []int) [][]int {
	output := [][]int{}
	index, n := 0, len(intervals)
	left, right := newInterval[0], newInterval[1]

	// 小于左边界的直接放进去
	for index < n && left > intervals[index][1] {
		output = append(output, intervals[index])
		index++
	}
	tmp := []int{left, right}
	// 找重合区域,知道right比区间的左边界还小结束
	for index < n && right >= intervals[index][0] {
		//左的最左，右的最右
		tmp[0] = min(tmp[0], intervals[index][0])
		tmp[1] = max(tmp[1], intervals[index][1])
		index++
	}
	output = append(output, tmp)
	// 后面如果还有，放进去
	for index < n {
		output = append(output, intervals[index])
		index++
	}
	return output
}


// 想太多。。。
func insertuseless(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	// 无区间
	if len(intervals) == 0 || len(newInterval) == 0{
		return append(res, newInterval)
	}

	// 新区间包含整个旧区间
	if newInterval[0] <= intervals[0][0] && intervals[len(intervals)-1][1] <=  newInterval[1] {
		return append(res, newInterval)
	}

	// 新区间在区间列表的左边
	if newInterval[1] < intervals[0][0] {
		res = append(res, newInterval)
		return append(res, intervals...)
	}

	// 新区间在区间列表的右边
	if intervals[len(intervals)-1][1] < newInterval[0] {
		res = intervals
		return append(res, newInterval)
	}

	left := newInterval[0]
	right := newInterval[1]

	// 要被新区间替代的区间的最左和最右两个
	deleteArr := []int{}
	mergeInterval := []int{}
	for i:=0;i<len(intervals);i++{
		// 找到要合并的区间, 右边界比左大,左边界比右小, 因为边界接触也要合并,所以是等于
		//if (intervals[i][0] <= left && left <= intervals[i][1]) || (intervals[i][0] <= right && right <= intervals[i][1]){
			if intervals[i][1] > left  &&  intervals[i][0] <= right{
			// 确定要合并区间的左边界
			if len(mergeInterval) == 0 {
				mergeInterval = append(mergeInterval, min(intervals[i][0], left))
			}
			deleteArr = append(deleteArr, i)
		}
		// 确定右边界, 第一个左边界比右要大的前一个区间
		// [i-1] right [i]  [ right i-1] [i]
		if i > 0 && right < intervals[i][0] && len(mergeInterval) == 1 {
			mergeInterval = append(mergeInterval, max(intervals[i-1][1], right))
		}
	}

	// right比所有区间都大的情况,或只有一个区间
	if len(mergeInterval) == 1 {
		mergeInterval = append(mergeInterval, max(right, intervals[len(intervals)-1][1]))
	}
	fmt.Println(mergeInterval, deleteArr)
	part1 := intervals[:deleteArr[0]]
	part2 := []int{mergeInterval[0], mergeInterval[1]}
	part3 := intervals[deleteArr[len(deleteArr)-1]+1:]
	fmt.Println(part1, part2, part3)
	res = part1
	res = append(res, part2)
	res = append(res, part3...)

	return res
}


func max(a, b int) int {

	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
