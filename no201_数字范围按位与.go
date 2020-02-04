// 给定范围 [m, n]，其中 0 <= m <= n <= 2147483647，返回此范围内所有数字的按位与（包含 m, n 两端点）。
// 示例 1:
// 输入: [5,7]
// 输出: 4
// 示例 2:
// 输入: [0,1]
// 输出: 0

package main

import "fmt"

func main() {
	fmt.Println(rangeBitwiseAnd3(13, 15))

}

func rangeBitwiseAnd(m int, n int) int {

	if m == n {
		return m
	}

	tmp1 := m & n

	mbits := getbit(m)
	nbits := getbit(n)
	if mbits == nbits {
		return tmp1
	}
	for i := 0; i < mbits; i++ {
		n = n >> 1
	}

	tmp2 := 1
	for i := 0; i < nbits-mbits; i++ {
		tmp2 = tmp2 << 1
	}
	res := 1
	for mbits != 0 {
		res = res << 1
		mbits--
	}
	return res + tmp1

}

func getbit(n int) int {
	res := 0
	for n != 0 {
		res++
		n = n >> 1
	}
	return res
}

// 5 101    10    1
//   110 -> 11 -> ok
// 7 111	11	  1

func rangeBitwiseAnd2(m int, n int) int {
	var cnt uint
	for m != n {
		// m比n位少,一定是0
		if m == 0 {
			return 0
		}
		m = m >> 1
		n = n >> 1
		cnt++
	}
	n = n << cnt
	return n
}

// 111 & 110 -> 110
// 110 & 101 -> 100

// eg 101101, 1101
// 101101 & 101100 -> 101100
// 101100 & 101011 -> 101000
// 101000 & 100111 -> 100000
// 100000 & 011111 -> 0
// 首先要同位数,然后二进制从左到右相同的,其他是0
func rangeBitwiseAnd3(m int, n int) int {
	for n > m {
		n = n & (n - 1)
	}
	return n
}

func rangeBitwiseAnd4(m int, n int) int {
	if m == n {
		return m
	}
	i := m^n
	i = i | (i>>1)
	i = i | (i>>2)
	i = i | (i>>4)
	i = i | (i>>8)
	i = i | (i>>16)
	return m&^i
}


// public int rangeBitwiseAnd(int m, int n) {
//     if (m == n) {
//         return m;
//     }
//     int i = m ^ n;
//     i |= (i >>> 1);
//     i |= (i >>> 2);
//     i |= (i >>> 4);
//     i |= (i >>> 8);
//     i |= (i >>> 16);
//     return m & ~i;
// }

// 作者：windliang
// 链接：https://leetcode-cn.com/problems/bitwise-and-of-numbers-range/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by--41/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。