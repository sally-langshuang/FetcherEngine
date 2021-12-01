package heap

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int)  {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{}  {
	x := (*h)[len(*h) - 1]
	*h = (*h)[:len(*h) - 1]
	return x
}

func (h *IntHeap) Push(x interface{})  {
	*h = append(*h, x.(int))
}


