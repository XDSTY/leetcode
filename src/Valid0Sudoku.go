package main

/**
Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
判断一个9x9的数独是否正确。需要保证以下规则成立
Each row must contain the digits 1-9 without repetition.
每行都要包含1-9的不重复的数字
Each column must contain the digits 1-9 without repetition.
每列都包含1-9的不重复的数字
Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.
每个3x3的小格子都要包含1-9的不重复数字
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-sudoku
*/
func isValidSudoku(board [][]byte) bool {
	// 判断每行和每列
	for i := 0; i < 9; i++ {
		rows := [10]int{}
		columns := [10]int{}
		gards := [10]int{}
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if rows[board[i][j]-'0'] == 1 {
					return false
				}
				rows[board[i][j]-'0'] = 1
			}

			if board[j][i] != '.' {
				if columns[board[j][i]-'0'] == 1 {
					return false
				}
				columns[board[j][i]-'0'] = 1

			}
		}
		li := i / 3 * 3
		lj := i % 3 * 3
		for ti := li; ti < li+3; ti++ {
			for j := lj; j < lj+3; j++ {
				if board[ti][j] == '.' {
					continue
				}
				if gards[board[ti][j]-'0'] == 1 {
					return false
				}
				gards[board[ti][j]-'0'] = 1
			}
		}
	}
	return true
}
