package main

//根据每日 气温 列表，请重新生成一个列表，对应位置的输入是你需要再等待多久温度才会升高超过该日的天数。
//如果之后都不会升高，请在该位置用 0 来代替。
//例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。
//提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。

//从当前标记向左走，
// 遇到比当前数大的，break，因为肯定该数能够满足前面数的要求
// 遇到已经赋值的数字，跳过
// 剩下的数字，i-j就是等待的天数
func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))

	for i := 1; i < len(T); i++ {
		for j := i - 1; j >= 0; j-- {
			if T[j] >= T[i] {
				break
			}
			if res[j] != 0 {
				continue
			} else {
				if T[i] > T[j] {
					res[j] = i - j
				}
			}
		}
	}
	return res
}

func dailyTemperatures2(T []int) []int {

	res := make([]int, len(T))
	stack := []int{}
	for i := len(T) - 1; i >= 0; i-- {
		for len(stack) > 0 && T[i] >= T[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = 0
		} else {
			res[i] = stack[len(stack)] - i
		}
		stack = append(stack, i)
	}
	return res
}

//用栈保存
// 当前结果比栈顶小的入栈，遇到比栈顶大的元素，for循环一直抛出直到当前元素比栈顶小
func dailyTemperatures3(T []int) []int {
	res := make([]int, len(T))
	stack := make([]int, 0)
	for i := 0; i < len(T); i++ {
		for len(stack) > 0 && T[i] > T[stack[len(stack)-1]] {
			t := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[t] = i - t
		}
		stack = append(stack, i)
	}

	return res
}
