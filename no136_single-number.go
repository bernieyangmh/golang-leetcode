// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
// 说明：
// 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
// 示例 1:
// 输入: [2,2,1]
// 输出: 1
// 示例 2:
// 输入: [4,1,2,1,2]
// 输出: 4

package main

func main() {
	n := []int{1, 2, 1, 2, 4}
	println(singleNumber(n))
}

// 异或，两次相同数字的异或结果仍为0
// 剩下的数字即为出现一次的数字
func singleNumber(nums []int) int {
	var singleNum int
	for _, i := range nums {
		singleNum ^= i
		println(singleNum)
	}
	return singleNum
}


