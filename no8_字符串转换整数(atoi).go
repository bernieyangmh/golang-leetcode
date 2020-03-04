package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s:="-"
	fmt.Println(string(s[0]))
	a := "   -42"
	fmt.Println(myAtoi(a))

}

func myAtoi(str string) int {
	trim_str := strings.TrimSpace(str)
	//fmt.Println(trim_str)
	if len(trim_str) == 0 {
		return 0
	}
	var (
		flag   bool
		resStr string
	)

	for i := 0; i < len(trim_str); i++ {
		c := string(trim_str[i])
		if i == 0 {
			if c == "-" {
				continue
			} else if c == "+"{
				flag = true
				continue
			} else if !isNum(string(trim_str[i])) {
				//fmt.Println(c, i)
				return 0
			} else {
				flag = true
				resStr += c
			}
			continue
		}
		if isNum(c) {
			resStr += c
		} else {
			break
		}
	}
	resNum, _ := strconv.ParseInt(resStr, 10, 64)
	//fmt.Println(resNum)
	if !flag {
		resNum = resNum * -1
	}
	if resNum > (1<<31 - 1) {
		return 1<<31 - 1
	}
	if resNum < (1 << 31 * -1) {
		return 1 << 31 * -1
	}
	return int(resNum)

}

func isNum(s string) bool {
	if s == "0" || s == "1" || s == "2" || s == "3" || s == "4" || s == "5" || s == "6" || s == "7" || s == "8" || s == "9" {
		return true
	}
	return false
}
