package main

import (
	"bytes"
	"fmt"
	"strconv"
)

//给定两个整数，分别表示分数的分子 numerator 和分母 denominator，以字符串形式返回小数。
//如果小数部分为循环小数，则将循环的部分括在括号内。
//示例 1:
//输入: numerator = 1, denominator = 2
//输出: "0.5"
//示例 2:
//输入: numerator = 2, denominator = 1
//输出: "2"
//示例 3:
//输入: numerator = 2, denominator = 3
//输出: "0.(6)"

func main() {
	println(fractionToDecimal(1, 8))

}

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	buffer := bytes.Buffer{}
	if (numerator < 0 && denominator > 0) || (numerator > 0 && denominator < 0) {
		buffer.WriteString("-")
	}
	num := abs(numerator)
	den := abs(denominator)

	buffer.WriteString(strconv.Itoa(num / den))
	num = num % den
	if num == 0 {
		return buffer.String()
	}
	buffer.WriteString(".")
	// 记录不同的余数，余数如果再出现，就说明是循环，int记录位置
	m := make(map[int]int, 10)
	repeatPos := -1
	for {
		num *= 10
		pos, ok := m[num]
		if !ok {
			m[num] = buffer.Len()
			//fmt.Println(num, m[num], buffer.String())
		} else {
			repeatPos = pos
			break
		}
		// 写结果，赋值余数
		buffer.WriteString(strconv.Itoa(num / den))
		//fmt.Println(buffer.String(), num, num%den, m)
		num %= den
		//能被除尽
		if num == 0 {
			break
		}
	}
	if repeatPos == -1 {
		return buffer.String()
	}
	res := buffer.String()
	//fmt.Println(res, repeatPos)
	return fmt.Sprintf("%s(%s)", res[0:repeatPos], res[repeatPos:])
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a

}
