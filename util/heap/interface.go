package heap

type SortInterface interface {
	Less (i, j int) bool
	Swap (i, j int)
	Len() int
}

type HeapInterface interface {
	SortInterface
	Push(x interface{}) // push to tail
	Pop() interface{} // pop last
}

// O(n)?
func Init(h HeapInterface)  {
	lastParent := h.Len() / 2 - 1
	for i := lastParent; i >= 0; i-- {
		down(h, i, h.Len())
	}

}

func Push(h HeapInterface, x interface{})  {
	h.Push(x)
	up(h, h.Len() - 1)
}

func Pop(h HeapInterface) interface{} {
	// pop root
	last := h.Len() - 1
	h.Swap(0, last)
	down(h, 0, last)
	return h.Pop()
}

func Remove(h HeapInterface, i int) interface{} {
	last := h.Len() - 1
	if last != i {
		h.Swap(i, last)
		down(h, i, last)
		//Update(h, i, last)?
	}
	return h.Pop()
}

func Update(h HeapInterface, i, n int)  {
	if !down(h, i, n) {
		up(h, i)
	}
}

func Fix(h HeapInterface, i int)  {
	Update(h, i, h.Len())
}

func up(h HeapInterface, i int)  {
	for {
		p := (i -1) / 2
		if p == i || !h.Less(i, p) {
			break
		}
		h.Swap(p, i)
		i = p
	}
}

func down(h HeapInterface, i0, n int) bool {
	i := i0
	for {
		c := i * 2 + 1 // left child
		if c >= n || c < 0{
			break
		}
		if c2 := c + 1; c2 < n && h.Less(c2, c) {
			c = c2
		}
		if !h.Less(c, i) {
			break
		}
		h.Swap(i, c)
		i = c
	}
	return i > i0
}
