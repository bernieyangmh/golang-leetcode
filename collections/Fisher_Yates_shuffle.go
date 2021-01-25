package main

import (
	"math/rand"
	"time"
)

func FYShuffle(arr *[]int) *[]int {
	if len(*arr) < 2 {
		return arr
	}
	rand.Seed(time.Now().Unix())
	//选一个数，从剩下的数中随机挑一个，进行交换
	for i := len(*arr) - 1; i > 0; i-- {
		curPoint := rand.Intn(i+1)
		if curPoint == i {
			continue
		}
		(*arr)[i], (*arr)[curPoint] = (*arr)[curPoint], (*arr)[i]
	}
	return arr
}
