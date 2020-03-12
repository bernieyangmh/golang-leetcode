package main


//格雷编码是一个二进制数字系统，在该系统中，两个连续的数值仅有一个位数的差异。
//给定一个代表编码总位数的非负整数 n，打印其格雷编码序列。格雷编码序列必须以 0 开头。
//示例 1:
//输入: 2
//输出: [0,1,3,2]
//解释:
//00 - 0
//01 - 1
//11 - 3
//10 - 2
//对于给定的 n，其格雷编码序列并不唯一。
//例如，[0,2,3,1] 也是一个有效的格雷编码序列。
//00 - 0
//10 - 2
//11 - 3
//01 - 1

//https://leetcode-cn.com/problems/gray-code/solution/gray-code-jing-xiang-fan-she-fa-by-jyd/

// Gn    G'n  Rn
// 00 -> 000, 100
// 01 -> 001, 101
// 11 -> 011, 111
// 10 -> 010, 110
func grayCode(n int) []int {
	res := []int{0}
	head := 1
	for i:=0;i<n;i++{
		// 倒序遍历，因为 010和110才只差一位，要让他俩连一起才满足 "两个连续的数值仅有一个位数的差异"
		for j := len(res)-1;j>=0;j-- {
			res = append(res, head + res[j])
		}
		//0		00		000
		//1		10		100
		head <<= 1
	}
	return res
}