package graph

import (
	"fmt"
	"testing"
)

func TestNum1494(t *testing.T)  {
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
