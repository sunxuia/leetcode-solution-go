package q036_valid_sudoku

import (
	"testing"
)

func TestIsValidSudoku(t *testing.T) {
	doTest(t, isValidSudoku)
}

func isValidSudoku(board [][]byte) bool {
	return check(board, func(m, n int) (int, int) {
		return m, n
	}) && check(board, func(m, n int) (int, int) {
		return n, m
	}) && check(board, func(m, n int) (int, int) {
		return m/3*3 + n/3, m%3*3 + n%3
	})
}

func check(board [][]byte, provider func(int, int) (int, int)) bool {
	for m := 0; m < 9; m++ {
		set := 0
		for n := 0; n < 9; n++ {
			i, j := provider(m, n)
			if board[i][j] != '.' {
				offset := board[i][j] - '0'
				if set>>offset&1 == 1 {
					return false
				}
				set |= 1 << offset
			}
		}
	}
	return true
}
