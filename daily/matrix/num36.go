package matrix

func isValidSudoku(board [][]byte) bool {
	f := [3][9]int{}
	for i := range board {
		for j := range board[i] {
			if val := board[i][j] - '0'; val < 10 {
				m := 1 << (val - 1)
				if f[0][i] & m != 0 || f[1][j] & m != 0  || f[2][i / 3 * 3  + j / 3] & m != 0{
					return false
				}
				f[0][i] ^= m
				f[1][j] ^= m
				f[2][i / 3 * 3 + j / 3] ^= m
			}
		}
	}
	return true
}
