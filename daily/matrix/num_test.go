package matrix

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestNum807(t *testing.T) {
	datas := []struct {
		grid [][]int
		ans  int
	}{
		{[][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}, 35},
	}
	for i := range datas {
		fmt.Println(maxIncreaseKeepingSkyline(datas[i].grid) == datas[i].ans)
	}
}

func TestA(t *testing.T) {
	s1 := "abc"
	s3 := s1[1:]
	s4 := s1[:1]
	s2 := "abcd"
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s1)))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s2)))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s3)))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s4)))
	//fmt.Println(!true || false)
	//fmt.Println(!false && true)
	//fmt.Println(4 % 2)
}
func TestNum794(t *testing.T) {
	datas := []struct {
		board []string
		ans   bool
	}{
		{[]string{"XXX", "OOX", "OOX"}, true},
		{[]string{"XXO", "XOX", "OXO"}, false},
	}
	for _, x := range datas {
		fmt.Println(validTicTacToe(x.board) == x.ans)
	}

}
func TestNum1034(t *testing.T) {
	datas := []struct {
		grid, ans       [][]int
		row, col, color int
	}{
		{[][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, [][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}, 1, 1, 2},
		{[][]int{{1, 2, 1, 2, 1, 2}, {2, 2, 2, 2, 1, 2}, {1, 2, 2, 2, 1, 2}}, [][]int{{1, 1, 1, 1, 1, 2}, {1, 2, 1, 1, 1, 2}, {1, 1, 1, 1, 1, 2}}, 1, 3, 1},
	}
	for _, d := range datas {
		fmt.Println(reflect.DeepEqual(colorBorder(d.grid, d.row, d.col, d.color), d.ans))
	}
}
func TestNum37(t *testing.T) {
	datas := []struct {
		board, ans [][]byte
	}{
		{
			board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			ans: [][]byte{
				{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
				{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
				{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
				{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
				{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
				{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
				{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
				{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
				{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
			},
		},
		{
			board: [][]byte{
				{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
				{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
				{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
				{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
				{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
				{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
				{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
			},
			ans: [][]byte{
				{'5', '1', '9', '7', '4', '8', '6', '3', '2'},
				{'7', '8', '3', '6', '5', '2', '4', '1', '9'},
				{'4', '2', '6', '1', '3', '9', '8', '7', '5'},
				{'3', '5', '7', '9', '8', '6', '2', '4', '1'},
				{'2', '6', '4', '3', '1', '7', '5', '9', '8'},
				{'1', '9', '8', '5', '2', '4', '3', '6', '7'},
				{'9', '7', '5', '8', '6', '3', '1', '2', '4'},
				{'8', '3', '2', '4', '9', '1', '7', '5', '6'},
				{'6', '4', '1', '2', '7', '5', '9', '8', '3'},
			},
		},
	}
	for _, x := range datas {
		solveSudoku(x.board)
		fmt.Println(reflect.DeepEqual(x.board, x.ans))
	}
}

func TestNum36(t *testing.T) {
	datas := []struct {
		board [][]byte
		ans   bool
	}{
		{
			board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			ans: true,
		},
		{
			board: [][]byte{
				{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			ans: false,
		},
	}
	for _, x := range datas {
		fmt.Println(isValidSudoku(x.board) == x.ans)
	}
}
