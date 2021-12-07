package matrix

import "fmt"

func onlyOne(x int) int {
	idx := -1
	for i:=0; i< 9; i++ {
		if x % 2 != 0 {
			if idx == -1 {
				idx = i
			} else {
				return -1
			}
		}
		x = x >> 1
	}
	if x != 0 {
		return -1
	}
	return idx
}

func findRoot(board [][]byte, f [3][9]int) (int, int, byte) {
	for i := range board {
		for j := range board[i] {
			//fmt.Println("i=",i,"j=",j)
			if val := board[i][j] - '0'; val > 9 || val < 1 {
				m := f[0][i] ^ f[1][j] ^ f[2][i/3*3+j/3]
				if i == 8 && j == 6{
					x := fmt.Sprintf("%09b", m)
					fmt.Println(x)
				}
				if k:= onlyOne(m); k != -1 {
					return i, j, '0' + byte(k)
				}
			}
		}
	}
	return -1, -1, '0'
}

func getMask(i, j int, f [3][9]int) int {
	return  f[0][i] ^ f[1][j] ^ f[2][i / 3 * 3 + j / 3]
}

func updateState(v byte, i, j int, f *[3][9]int, board [][]byte) (bool){
	if v >= '1' && v <= '9'{
		m := 1 << (v -'1')
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

	for {
		i, j, k := findRoot(board, f)
		if i != -1 {
			updateState(k, i, j, &f, board)
		} else {
			break
		}
	}
}


