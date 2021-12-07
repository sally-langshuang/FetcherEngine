package matrix

import (
	"fmt"
	"strings"
)


func printArr(arr [][]byte)  {
	b := &strings.Builder{}
	for i:= range arr {
		for j := range arr[i] {
			fmt.Fprintf(b, fmt.Sprintf("%c ", arr[i][j]))
		}
		fmt.Fprintf(b, "\n")
	}
	print(b.String())
}


func updateState(v byte, i, j int, f [][]int, board [][]byte) bool {
	board[i][j] = v
	if v >= '1' && v <= '9'{
		m := 1 << (v -'1')
		if f[0][i] & m != 0 || f[1][j] & m != 0  || f[2][i / 3 * 3  + j / 3] & m != 0{
			return false
		}
		f[0][i] |= m
		f[1][j] |= m
		f[2][i / 3 * 3 + j / 3] |= m
	}
	return true
}
func clearState(v byte, i, j int, f [][]int, board [][]byte) (bool){
	board[i][j] = '.'
	if v >= '1' && v <= '9'{
		m := 1 << (v -'1')
		if f[0][i] & m != 0 {
			f[0][i] ^= m
		}
		if f[1][j] & m != 0 {
			f[1][j] ^= m
		}
		if f[2][i / 3 * 3 + j / 3] & m != 0 {
			f[2][i / 3 * 3 + j / 3] ^= m
		}
	}
	return true
}

func solve(board [][]byte, i, j int, f [][]int) bool {
	if i == -1 && j == -1 {
		return true
	}
	m := f[0][i] | f[1][j] | f[2][i/3*3+j/3]
	for x := 0; x < 9; x++ {
		if m & (1 << x) == 0 {
			ok := updateState(byte('1' + x), i, j, f, board)
			if !ok {
				clearState(byte('1' + x), i, j, f, board)
			} else {
				ni, nj := findNextBlank(board)
				ok1 := solve(board, ni, nj, f)
				if ok1 {
					return true
				} else {
					clearState(byte('1' + x), i, j, f, board)
				}
			}
		}
	}
	return false
}

func findNextBlank(board [][]byte) (int, int){
	for i := range board {
		for j := range board {
			if board[i][j] == '.' {
				return i, j
			}
		}
	}
	return -1, -1
}


func solveSudoku(board [][]byte)  {
	f := [][]int{}
	for i:=0; i < 3; i++ {
		f = append(f, make([]int, 9))
	}
	for i := range board {
		for j := range board[i] {
			updateState(board[i][j], i, j, f, board)
		}
	}
	i, j := findNextBlank(board)
	if i == -1 || j == -1 {
		return
	}
	solve(board, i, j, f)
	//printArr(board)
}


