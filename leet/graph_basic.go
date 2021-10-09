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

// 有限栈

type Stack1 struct {
	Arr []int
	TopToPushIdx int
}

func (s Stack1) Empty() bool {
	return s.TopToPushIdx == 0
}

func (s Stack1) Push(val int) error {
	if s.TopToPushIdx == len(s.Arr) {
		return fmt.Errorf("stack overflows")
	}
	s.Arr[s.TopToPushIdx] = val
	s.TopToPushIdx += 1
	return nil
}

func (s Stack1) Pop() (v int, e error){
	if s.Empty() {
		return -1, fmt.Errorf("no element to pop")
	}
	v = s.Arr[s.TopToPushIdx - 1]
	s.TopToPushIdx -= 1
	return
}