package tree

import "testing"

func TestPreorder(t *testing.T)  {
	var s TreeSerialize
	s = pre
	tree := s.Deserialize("[1,2,#,4,#,#,3,#,#]")
	tree.Print()
}
