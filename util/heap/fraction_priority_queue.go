package heap

import (
	"fmt"
	"strings"
)

type Fraction []int

type FractionHeap struct {
	FractionIdx [][]int
	Arr []int
}
func (h FractionHeap) GetFractionByIdx(idx []int) Fraction {
	return Fraction{h.Arr[idx[0]], h.Arr[idx[1]]}
}
func (h FractionHeap) GetFraction(i int) Fraction {
	idx := h.FractionIdx[i]
	return Fraction{h.Arr[idx[0]], h.Arr[idx[1]]}
}

func (h FractionHeap) Len() int {
	return len(h.FractionIdx)
}
func (h FractionHeap) Less(i, j int) bool {
	x, y := h.GetFraction(i), h.GetFraction(j)
	return x[0] * y[1] < y[0] * x[1]
}
func (h FractionHeap) Swap(i, j int)  {
	h.FractionIdx[i], h.FractionIdx[j] = h.FractionIdx[j], h.FractionIdx[i]
}

func (h *FractionHeap) Pop() interface{}  {
	idx := (*h).FractionIdx[h.Len() - 1]
	(*h).FractionIdx = (*h).FractionIdx[:h.Len() - 1]
	return idx
}

func (h *FractionHeap) Push(x interface{})  {
	(*h).FractionIdx = append((*h).FractionIdx, x.([]int))
}
func (h *FractionHeap) String() string {
	b := &strings.Builder{}
	fmt.Fprintf(b, "[")
	for i := range h.FractionIdx {
		fmt.Fprintf(b, "%d/%d %d/%d(%.3f), ",h.GetFraction(i)[0],h.GetFraction(i)[1],h.FractionIdx[i][0],h.FractionIdx[i][1], float64(h.GetFraction(i)[0]) / float64(h.GetFraction(i)[1]))
	}
	fmt.Fprintf(b, "]\n")
	return b.String()
}


