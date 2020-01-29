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
	println(divide1(-2147483648, -1))

}

// 常规做法是，减数一次一次减去被减数，不断更新差，直到差小于0，我们减了多少次，结果就是多少。
func divide1(dividend int, divisor int) int {
	var (
		res      int
		maxValue = (1 << 31) - 1
		minValue = -1 << 31
	)
	// 这里也需要处理最大最小值
	if divisor == 1 {
		return min(maxValue, max(dividend, minValue))
	}
	if divisor == -1 {
		return min(maxValue, max(-dividend, minValue))
	}
	sign := (dividend > 0) == (divisor > 0)
	if divisor < 0 {
		divisor = -divisor
	}
	if dividend < 0 {
		dividend = -dividend
	}
	//注意边界条件, 能再减一次的时候,res才能够+1
	for dividend >= divisor {
		res += 1
		dividend = dividend - divisor
	}
	if !sign {
		res = -res
	}
	return min(maxValue, max(res, minValue))
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

// 移位法
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

// 二分法
func divide3(dividend int, divisor int) int {
	var (
		MAX_VALUE = 1<<31 - 1
		MIN_VALUE = -1 << 31
		result    int
		sign      bool
	)
	if divisor == 0 {
		return dividend
	}
	// 1和-1特殊处理 优化一下
	if divisor == 1 {
		result = dividend
		return min(MAX_VALUE, max(MIN_VALUE, result))
	}
	if divisor == -1 {
		result = -dividend
		return min(MAX_VALUE, max(MIN_VALUE, result))
	}
	// 确定最后的正负
	sign = (dividend > 0) == (divisor > 0)

	// 换成绝对值
	if divisor < 0 {
		divisor = -divisor
	}
	if dividend < 0 {
		dividend = -dividend
	}
	// divisor 在multiDivide翻倍,在该while循环仍保持初始值
	for divisor <= dividend {
		println(result, dividend, divisor)
		r, d := multiDivide(dividend, divisor)
		dividend = d
		result += r
	}
	if !sign {
		result = -result
	}
	return min(MAX_VALUE, max(MIN_VALUE, result))
}

// 翻倍除法，如果可以被除，则下一步除数翻倍，直至除数大于被除数，
// 返回商加总的结果与被除数的剩余值；
// 10/3 1.减1次3 2.减2次3 3.减3次3 4.减4次3
func multiDivide(dividend int, divisor int) (res, d int) {
	count := 1

	// divisor 3, 6, 12, 24
	// count 1, 2, 4 ,8
	// res 1, 3, 7, 15
	for divisor <= dividend {
		dividend -= divisor
		res += count
		count += count
		divisor += divisor
	}
	// 抛出减后的被除数的剩余,从初始除数继续计数
	return res, dividend
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
