// 给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] + D[l] = 0。
// 为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。所有整数的范围在 -228 到 228 - 1 之间，最终结果不会超过 231 - 1 。
// 例如:
// 输入:
// A = [ 1, 2]
// B = [-2,-1]
// C = [-1, 2]
// D = [ 0, 2]
// 输出:
// 2
// 解释:
// 两个元组如下:
// 1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
// 2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
package main



// 一个map保存ab的和,一个map保存cd的和,看ab和的相反数在不在cd和的map里
func fourSumCount(A []int, B []int, C []int, D []int) int {
	mapA := make(map[int]int)
	mapB := make(map[int]int)

	var res int

	for _, v1 := range A {
		for _, v2 := range B {
			if _, ok := mapA[v1+v2]; ok {
				mapA[v1+v2] = mapA[v1+v2] + 1
			} else {
				mapA[v1+v2] = 1
			}
		}
	}

	for _, v1 := range C {
		for _, v2 := range D {
			if _, ok := mapB[v1+v2]; ok {
				mapB[v1+v2] = mapB[v1+v2] + 1
			} else {
				mapB[v1+v2] = 1
			}
		}
	}

	// fmt.Println(mapA, mapB)

	// 次数相乘
	for valueA, cntA := range mapA {
		if cntB, ok := mapB[-valueA]; ok {
			res += cntA * cntB
		}
	}
	return res

}
