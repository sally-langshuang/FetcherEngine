package tree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum5(t *testing.T) {
	fmt.Println(Palindrome{center: 3, armLen: 4}.S(produceS("abc")))
   //fmt.Println(longestPalindrome("babad"))
   //fmt.Println(longestPalindrome("ccc"))
   //fmt.Println(longestPalindrome("bananas"))
}
func Test488(t *testing.T)  {
	m1 := map[rune]int{'R':3, 'L':2}
	m2 := map[rune]int{'L':2, 'R':3, 'x':6}
	if  reflect.DeepEqual(m1, m2) {
		fmt.Println("m1 == m2")
	} else {
		fmt.Println("m1 != m2")
	}

}
func TestNum116(t *testing.T)  {
	tree := InitTree("[1,2,3,4,5,6,7]")
	//tree := InitTree("[1,2,3]")
	connect(tree)
	fmt.Println(getGraphLink(tree, -1))
}
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
