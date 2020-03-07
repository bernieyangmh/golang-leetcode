package main

//实现 pow(x, n) ，即计算 x 的 n 次幂函数。
//示例 1:
//输入: 2.00000, 10
//输出: 1024.00000
//示例 2:
//输入: 2.10000, 3
//输出: 9.26100
//示例 3:
//输入: 2.00000, -2
//输出: 0.25000
//解释: 2-2 = 1/22 = 1/4 = 0.25
//

// 暴力
func myPow1(x float64, n int) float64 {
	N := n
	// 如果是负数
	if N < 0 {
		x = 1/ x
		N = -N
	}
	res := 1.0
	for i :=0;i<N;i++{
		res = res * x
	}
	return res
}


//二分法
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n == -1 {
		return 1 / x
	}
	half := myPow(x, n/2)
	rest := myPow(x, n%2)
	//n = 2*half+rest
	return rest * half * half
}

func myPow3(x float64, n int) float64 {
	N := n
	if N < 0 {
		x = 1/ x
		N = -N
	}
	res := 1.0
	cur_mutil := x
	//递归，每次除以2
	for i:=N;i>0;i/=2{
		//剩1
		if i%2 == 1 {
			res = res * cur_mutil
		}
		cur_mutil = cur_mutil * cur_mutil
	}
	return res
}



func myPow2(x float64, n int) float64 {
	if n < 0 {
		x = 1/ x
		n = -n
	}
	res := 1.0
	for n>0 {
		if n &1 == 1 {
			res *= x
		}
		x*=x
		n>>=1
	}
	return res
}