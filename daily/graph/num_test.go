package graph

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	n := 10
	arr := make([]int, n)
	for i := range arr {
		arr[i] = rand.Intn(20)
	}
	l := func(i, j int) bool {
		return i <= j
	}
	a := &[]int{0, 0, 5, 1, 1, 18, 7, 19, 16, 7}
	newQueue(a, l)
	for i := range *a {

		fmt.Println(i, Pop(a, l))
	}
}

type Elem interface {
	int | string
}

func PrintSlice[T Elem](s []T) {
	for _, v := range s {
		println(v)
	}
}

func TestX(t *testing.T) {
	PrintSlice([]int{1, 2, 3})
	PrintSlice([]string{"a", "b", "c"})
}

func TestNum851(t *testing.T) {
	s := `a
	b
	`
	fmt.Println(s)
	datas := []struct {
		richer     [][]int
		quiet, ans []int
	}{
		{
			[][]int{{1, 0}, {2, 1}, {3, 1}, {3, 7}, {4, 3}, {5, 3}, {6, 3}},
			[]int{3, 2, 5, 4, 6, 1, 7, 0},
			[]int{5, 5, 2, 5, 4, 5, 6, 7},
		},
		{[][]int{}, []int{0}, []int{0}},
		{[][]int{{0, 1}, {1, 2}}, []int{1, 0, 2}, []int{0, 1, 1}},
	}
	for _, x := range datas {
		if ans := loudAndRich(x.richer, x.quiet); !reflect.DeepEqual(ans, x.ans) {
			fmt.Printf("wrong ans = %v, expected = %v\n", ans, x.ans)
		}
	}
}
func TestNum1494(t *testing.T) {
	var nexts []int
	nexts = append(nexts, 2)
	fmt.Println(nexts)
	datas := []struct {
		dependencies [][]int
		n, k, ans    int
	}{
		{[][]int{{2, 1}, {3, 1}, {1, 4}}, 4, 2, 3},
		{[][]int{{2, 1}, {3, 1}, {4, 1}, {1, 5}}, 5, 2, 4},
		{[][]int{}, 11, 2, 2},
	}
	for _, x := range datas {
		fmt.Println(minNumberOfSemesters(x.n, x.dependencies, x.k) == x.ans)
	}
}
