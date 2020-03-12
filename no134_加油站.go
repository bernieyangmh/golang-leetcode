package main

//在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
//你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。
//你从其中的一个加油站出发，开始时油箱为空。
//如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。
//说明: 
//如果题目有解，该答案即为唯一答案。
//输入数组均为非空数组，且长度相同。
//输入数组中的元素均为非负数。

//找到最低点，最低点右边一定上升且如果能走一圈就一定走到
func canCompleteCircuit(gas []int, cost []int) int {
	length := len(gas)
	spare := 0
	minSpare := 1<<31 - 1
	minIndex := 0
	for i := 0; i < length; i++ {
		spare += gas[i] - cost[i]
		if spare < minSpare {
			minSpare = spare
			minIndex = i
		}
	}

	if spare < 0 {
		return -1
	}
	return (minIndex + 1) % length
}

// 如果总数<0，一定不行
//如果遇到走不过去的加油站，从下一个加油站开始
// 因为题目答案唯一，所以能走到，满足的加油站一定是
func canCompleteCircuit2(gas []int, cost []int) int {
	length := len(gas)
	total, cur, start := 0, 0, 0
	for i := 0; i < length; i++ {
		total += gas[i] - cost[i]
		cur += gas[i] - cost[i]
		if cur < 0 {
			start = i + 1
			cur = 0
		}
	}
	if total >= 0 {
		return start
	}
	return -1
}
