// 给定一个整数 (32 位有符号整数)，请编写一个函数来判断它是否是 4 的幂次方。
// 示例 1:
// 输入: 16
// 输出: true
// 示例 2:
// 输入: 5
// 输出: false
// 进阶：
// 你能不使用循环或者递归来完成本题吗？

package main

func main() {
	isPowerOfFour(8)

}

// 不停除以 4 直到不能 整 除，然后判断是否为 1 即可
// 注意是整除，所以判断余数
func isPowerOfFour(num int) bool {
	for num%4 == 0 && num != 0 {
		num = num / 4
		println(num)
	}
	return num == 1
}

//  如果一个数字 n 是 2 的幂次方，那么 n & (n - 1) 一定等于 0
// 1431655765 -> 01010101010101010101010101010101 奇数是0
// 4的米
func isPowerOfFour2(num int) bool {
	if num == 1 {
		return true
	}
	if num < 4 {
		return false
	}

	if num&(num-1) != 0 {
		return false
	}
	return num&1431655765 == num
}

// 是 2 的幂次方；减去 1 是三的倍数
func isPowerOfFour3(num int) bool {
	var uniqueMap = make(map[int]int)

	return num > 0 && (num&(num-1)) == 0 && (num-1)%3 == 0
}
