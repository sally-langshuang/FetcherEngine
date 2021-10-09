package leet

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

const (
	Less        = 1 << iota
	Equal
	Larger
	LessEqual   = Less & Equal
	LargerEqual = Larger & Equal
)

type Element interface {
}

type Key interface {
}

type CompareElement interface {
	Compare(e CompareElement) byte
}

type Set interface{
	Search(s Set, key Key) *Element
	Insert(s Set, x *Element)
	Delete(s Set, x *Element)
	Minimum(s Set) *Element
	Maximum(s Set) *Element
	Successor(s Set, x *Element) *Element
	Predecessor(s Set, x *Element) *Element
}
