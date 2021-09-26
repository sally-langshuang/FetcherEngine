package leet

import (
	"fmt"
	"testing"
)

func TestIdx(t *testing.T)  {
	fmt.Println(3/2)
}


var data = [] struct{
	Root []int; Target,K int; Res []int
} {
	{[]int{3,5,1,6,2,0,8,null,null,7,4}, 5, 2, []int{7,4,1}},
}

func TestDistinct(t *testing.T)  {
	for _, d:= range data{
		DistinctK(d.Root, d.Target, d.K)
	}
}

func TestK(t *testing.T)  {
	yy := []int{1, 2}

	fmt.Println(yy[3:])
	type X struct {
		V int
		Y *X
	}
	e := &X{V: 3}
	arr := [1]*X{}
	e.Y = &X{V: 10}
	arr[0] = e
	fmt.Println(&arr[0], &e)
}
