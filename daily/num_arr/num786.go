package num_arr

import (
	. "github.com/sally-langshuang/FetcherEngine/util/heap"
)

// tosee
func kthSmallestPrimeFraction(arr []int, k int) []int {
	q := &FractionHeap{Arr: arr, FractionIdx: [][]int{}}
	for i:= 0; i < len(arr) - 1; i++ {
		Push(q, []int{i, len(arr) - 1})
	}
	for i := 1; q.Len() != 0; i++ {
		idx := Pop(q).([]int)
		if i == k {
			return q.GetFractionByIdx(idx)
		}
		if idx[1] - idx[0] > 1 {
			Push(q, []int{idx[0], idx[1] - 1})
		}
	}
	return nil
}

