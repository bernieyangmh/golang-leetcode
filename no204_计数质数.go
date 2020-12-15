package main

//统计所有小于非负整数 n 的质数的数量。
//示例:
//输入: 10
//输出: 4
//解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。

func countPrimes(n int) int {
	primArry := constructPrimArray(n)
	count := 0
	for i := 2; i < n; i++ {
		if primArry[i] {
			count++
		}
	}
	return count
}

//构建质数数组
func constructPrimArray(n int) []bool {
	isPrim := []bool{}
	for i := 0; i < n; i++ {
		isPrim = append(isPrim, true)
	}

	for i := 2; i*i < n; i++ {
		if isPrim[i] {
			for j := i * i; j < n; j += i {
				isPrim[j] = false
			}
		}
	}
	return isPrim
}
