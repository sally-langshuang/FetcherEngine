package heap

import (
	"fmt"
	"testing"
)


func TestA(t *testing.T)  {
	p := &IntHeap{7, 10, 3}
	Init(p)
	fmt.Println(p)
	Push(p, 6)
	Push(p,5)
	Push(p,2)

	fmt.Println(Pop(p))
	fmt.Println(p)
	fmt.Println(Pop(p))
	fmt.Println(Pop(p))
	fmt.Println(Pop(p))
	fmt.Println(Pop(p))
	fmt.Println(Pop(p))
}
