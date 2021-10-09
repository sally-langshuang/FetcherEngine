package datastruct

import "fmt"

// finite length stack
type FiniteStackArr struct {
	Arr        []int
	TopSuccIdx int
}

func (s FiniteStackArr) Empty() bool {
	return s.TopSuccIdx == 0
}

func (s FiniteStackArr) Full() bool {
	return s.TopSuccIdx == len(s.Arr)
}

func (s FiniteStackArr) Push(val int) error {
	if s.Full() {
		return fmt.Errorf("Stack overflows, cannot push")
	}
	s.Arr[s.TopSuccIdx] = val
	s.TopSuccIdx += 1
	return nil
}

func (s FiniteStackArr) Pop() (v int, e error){
	if s.Empty() {
		return -1, fmt.Errorf("no element to pop")
	}
	v = s.Arr[s.TopSuccIdx- 1]
	s.TopSuccIdx -= 1
	return
}