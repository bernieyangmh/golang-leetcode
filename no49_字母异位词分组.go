package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	a := "eqwds"
	for _, value := range a {
		fmt.Println(string(value))
	}

}

// 暴力
// 拿到一个字符串,转成单个字符的数组,排序后的数组即是唯一值,拼成字符串放到hashmap里
func groupAnagrams(strs []string) [][]string {
	var hashMap = make(map[string][]string)
	var res [][]string
	for _, value := range strs {
		arr := strings.Split(value, "")
		sort.Strings(arr)
		var str string
		for _, s := range arr {
			str += s
		}
		if _, ok := hashMap[str]; ok {
			hashMap[str] = append(hashMap[str], value)
		} else {
			hashMap[str] = []string{value}
		}
	}
	for _, v := range hashMap {
		res = append(res, v)
	}
	return res
}
