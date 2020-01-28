// 给定两个字符串 s1 和 s2，
// 要求判定 s2 是否能被 s1 循环移位得到的字符串包含。
// 比如，给定 s1 = AABCD 和 s2 = CDAA，返回 true 。
// 给定 s1 = ABCD，s2 = ACBD， 则返回 false。

package main

import "strings"

func main() {
	println(isContain("123123", "234"))

}

func isContain(a, b string) bool {
	aArr := strings.Split(a, "")
	bArr := strings.Split(b, "")

	var (
		n = len(aArr)
		m = len(bArr)
	)

	for i := 0; i < n; i++ {
		var p1 = i
		var p2 int

		// 假设有两个a数组拼接在一起,p1指针可以在0到2n移动
		// p2在b数组上移动
		for (p1 < 2*n) && p2 < m {
			// p1 走到底再从头开始, 虚拟两个数组
			if aArr[p1%n] == bArr[p2] {
				p1 += 1
				p2 += 1
			} else {
				break
			}
		}
		// 有子字符串包含,则p2走完全程, 最后+1; 判断的时候不-1
		if p2 == m {
			return true
		}
	}
	return false
}
