// 给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
// 返回被除数 dividend 除以除数 divisor 得到的商。
// 示例 1:
// 输入: dividend = 10, divisor = 3
// 输出: 3
// 示例 2:
// 输入: dividend = 7, divisor = -3
// 输出: -2
// 说明:
// 被除数和除数均为 32 位有符号整数。
// 除数不为 0。
// 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,  2^31 − 1]。本题中，如果除法结果溢出，则返回 2^31 − 1。
package main

func main() {
	println(divide2(312312, 3))

}

func divide(dividend int, divisor int) int {
	// 正false 负true
	sign := (dividend > 0) != (divisor > 0)
	var (
		result    int
		MAX_VALUE = 1<<31 - 1
		MIN_VALUE = -1 << 31
	)
	// 全部转化为负数
	if dividend > 0 {
		dividend = -dividend
	}
	if divisor > 0 {
		divisor = -divisor
	}
	for dividend <= divisor {
		var (
			temp_result  = -1
			temp_divisor = divisor
		)
		for dividend <= (temp_divisor << 1) {
			if temp_divisor <= (MIN_VALUE >> 1) {
				break
			}
			temp_result = temp_result << 1
			temp_divisor = temp_divisor << 1
		}
		dividend = dividend - temp_divisor
		result += temp_result
		println(dividend, temp_divisor, result)
	}

	// 如果结果应是正数
	if !sign {
		if result <= MIN_VALUE {
			return MAX_VALUE
		}
		result = -result
	}
	return result
}

func divide2(dividend int, divisor int) int {
	sign := (dividend > 0) != (divisor > 0)
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	var count, result int
	// 除数不断左移，直到它大于被除数, count就是要进行的n次操作,也是商的位数
	// 二进制 101101/10 = 10110 count=5
	for dividend >= divisor {
		count += 1
		divisor = divisor << 1
	}
	println("count divisor dividend result")
	// 注意次数和边界,保证执行count次
	for count > 0 {
		print(count, " ", divisor, " ", dividend, " ", result)
		count -= 1
		// 除数右移减少, 除100000, 除10000, 除1000...
		divisor = divisor >> 1
		// 除数比被除数还大,跳过等除数再移一位
		if divisor <= dividend {
			result += (1 << uint32(count))
			dividend -= divisor
		}
		println(" | ", count, divisor, dividend, result)
	}
	if sign {
		result = -result
	}
	if -(1<<31) <= result && result <= (1<<31)-1 {
		return result
	}
	return 1<<31 - 1
}
