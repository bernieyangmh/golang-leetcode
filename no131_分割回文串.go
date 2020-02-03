// 给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
// 返回 s 所有可能的分割方案。
// 示例:
// 输入: "aab"
// 输出:
// [
//   ["aa","b"],
//   ["a","a","b"]
// ]

package main

import "fmt"

func main() {
	fmt.Println(partition("acaab"))

}

// 1、每一个结点表示剩余没有扫描到的字符串，产生分支是截取了剩余字符串的前缀；
// 2、产生前缀字符串的时候，判断前缀字符串是否是回文。
// 如果前缀字符串是回文，则可以产生分支和结点；
// 如果前缀字符串不是回文，则不产生分支和结点，这一步是剪枝操作。
// 3、在叶子结点是空字符串的时候结算，此时从根结点到叶子结点的路径，就是结果集里的一个结果，使用深度优先遍历，记录下所有可能的结果。
// 采用一个路径变量 path 搜索，path 全局使用一个（注意结算的时候，需要生成一个拷贝），因此在递归执行方法结束以后需要回溯，即将递归之前添加进来的元素拿出去；
// path 的操作只在列表的末端，因此合适的数据结构是栈。
// 作者：liweiwei1419
// 链接：https://leetcode-cn.com/problems/palindrome-partitioning/solution/hui-su-you-hua-jia-liao-dong-tai-gui-hua-by-liweiw/

// eg:acaab
// "" -> [a(ok), ac(X), aca(ok), acaa(X), acaab(X)]
// a->[c(ok), ca(X), caa(X), caab(X)]; aca->[a(ok), ab(X)]
// a->c->[a(ok), aa(ok), aab(X)]; aca->a->[b(ok)到底]
// a->c->a->[a(ok), ab(X)]; a->c->a->aa->[b(ok)到底]
// a->c->a->a->[b(ok)到底]
func partition(s string) [][]string {
	var res [][]string
	if len(s) == 0 {
		return res
	}
	helper(s, 0, len(s), []string{}, &res)
	return res
}

func helper(str string, start, l int, path []string, res *[][]string) {
	// 走完字符串,深拷贝一个,避免之后path的改动影响当前的结果
	fmt.Println(start, path)
	if start == l {
		tmp := append([]string(nil), path...)
		*res = append(*res, tmp)
	}
	// 从开始走到结束
	for i := start; i < l; i++ {

		// a + [c,ca,caa,caab] ac不是回文,后面的ca,caa,caab不需再走
		// 前缀不是回文,后面就不用深入了
		if !isPalindrome(str[start : i+1]) {
			continue
		}
		path = append(path, str[start:i+1])
		helper(str, i+1, l, path, res)
		// 回退到start的时候path的状态也回溯
		path = path[:len(path)-1]
	}
}

func isPalindrome(str string) bool {
	s := 0
	e := len(str) - 1
	for s < e {
		if str[s] != str[e] {
			return false
		}
		s++
		e--
	}
	return true
}
