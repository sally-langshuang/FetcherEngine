package leet

import "testing"

func Test743(t *testing.T) {
	data := []struct {
		Times   [][]int;
		N, K, R int
	}{
		{[][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2, 2},
		{[][]int{{1, 2, 1}}, 2, 1, 1},
		{[][]int{{1, 2, 1}}, 2, 2, -1},
	}
	for _, d:= range data{
		r := F743(d.Times, d.N, d.K)
		if r != d.R {
			t.Errorf("%v ==> %d is wrong", d, r)
		}
	}
}
