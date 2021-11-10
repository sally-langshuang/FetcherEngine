package tree

//RYGBW
func handSet(hand string) map[rune]int {
	s := make(map[rune]int, 0)
	for _, b:= range hand {
		n := s[b]
		s[b] = n + 1
	}
	return s
}

func handNum(h map[rune]int) int {
	n := 1
	for _, v:= range h {
		n *= v
	}
	return n
}



func eliminate(hand map[rune]int, board string) string {
	if len(hand) == 0 {
		return board
	}
	return ""
}

func Zuma(board, hand string) {

}
