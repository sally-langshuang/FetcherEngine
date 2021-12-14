package graph

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum851(t *testing.T)  {
	datas := []struct{
		richer [][]int
		quiet, ans []int
	}{
		{

		[][]int{{1,0},{2,1},{3,1},{3,7},{4,3},{5,3},{6,3}},
		[]int{3,2,5,4,6,1,7,0},
		[]int{5,5,2,5,4,5,6,7},
		},
		{[][]int{}, []int{0}, []int{0}},
	}
	for _, x := range datas {
		if ans := loudAndRich(x.richer, x.quiet); !reflect.DeepEqual(ans, x.ans) {
			fmt.Printf("wrong ans = %v, expected = %v\n", ans, x.ans)
		}
	}
}
func TestNum1494(t *testing.T)  {
	var nexts []int
	nexts = append(nexts, 2)
	fmt.Println(nexts)
	datas := []struct{
		dependencies [][]int
		n, k, ans int
	}{
		{[][]int{{2,1},{3,1},{1,4}}, 4, 2, 3},
		{[][]int{{2,1},{3,1},{4,1},{1,5}}, 5, 2, 4},
		{[][]int{}, 11, 2, 2},
	}
	for _, x:= range datas{
		fmt.Println(minNumberOfSemesters(x.n, x.dependencies, x.k) == x.ans)
	}
}
