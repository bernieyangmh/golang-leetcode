// 给定一个大小为 n 的数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
// 说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1)。
// 示例 1:
// 输入: [3,2,3]
// 输出: [3]
// 示例 2:
// 输入: [1,1,1,3,3,2,2,2]
// 输出: [1,2]

package main

// 两次遍历;第一次找出数量前2的候选者;第二次看候选者超没超过n/3
func majorityElement(nums []int) []int {
	var res []int
	if len(nums) == 0 {
		return res
	}
	// 初始化
	candidateA := nums[0]
	candidateB := nums[0]
	countA := 0
	countB := 0
	for i := 0; i < len(nums); i++ {
		// A获得1票
		if nums[i] == candidateA {
			countA++
			continue
		}
		// B获得1票
		if nums[i] == candidateB {
			countB++
			continue
		}

		// 值和AB都不同, 如果候选者票数为0,换新的候选者
		// 因为优先A, 所以111遇到223的时候,A更新成3,B仍为1
		if countA == 0 {
			candidateA = nums[i]
			countA++
			continue
		}
		if countB == 0 {
			candidateB = nums[i]
			countB++
			continue
		}
		// 元素和AB不相同,票数不为0, AB均减1
		countA--
		countB--
	}
	countA = 0
	countB = 0
	// 找出AB的数量
	for i := 0; i < len(nums); i++ {
		if nums[i] == candidateA {
			countA++
		} else if nums[i] == candidateB {
			countB++
		}
	}
	if countA > len(nums)/3 {
		res = append(res, candidateA)
	}
	if countB > len(nums)/3 {
		res = append(res, candidateB)
	}
	return res
}
