package collections

//给定一个十进制数M，以及需要转换的进制数N。将十进制数M转化为N进制数


var char = []string{"0","1","2","3","4","5","6","7","8","9","A","B","C","D","E","F"}

func solve(M int, N int) string {
	// write code here

	f := false
	tmp := []string{}
	res := ""

	if M < 0 {
		M = -M
		f = true
	}

	for M > 0 {
		tmp = append(tmp, char[M%N])
		M /= N
	}
	if f {
		res = "-"
	}

	for i := len(tmp) - 1; i >= 0; i-- {
		res += tmp[i]
	}
	return res
}

//
//// 任意进制转10进制
//func anyToDecimal(num string, n int) int {
//	var new_num float64
//	new_num = 0.0
//	nNum := len(strings.Split(num, "")) - 1
//	for _, value := range strings.Split(num, "") {
//		tmp := float64(findkey(value))
//		if tmp != -1 {
//			new_num = new_num + tmp*math.Pow(float64(n), float64(nNum))
//			nNum = nNum - 1
//		} else {
//			break
//		}
//	}
//	return int(new_num)
//}