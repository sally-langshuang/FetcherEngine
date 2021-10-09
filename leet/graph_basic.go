package leet

import "fmt"

type Node struct {
	ID int
}
type Edge struct {
	Node *Node
	Weight int
}

type Vertex struct {
	Node Node
	AdjacentNodeList *Edge
}

type Entry struct {
	Weight int
}
type Vertice struct {
	ID int
}

type G struct {
	AdjacencyList []Vertex
	AdjacencyMatrix struct {
		Matrix [][]Entry
		Vertices []Vertice
	}
}


type Element interface {
}


type Set interface{
	Search(s Set, key int) *Element
	Insert(s Set, x *Element)
	Delete(s Set, x *Element)
	Minimum(s Set) *Element
	Maximum(s Set) *Element
	Successor(s Set, x *Element) *Element
	Predecessor(s Set, x *Element) *Element
}

// finite length
type FiniteIntElementSet interface {
	Empty() bool
	Full() bool
	Push(int) error
	Pop() (int, error)
}

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

// finite length queue
type FiniteQueueArr struct {
	Arr []int
	head, tail int
}

func (q FiniteQueueArr) Empty() bool {
	return q.tail == q.head
}

func (q FiniteQueueArr) Full() bool {
	return q.tail % len(q.Arr) == q.head
}

func (q FiniteQueueArr) Push(val int) error {
	if q.Full() {
		return fmt.Errorf("cannot push to full queue")
	}
	q.Arr[q.tail] = val
	q.tail += 1
	return nil
}

func (q FiniteQueueArr) Pop() (val int, e error) {
	if q.Empty() {
		return -1, fmt.Errorf("cannot pop from empty queue")
	}
	val = q.Arr[q.head]
	q.head -= 1
	return
}