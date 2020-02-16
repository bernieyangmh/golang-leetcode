// 给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
// 例如，给出 n = 3，生成结果为：
// [
//   "((()))",
//   "(()())",
//   "(())()",
//   "()(())",
//   "()()()"
// ]
package main

import "fmt"

func main() {
	generateParenthesis2(3)

}

func generateParenthesis(n int) []string {
	res := []string{}
	res = backtrack(res, "", 0, 0, n)
	return res
}

func backtrack(res []string, cur string, open, close, max int) []string {
	if len(cur) == max*2 {
		res = append(res, cur)
		return res
	}

	// 先放左,相当深度优先
	// (( 之后,先放(((, 回来之后再放(()
	if open < max {
		res = backtrack(res, cur+"(", open+1, close, max)
	}
	if close < open {
		res = backtrack(res, cur+")", open, close+1, max)
	}
	return res
}


// 作者：yuyu-13
// 链接：https://leetcode-cn.com/problems/generate-parentheses/solution/zui-jian-dan-yi-dong-de-dong-tai-gui-hua-bu-lun-da/
// func generateParenthesis2(n int) []string {
// 	if n == 0 {
// 		return []string{}
// 	}
// 	res := [][]string{}
// 	res = append(res, []string{})
// 	res = append(res, []string{"()"})

// 	for i := 2; i < n+1; i++ {
// 		l := []string{}
// 		for j := 0; j < i; j++ {
// 			n1 := res[j]
// 			n2 := res[i-j-1]
// 			fmt.Println(n1, n2)
// 			for _, k1 := range n1 {
// 				for _, k2 := range n2 {
// 					el := "(" + k1 + ")" + k2
// 					l = append(l, el)
// 				}
// 			}

// 		}
// 		// fmt.Println(l)
// 		res = append(res, l)
// 	}
// 	// fmt.Println(res)
// 	return res[n]
// }



func generateParenthesis3(n int) []string {
	res := []string{}
	if n == 0 {
		res = append(res, "")
	} else {
		for c:=0; c< n; c++ {
			for _,left := range generateParenthesis3(c) {
				for _, right := range generateParenthesis3(n-1-c){
					res = append(res, "(" + left + ")" + right)
				}
			}
		}
	}
	return res
}