package tree

import (
	"fmt"
	"strconv"
	"strings"
)
const (
	Null = "#"
    pre PreOrderTraverseSerialize = iota
)

func sepString(data string) []*TreeNode {
	datas := strings.Split(strings.TrimFunc(data, func(c rune) bool{return c == '[' || c == ']' || c == ' '}), ", ")
	ans := make([]*TreeNode, len(datas))
	for i := range datas {
		ans[i] = parseValue(datas[i])
	}
	return ans
}

func parseValue(val string) *TreeNode {
	if val == Null {
		return nil
	}
	num, _:= strconv.Atoi(val)
	return &TreeNode{Val: num}
}


type TreeSerialize interface {
	Deserialize(data string) *TreeNode
	Serialize(root *TreeNode ) string
}

type PreOrderTraverseSerialize int

func (s PreOrderTraverseSerialize) handleDeserialize(nodes *[]*TreeNode) *TreeNode{
	if len(*nodes) == 0 {
		return nil
	}
	r := (*nodes)[0]
	*nodes = (*nodes)[1:]
	if r == nil {
		return nil
	}
	r.Left = s.handleDeserialize(nodes)
	r.Right = s.handleDeserialize(nodes)
	return r
}

func (s PreOrderTraverseSerialize) Deserialize(data string) *TreeNode  {
	nodes := sepString(data)
	return s.handleDeserialize(&nodes)
}

func (s PreOrderTraverseSerialize) handleSerialize(root *TreeNode, b *strings.Builder) {
	if root == nil {
		return
	}
	fmt.Fprintf(b, strconv.Itoa(root.Val))
	s.handleSerialize(root.Left, b)
	s.handleSerialize(root.Right, b)
	return
}

func (s PreOrderTraverseSerialize) Serialize(root *TreeNode ) string {
	b := strings.Builder{}
	s.handleSerialize(root, &b)
	return b.String()
}