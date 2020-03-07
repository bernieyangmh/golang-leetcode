package main

import (
	"fmt"
)

//判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。
//数字 1-9 在每一行只能出现一次。
//数字 1-9 在每一列只能出现一次。
//数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。

func main()  {
	a := "012"
	for _, s := range a {
		fmt.Println(s-48)

	}

}


func isValidSudoku(board [][]byte) bool {
	var row, col, box []map[byte]bool
	for i := 0; i < 9; i++ {
		row = append(row, make(map[byte]bool))
		col = append(col, make(map[byte]bool))
		box = append(box, make(map[byte]bool))
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			//"0" ==> 48
			//j/3+(i/3)*3 ==> 0+(0,1,2); 3+(0,1,2); 6+(0,1,2)
			//j/3 =>0，1，2    i/3*3 => 0,3,6
			curNum := board[i][j]
			if row[i][curNum]  {
				return false
			}
			if col[j][curNum]  {
				return false
			}
			if box[j/3 + (i/3) * 3][curNum]  {
				return false
			}
			row[i][curNum] = true
			col[j][curNum] = true
			box[j/3+(i/3)*3][curNum] = true

		}
	}
	return true
}


//位运算
func isValidSudoku(board [][]byte) bool {
	n := len(board)
	rectSets := make([]int16, n*n/9) // sets for inner rectangles
	for i := 0; i < n; i++ {
		rowSet, colSet := int16(0), int16(0) // sets for a row and a column
		for j := 0; j < n; j++ {
			// checking a column and an inner rectangle
			if num := board[i][j]; num != '.' {
				numBit := int16(1 << (num - '0')) // calc a #num bit
				if rowSet&numBit > 0 {
					return false
				}
				rowSet |= numBit // set the bit

				// set the bit for a inner rectangle: ((n/3)*(i/3), j/3) - its coordinates
				if rectSets[(n/3)*(i/3)+j/3]&numBit > 0 {
					return false
				}
				rectSets[(n/3)*(i/3)+j/3] |= numBit
			}

			// checking column
			if num := board[j][i]; num != '.' {
				numBit := int16(1 << (num - '0'))
				if colSet&numBit > 0 {
					return false
				}
				colSet ^= numBit
			}
		}
	}
	return true
}