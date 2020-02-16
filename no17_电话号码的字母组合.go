// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
// 示例:
// 输入："23"
// 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	intLetter := make(map[string][]string)
	intLetter["2"] = []string{"a", "b", "c"}
	intLetter["3"] = []string{"d", "e", "f"}
	intLetter["4"] = []string{"g", "h", "i"}
	intLetter["5"] = []string{"j", "k", "l"}
	intLetter["6"] = []string{"m", "n", "o"}
	intLetter["7"] = []string{"p", "q", "r", "s"}
	intLetter["8"] = []string{"t", "u", "v"}
	intLetter["9"] = []string{"w", "x", "y", "z"}

	res := []string{""}

	for _, s := range digits {
		// 每取一个数字,新建一个临时结果,将上次结果依次处理,再放回结果里
		var tmp []string
		for _, v := range intLetter[string(s)] {
			for _, originWord := range res {
				originWord = originWord + v
				tmp = append(tmp, originWord)
			}
		}
		res = tmp
	}
	return res
}

// 队列
// a,b,c -> b, c, ad, ae, af -> ad, ae ,af .....
func letterCombinations2(digits string) []string {
	if digits == "" {
		return []string{}
	}

	intLetter := make(map[string][]string)
	intLetter["2"] = []string{"a", "b", "c"}
	intLetter["3"] = []string{"d", "e", "f"}
	intLetter["4"] = []string{"g", "h", "i"}
	intLetter["5"] = []string{"j", "k", "l"}
	intLetter["6"] = []string{"m", "n", "o"}
	intLetter["7"] = []string{"p", "q", "r", "s"}
	intLetter["8"] = []string{"t", "u", "v"}
	intLetter["9"] = []string{"w", "x", "y", "z"}

	res := []string{""}

	for _, s := range digits {
		size := len(res)
		// 循环队列长度次,之后新添的不处理
		for j := 0; j < size; j++ {
			// res.remove(0)
			tmp := res[0]
			res = res[1:]
			for _, l := range intLetter[string(s)] {
				res = append(res, tmp+l)
			}
		}
	}

	return res
}
