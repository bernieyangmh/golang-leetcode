package main

import (
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	add := 0
	ans := ""
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result := x + y + add
		ans = strconv.Itoa(result%10) + ans
		add = result / 10
	}
	return ans
}

func addStrings(num1 string, num2 string) string {
	var arr []byte
	// 进位
	Carrying := 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		n1, n2 := 0, 0
		if i >= 0 {
			n1 = int(num1[i] - '0')
		}
		if j >= 0 {
			n2 = int(num2[j] - '0')
		}
		// r + 10 * Carrying = sum
		sum := n1 + n2 + Carrying
		r := sum % 10
		Carrying = sum / 10
		arr = append(arr, byte(r+int('0')))
	}
	if Carrying > 0 {
		arr = append(arr, '1')
	}
	//reverse 调转一半
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	//convert string
	return string(arr)
}