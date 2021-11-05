package tree

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func (t *TreeNode) LevelsNode() [][]*TreeNode {
	if t == nil {return nil}
	res := make([][]*TreeNode, 0)
	currentLevel := make([]*TreeNode, 0)
	nextLevel := make([]*TreeNode, 0)
	for ;len(currentLevel) > 0; {
		for _, n := range currentLevel {
			if n != nil {
				nextLevel = append(nextLevel, n.Left, n.Right)
			}
		}
		res = append(res, nextLevel)
	}
	return res
}

func (t TreeNode) MaxVal() int {
	return max(t.Val, t.Left.MaxVal(), t.Right.MaxVal() )
}

func MaxDepth(t *TreeNode) int {
	if t == nil {
		return 0
	}
	return max(MaxDepth(t.Left), MaxDepth(t.Right)) + 1
}

func maxInt(num int, nums...int) int {
	res := num
	for _, i:= range nums{
		if i > res {
			res = i
		}
	}
	return res
}

//
//func (TreeNode) Init(s string) *TreeNode {
//
//}
type NumArr struct {
	arr []*int
}

func (a NumArr)String() string {
	var b strings.Builder
	fmt.Fprint(&b, "[")
	for idx, i:= range a.arr {
		if i == nil {
			fmt.Fprint(&b, "null" )
		} else {
			fmt.Fprintf(&b, "%d",*i)
		}
		if idx != len(a.arr) - 1 {
			fmt.Fprint(&b, ",")
		}

	}
	fmt.Fprint(&b, "]")
	return b.String()
}

func exactElem(s string) []string {
	s = strings.TrimFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	return strings.Split(s, ",")
}

func InitTree(s string) *TreeNode {
	numArr := exactInt(s)
	if len(numArr.arr) == 0 {
		return nil
	}
	root := &TreeNode{Val: *numArr.arr[0]}
	i := 1
	for parentLevel := []*TreeNode{root}; len(parentLevel) > 0; {
		nextLevel := make([]*TreeNode, 0)
		for _, n := range parentLevel {
			if i < len(numArr.arr) {
				if numArr.arr[i] != nil {
					n.Left = &TreeNode{Val: *numArr.arr[i]}
					nextLevel = append(nextLevel, n.Left)
				}
				i++
			} else{
				break
			}
			if i < len(numArr.arr) {
				if numArr.arr[i] != nil{
					n.Right = &TreeNode{Val: *numArr.arr[i]}
					nextLevel = append(nextLevel, n.Right)
				}
				i++
			} else {
				break
			}
		}
		parentLevel, nextLevel = nextLevel, make([]*TreeNode, 0)
	}
	return root
}

func exactInt(s string) NumArr {
	res := make([]*int, 0)
	elems := exactElem(s)
	for _, i:= range elems {
		if i == "null" {
			res = append(res, nil)
		} else{
			num, err := strconv.Atoi(i)
			if err == nil {
				res = append(res, &num)
			} else {
				panic(err)
			}
		}
	}
	return NumArr{arr: res}
}
// [1,2,3,4,null,2,4,null,null,4]
