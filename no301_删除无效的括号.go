package main

//删除最小数量的无效括号，使得输入的字符串有效，返回所有可能的结果。
//说明: 输入可能包含了除 ( 和 ) 以外的字符。
//示例 1:
//输入: "()())()"
//输出: ["()()()", "(())()"]
//示例 2:
//输入: "(a)())()"
//输出: ["(a)()()", "(a())()"]
//示例 3:
//输入: ")("
//输出: [""]

var res []string

var size int

func removeInvalidParentheses(s string) []string {
	left := 0
	right := 0
	size = len(s)
	for _, l := range s {
		if string(l) == "(" {
			left++
		}
		if string(l) == ")" {
			if left>0 {
				left--
			}else {
				right++
			}
		}
	}
	dfs(s, 0, left, right)
	return res
}

func dfs(s string, st, l, r int)  {
	if l ==0 && r==0 {
		if check(s) {
			res = append(res, s)
		}
		return
	}
	for i:=st; i<len(s); i++{
		if i-1 >= st && s[i] == s[i-1] {
			continue
		}
		if l >0 && string(s[i]) == "(" {
			dfs(s[0:i]+s[i+1:], i, l-1, r)
		}
		if r > 0 && string(s[i]) == ")" {
			dfs(s[0:i]+s[i+1:], i, l, r-1)
		}
	}

}

func check(s string) bool {
	cnt := 0
	for _, l := range s{
		if string(l) == "(" {
			cnt++
		}
		if string(l)== ")"{
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return  cnt == 0
}