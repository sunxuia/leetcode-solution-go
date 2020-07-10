package q037_sudoku_solver

import (
	"testing"
)

func TestSolveSudoku(t *testing.T) {
	doTest(t, solveSudoku)
}

func solveSudoku(board [][]byte) {
	row, col, area := [9]int{}, [9]int{}, [9]int{}
	stack := [81]byte{}

	// 预先初始化已经有的数字, 这些状态是不会变更的
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				state := board[i][j] - '0'
				row[i] |= 1 << state
				col[j] |= 1 << state
				area[i/3*3+j/3] |= 1 << state
				// 0xff 是哨兵, 表示这部分数字是固定的
				stack[i*9+j] = 0xff
			}
		}
	}

	for pos, offset := 0, 1; pos < 81; pos += offset {
		if stack[pos] == 0xff {
			continue
		}

		i, j := pos/9, pos%9
		// 删除之前的状态
		row[i] &= 0xffff ^ (1 << stack[pos])
		col[j] &= 0xffff ^ (1 << stack[pos])
		area[i/3*3+j/3] &= 0xffff ^ (1 << stack[pos])

		// 寻找下一个状态
		state := stack[pos] + 1
		for row[i]>>state&1 == 1 ||
			col[j]>>state&1 == 1 ||
			area[i/3*3+j/3]>>state&1 == 1 {
			state++
		}

		if state == 10 {
			// 没找到, 回退
			stack[pos] = 0
			offset = -1
		} else {
			// 找到, 前进
			board[i][j] = state + '0'
			stack[pos] = state
			row[i] |= 1 << state
			col[j] |= 1 << state
			area[i/3*3+j/3] |= 1 << state
			offset = 1
		}
	}
}
