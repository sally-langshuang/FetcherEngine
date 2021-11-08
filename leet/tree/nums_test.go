package tree

import (
	"fmt"
	"testing"
)

func TestNum105(t *testing.T)  {
	//tree := InitTree("[1,2,3,4,null,2,4,null,null,4]")
	//x := tree.LevelsNode()
	//fmt.Println(x)
	tree := InitTree("[1,2,3,4,null, 5,6343874,null,null,7,678]")
	//tree := InitTree("[1,2]")
	fmt.Println(tree.MaxVal())
	g := getGraphLink(tree, -1)
	fmt.Printf("%v", g)
	fmt.Printf("%v", []string{"a", "eh kdj"})
}
