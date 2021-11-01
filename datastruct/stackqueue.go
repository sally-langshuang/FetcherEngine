package datastruct

// finite length
type FiniteIntElementSet interface {
	Empty() bool
	Full() bool
	Push(int) error
	Pop() (int, error)

}

