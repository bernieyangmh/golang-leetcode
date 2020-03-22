package main

import (
	"fmt"
)

//给定一个正整数，返回它在 Excel 表中相对应的列名称。
//例如，
//1 -> A
//2 -> B
//3 -> C
//...
//26 -> Z
//27 -> AA
//28 -> AB

//示例 1:
//输入: 1
//输出: "A"
//示例 2:
//输入: 28
//输出: "AB"
//示例 3:
//输入: 701
//输出: "ZY"

func main() {
	fmt.Println(convertToTitle(701))

}

func convertToTitle(n int) string {
	res := ""
	r := 0
	for n > 0 {
		r = n % 26
		n = n / 26
		if r == 0 {
			n -= 1
			r = 26
		}
		res = string(byte(r+64)) + res
	}
	return res
}
