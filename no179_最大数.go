package main

import (
	"sort"
	"strconv"
	"strings"
)

//给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。
//示例 1:
//输入: [10,2]
//输出: 210
//示例 2:
//输入: [3,30,34,5,9]
//输出: 9534330

//排序比较
func largestNumber(nums []int) string {
	ss := make([]string, len(nums))
	for i, num := range nums {
		ss[i] = strconv.Itoa(num)
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i]+ss[j] >= ss[j]+ss[i]
	})
	o := strings.Join(ss, "")
	if o[0] == '0' {
		return "0"
	}
	return o
}