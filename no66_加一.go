package main

//给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
//最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
//你可以假设除了整数 0 之外，这个整数不会以零开头。
//示例 1:
//输入: [1,2,3]
//输出: [1,2,4]
//解释: 输入数组表示数字 123。
//示例 2:
//输入: [4,3,2,1]
//输出: [4,3,2,2]
//解释: 输入数组表示数字 4321。


//如果是9，当前设为0继续判断，如果不是+1 返回
//如果第一个是0，说明还有进位，在首位插入1
func plusOne(digits []int) []int {
	for i := len(digits)-1;i>=0;i--{
		if digits[i] != 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}

	return digits
}