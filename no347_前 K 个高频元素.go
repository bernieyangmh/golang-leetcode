package main

import (
	"sort"
)

// 计算次数，次数不可能大于nums的长度，以nums.length+1建立一个数组当桶排序，索引就是次数
// 统计之后，每个次数的数放进二维数组里，从大到小放进结果，取前k个
func topKFrequent(nums []int, k int) []int {
	cntMap := make(map[int]int)
	rankList := [][]int{}
	res := []int{}
	// 多一个，索引和长度不一致
	for i:=0; i<=len(nums);i++{
		rankList = append(rankList, []int{})
	}

	for i:=0;i<len(nums);i++{
		if _, ok := cntMap[nums[i]];ok{
			cntMap[nums[i]]++
		} else {
			cntMap[nums[i]] = 1
		}
	}

	for num, cnt:= range cntMap {
		rankList[cnt] = append(rankList[cnt], num)
	}

	for i:=len(rankList)-1;i>=0;i--{
		if len(rankList[i]) > 0 {
			res = append(res, rankList[i]...)
		}
		if len(res) >=k {
			return res[:k]
		}
	}
	return res
}


func topKFrequent2(nums []int, k int) []int {
	cntMap := make(map[int]int)
	cntList := []int{}
	res := []int{}

	for i:=0;i<len(nums);i++{
		if _, ok := cntMap[nums[i]];ok{
			cntMap[nums[i]]++
		} else {
			cntMap[nums[i]] = 1
		}
	}

	for _, cnt:= range cntMap {
		cntList = append(cntList, cnt)
	}
	sort.Ints(cntList)
	minCnt := cntList[len(cntList)-k]

	for num, cnt := range cntMap {
		if cnt >= minCnt{
			res = append(res, num)
		}
	}
	return res
}