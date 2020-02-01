package main

import "strconv"

import "fmt"

import "strings"

func main() {
	// println(factorial(4))
	println(getPermutation(3, 3))

}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	var res = 1
	var tmp = 1
	for tmp <= n {
		res = res * tmp
		tmp++
	}
	return res
}

//
func getPermutation(n int, k int) string {
	var used []bool
	var path []string
	var factorial []int
	if n == 0 {
		return ""
	}
	// 多一个flase,方便后续操作
	for i := 0; i <= n+1; i++ {
		used = append(used, false)
	}
	factorial = append(factorial, 1)
	for i := 1; i <= n; i++ {
		factorial = append(factorial, factorial[i-1]*i)
	}
	dfs(n, k, 0, &path, factorial, used)
	fmt.Println(path)

	return strings.Join(path, "")
}

func dfs(n, k, index int, path *[]string, factorial []int, used []bool) {
	fmt.Println(n, k, index+1, path, factorial, used)
	// 走到头
	if index == n {
		return
	}
	cnt := factorial[n-1-index]
	for i := 1; i <= n; i++ {
		// 同一个数只能用一次
		if used[i] {
			continue
		}
		// 一个n层高的树,有n-1!个结点,如果k超高该数,说明不在这个子树,跳过
		if cnt < k {
			k -= cnt
			continue
		}
		*path = append(*path, strconv.Itoa(i))
		used[i] = true
		dfs(n, k, index+1, path, factorial, used)
	}

}
