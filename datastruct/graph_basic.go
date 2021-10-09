package datastruct

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