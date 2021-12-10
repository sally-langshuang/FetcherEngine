package matrix

func validTicTacToe(board []string) bool {
	vals := map[byte]int{'X': 0, 'O': 1}
	stats := [4][]int{{0, 0},{0,0,0},{0,0,0},{0,0}}
	for i := range board{
		for j := range board[i] {
			if v, ok := vals[board[i][j]]; ok {
				stats[0][v] += 1
				stats[1][i] |= 1 << (j + v * 3)
				stats[2][j] |= 1 << (i + v * 3)
				if i == j {
					stats[3][0] |= 1 << (i + v * 3)
				}
				if i + j == 2 {
					stats[3][1] |= 1<< (i + v * 3)
				}
			}

		}
	}

	threeCount1, threeCount2 := 0, 0
	for i:= 1; i < len(stats); i++ {
		for j:= range stats[i] {
			if stats[i][j] & 0b111000 == 0b111000 {
				threeCount2 += 1
			} else if stats[i][j] & 0b111 == 0b111 {
				threeCount1 += 1
			}
		}
	}
	if threeCount2 >= 1 && threeCount1 >= 1 {
		return false
	}
	if threeCount1 >= 1 && stats[0][0] - stats[0][1] != 1 ||
		threeCount2 >= 1 && stats[0][0] - stats[0][1] != 0 ||
		(stats[0][0] - stats[0][1] > 1 || stats[0][0] - stats[0][1] < 0) {
		return false
	}
	return true
}
