package matrix

import "fmt"

func onlyOne(x int) int {
	idx := -1
	for i:=0; i< 9; i++ {
		if x % 2 != 0 {
			if idx != -1 {
				idx = i
			} else {
				return -1
			}
		}
		x = x >> 1
	}
	return idx
}

func findRoot(board [][]byte, f [3][9]int) (int, int) {
	for i := range board {
		for j := range board[i] {
			if val := board[i][j] - '0'; val > 9 {
				m := f[0][i] ^ f[1][j] ^ f[2][i/3*3+j/3]
				if onlyOne(m) != -1 {
					return i, j
				}
				fmt.Printf("%09b\n", m)
			}
		}
	}
	return -1, -1
}

func getMask(i, j int, f [3][9]int) int {
	return  f[0][i] ^ f[1][j] ^ f[2][i / 3 * 3 + j / 3]
}

func updateState(v byte, i, j int, f *[3][9]int, board [][]byte) (bool){
	if val := board[i][j] - '0'; val < 10 {
		m := 1 << (val - 1)
		if f[0][i] & m != 0 || f[1][j] & m != 0  || f[2][i / 3 * 3  + j / 3] & m != 0{
			return false
		}
		f[0][i] ^= m
		f[1][j] ^= m
		f[2][i / 3 * 3 + j / 3] ^= m
	}
	return true
}

func solveSudoku(board [][]byte)  {
	f := [3][9]int{}
	for i := range board {
		for j := range board[i] {
			updateState(board[i][j], i, j, &f, board)
		}
	}

	for i := range board {
		for j := range board[i] {
			if val := board[i][j] - '0'; val > 9 {
				m := f[0][i] ^ f[1][j] ^ f[2][i / 3 * 3 + j / 3]
				if onlyOne(m) != -1 {
				}
				fmt.Printf("%09b\n", m)
			}
		}
	}
}


