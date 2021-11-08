package tree

import (
	"fmt"
	"testing"
)

func TestNum105(t *testing.T)  {
	//tree := InitTree("[1,2,3,4,null,2,4,null,null,4]")
	//x := tree.LevelsNode()
	//fmt.Println(x)
	tree := InitTree("[1,2,3,4,null,2,4,null,null,4]")
	_, g := getGraphLink(tree, 2)
	fmt.Printf("%v", g)
	fmt.Printf("%v", []string{"a", "eh kdj"})
}
