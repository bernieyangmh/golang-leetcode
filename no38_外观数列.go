package main

import (
	"fmt"
	"strconv"
)

//「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。前五项如下：
//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//1 被读作  "one 1"  ("一个一") , 即 11。
//11 被读作 "two 1s" ("两个一"）, 即 21。
//21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。
//给定一个正整数 n（1 ≤ n ≤ 30），输出外观数列的第 n 项。
//注意：整数序列中的每一项将表示为一个字符串。

func main() {
	fmt.Println(countAndSay(4))

}

func countAndSay(n int) string {
	res := "1"
	if n == 1 {
		return res
	}
	// 总共n个,除去"1"第一个还剩n-1次
	for i := 1; i < n; i++ {
		//之前计数的字符
		curC := res[0]
		// 计数的数量
		curN := 0
		// 临时结果
		tmp := ""

		for i := 0; i < len(res); i++ {
			// 现在的数不是之前的字符, 把数量和字符都加上
			// curc换成现在的数，从1开始计数
			if res[i] != curC {
				tmp += strconv.Itoa(curN)
				tmp += string(curC)
				curC = res[i]
				curN = 1
			} else {
				curN++
			}
		}
		// 最后的数没有可以比较的，直接加上
		tmp += strconv.Itoa(curN)
		tmp += string(curC)
		res = tmp
	}
	return res
}

// 别人写的，更简洁
func countAndSay2(n int) string {
	res := "1"
	for i := 2; i <= n; i++ {
		tmp := ""
		pre := res[0]
		cnt := 1
		for j := 1; j < len(res); j++ {
			c := res[j]
			if c == pre {
				cnt++
			} else {
				tmp += strconv.Itoa(cnt)
				tmp += string(pre)
				pre = c
				cnt = 1
			}
		}
		tmp += strconv.Itoa(cnt)
		tmp += string(pre)
		res = tmp
	}
	return res
}
