package dp

import (
	"fmt"
	"strconv"
	"strings"
)

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
func printA(a []int, maxWidth int) string {
	b := strings.Builder{}
	fmt.Fprintf(&b,"[")
	for _, i:= range a {
		fmt.Fprintf(&b, fmt.Sprintf("%-[2]*[1]d ", i, maxWidth))
	}
	fmt.Fprintf(&b,"]\n")
	return b.String()
}

func ArrString(a [][]int) string {
	res := ""
	maxElem := findMaxElem(a)
	for _, x:= range a {
		res += printA(x, len(strconv.Itoa(maxElem)))
	}
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