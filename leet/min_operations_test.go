package leet

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

func TestMin(t *testing.T)  {
	a := []int {1, 3, 5}
	fmt.Println(sort.SearchInts(a, 2))
	fmt.Println(sort.SearchInts(a, 3))
	fmt.Println(sort.SearchInts(a, -1))
	fmt.Println(sort.SearchInts(a, 99))
}

var tables = []struct {
	Arr    []int
	Target []int
	Result int
}{
	//{[]int{9, 4, 2, 3, 4}, []int{5, 1, 3}, 2},
	{[]int{4,7,6,2,3,8,6,1}, []int{6,4,8,1,3,2}, 3},
}

func TestB(t *testing.T)  {
	maxInt := 1<<63 -1
	s:= strconv.FormatUint(uint64(1<<64 - 1), 2)
	s= strconv.FormatUint(uint64(maxInt + maxInt + 1), 2)
	fmt.Printf("%s len=%d\n" ,s, len(s))
}
func TestA(t *testing.T)  {
	for _, d:= range tables {
		res := minOperations(d.Target, d.Arr)
		fmt.Printf("res==>%v\n", res)
		if res != d.Result {
			t.Errorf("%v error", d)
		}
	}


}