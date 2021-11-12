package dp

import (
	"fmt"
	"strconv"
	"strings"
)
type point struct {
	x, y int
}
func findMaxElem(a [][]int)  int{
	var res int
	for _, i:= range a {
		for _, j := range i {
			if j > res {
				res = j
			}
		}
	}
	return res
}

func printRow(a []int, maxWidth int) string {
	b := strings.Builder{}
	fmt.Fprintf(&b,"[")
	for _, row:= range a {
		fmt.Fprintf(&b, fmt.Sprintf("%-[2]*[1]d ", row, maxWidth))
	}
	fmt.Fprintf(&b,"]\n")
	return b.String()
}

func ArrString(a [][]int, p *point) string {
	res := ""
	maxWidth := len(strconv.Itoa(findMaxElem(a)))
	for i, x:= range a {
		if p != nil && p.x == i{
			res += strings.Repeat(" ", 1 + (maxWidth + 1) * p.y ) + "V\n"
		}
		res += printRow(x, maxWidth)
	}
	res += strings.Repeat("-", len(a[0])*(maxWidth +1)+2)+ "\n"
	return res
}
func makeIntArr(m, n int) [][]int  {
	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]int, n)
	}
	return arr
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}