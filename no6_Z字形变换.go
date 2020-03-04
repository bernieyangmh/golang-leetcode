package main

//将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
//比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
//L   C   I   R
//E T O E S I I G
//E   D   H   N
//之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
//请你实现这个将字符串进行指定行数变换的函数：
//string convert(string s, int numRows);
//示例 1:
//输入: s = "LEETCODEISHIRING", numRows = 3
//输出: "LCIRETOESIIGEDHN"
//示例 2:
//输入: s = "LEETCODEISHIRING", numRows = 4
//输出: "LDREOEIIECIHNTSG"
//解释:
//L     D     R
//E   O E   I I
//E C   I H   N
//T     S     G

// 有几行就分几个字符串，遇到一个字符就加到对应的行，并改变行数
// 如果走到底或头，就改变方向
func convert(s string, numRows int) string {
	// 注意 只有1行的情况
	if numRows <= 1 {
		return s
	}

	res := ""
	rowStr := make([]string, numRows)
	curRow := 0
	// dir = 0 向下, 1 向上
	dir := 0
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		rowStr[curRow] += c
		// 如果走到底
		if curRow == numRows-1 {
			dir = 1
			curRow--
		} else if curRow == 0 {
			dir = 0
			curRow++
		} else if dir == 0 {
			curRow++
		} else {
			curRow--
		}
	}
	for i := 0; i < len(rowStr); i++ {
		res += rowStr[i]
	}
	return res
}

func convert2(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	res := []byte{}
	n := len(s)
	// 一次循环的长度
	cycLen := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < n; j += cycLen {
			res = append(res, s[j+i])
			//除去第一行和最后一行，其他行的值是j+i和j+c-i
			//j+i从循环开始向前走几行;j+c-i 循环结束回退几行
			if i != 0 && i != numRows-1 && j+cycLen-i < n {
				res = append(res, s[j+cycLen-i])

			}
		}
	}
	return string(res)
}
