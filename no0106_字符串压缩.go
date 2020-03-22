package main

import (
	"fmt"
	"strconv"
	"strings"
)

//
//字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。
//比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变短，
//则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。
//示例1:
//输入："aabcccccaaa"
//输出："a2b1c5a3"
//示例2:
//输入："abbccd"
//输出："abbccd"
//解释："abbccd"压缩后为"a1b2c2d1"，比原字符串长度更长。

func  main()  {
	s:="bbbac"
	fmt.Println(compressString(s))

}

func compressString(S string) string {
	if len(S) < 3 {
		return S
	}
	pre := S[0]
	count := 1
	res := ""

	for i:=1;i<len(S)-1;i++{
		if S[i] != pre {
			res += string(pre)
			res += strconv.Itoa(count)
			count = 1
			pre = S[i]
		} else {
			count++
		}
	}
	if S[len(S)-1] == pre {
		count++
		res += string(pre)
		res+=strconv.Itoa(count)
	} else {
		res += string(pre)
		res += strconv.Itoa(count)
		res += string(S[len(S)-1])
		res+="1"
	}
	if len(res) >= len(S) {
		return S
	}
	return res
}

func compressString2(S string) string {
	if S == "" {
		return ""
	}

	var sb strings.Builder
	curr := S[0]
	currLen := 1

	for i := 1; i < len(S); i++ {
		if S[i] == curr {
			currLen++
		} else {
			sb.WriteByte(curr)
			sb.WriteString(strconv.Itoa(currLen))
			curr = S[i]
			currLen = 1
		}
	}

	sb.WriteByte(curr)
	sb.WriteString(strconv.Itoa(currLen))

	if sb.Len() >= len(S) {
		return S
	}

	return sb.String()
}
