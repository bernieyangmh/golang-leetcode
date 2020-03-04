package main


//给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。
//找到所有在 [1, n] 范围之间没有出现在数组中的数字。
//您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。
//示例:
//输入:
//[4,3,2,7,8,2,3,1]
//输出:
//[5,6]

// 1.常见解法，hash保存出现的数字，然后遍历


// 遍历，数字的值相当于索引，将该索引的数值置为负，然后遍历一遍，正的数字的索引就是未出现的
// focus 拿索引要取正，置负数要判断，结果索引而不是数值
func findDisappearedNumbers(nums []int) []int {
	for i:=0;i<len(nums); i++{
		index := 0
		if  nums[i] > 0 {
			index = nums[i]
		} else {
			index = -nums[i]
		}
		if nums[index-1] > 0 {
			nums[index-1] = -nums[index-1]
		}
	}

	res := []int{}
	for index, v := range nums {
		if v > 0 {
			res = append(res, index+1)
		}
	}
	return res

}