package main

import (
	"fmt"
	"strconv"
)

//给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
//示例:
//输入: "25525511135"
//输出: ["255.255.11.135", "255.255.111.35"]

func main() {
	fmt.Println(restoreIpAddresses("0000"))

}

func restoreIpAddresses(s string) []string {
	res := []string{}
	if len(s) < 4 {
		return res
	}
	res = dfs(0, 1, "", s, res)
	return res
}

// 当前要判断的数,当前的部分

//细节太多

func dfs(i, p int, cur, s string, res []string) []string {

	//1. 当前的位置超过字符串长度了，肯定不行
	if i > len(s)-1 {
		return res
	}

	// 2.如果最后剩余的数量超过3,不符合要求,就返回
	if p == 4 && len(s)-i > 3 {
		return res
	}

	// 3.最后部分，只有[0]或[开头不是0,且小于等于255]的数字才能拼接
	if p == 4 {
		remainder := s[i:]
		//去掉开头是0且不是'0'的
		if remainder[0] == '0' && len(remainder) > 1 {
			return res
		}
		num, _ := strconv.Atoi(remainder)
		if num <= 255 {
			res = append(res, cur+string(remainder))
		}
		return res
	}
	//4.保存，方便回溯
	tmp := cur

	// 5.可以+1，这里不需要加条件
	cur = tmp + string(s[i]) + "."
	res = dfs(i+1, p+1, cur, s, res)

	// 6.可以+2 后面有两个数字且开头不是0
	if i+2 < len(s) && s[i] != '0' {
		cur = tmp + string(s[i:i+2]) + "."
		res = dfs(i+2, p+1, cur, s, res)
	}
	// 7.可以+3，后面有三个数字且开头不是0且小于等于255
	if i+3 < len(s) && s[i] != '0' {
		num, _ := strconv.Atoi(s[i : i+3])
		if num <= 255 {
			cur = tmp + string(s[i:i+3]) + "."
			res = dfs(i+3, p+1, cur, s, res)
		}
	}
	return res
}


//别人写的
var res []string

func restoreIpAddresses2(s string) []string {
	res = []string{}
	length := len(s)
	backtrack(0,length, 4, "", s)
	return res
}

func backtrack(i, l, flag int, tmp,s string)  {
	// 4个都已拼接完成且用完所有字符
	if i == l && flag == 0{
		res = append(res, tmp[0:len(tmp)-1])
		return
	}
	// 还能小于0说明有多余的
	if flag < 0 {
		return
	}

	for j:=i;j<i+3;j++{
		//边界检查
		if j<l{
			//0只能单独构成一个部分
			if i==j && s[j] == '0' {
				backtrack(j+1, l, flag-1, tmp+string(s[j])+".", s)
				break
			}
		}
		// 边界检查
		if j+1 <= l {
			//1，11，111
			num, _ := strconv.Atoi(s[i : j+1])
			if num <= 255 {
				backtrack(j+1, l, flag-1, tmp+s[i:j+1]+".", s)
			}
		}
	}
}