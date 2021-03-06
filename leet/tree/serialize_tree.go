package tree

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	Null                            = "#"
	Sep                             = ","
	pre  PreOrderTraverseSerialize  = iota
	post PostOrderTraverseSerialize = iota
	level LevelTraverseSerialize = iota
)

func sepString(data string) []*TreeNode {
	datas := strings.Split(strings.TrimFunc(data, func(c rune) bool { return c == '[' || c == ']' || c == ' ' }), Sep)
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
	num, _ := strconv.Atoi(val)
	return &TreeNode{Val: num}
}

type TreeSerialize interface {
	Deserialize(data string) *TreeNode
	Serialize(root *TreeNode) string
}

type PreOrderTraverseSerialize int

func (s PreOrderTraverseSerialize) handleDeserialize(nodes *[]*TreeNode) *TreeNode {
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

func (s PreOrderTraverseSerialize) Deserialize(data string) *TreeNode {
	nodes := sepString(data)
	return s.handleDeserialize(&nodes)
}

func (s PreOrderTraverseSerialize) handleSerialize(root *TreeNode, b *strings.Builder) {
	if root == nil {
		fmt.Fprintf(b, Null)
		fmt.Fprintf(b, Sep)
		return
	}
	fmt.Fprintf(b, strconv.Itoa(root.Val))
	fmt.Fprintf(b, Sep)
	s.handleSerialize(root.Left, b)
	s.handleSerialize(root.Right, b)
	return
}

func (s PreOrderTraverseSerialize) Serialize(root *TreeNode) string {
	b := strings.Builder{}
	s.handleSerialize(root, &b)
	str := b.String()
	if len(str) != 0 {
		str = str[:len(str)-1]
	}
	return str
}

type PostOrderTraverseSerialize int

func (s PostOrderTraverseSerialize) handleDeserialize(nodes *[]*TreeNode) *TreeNode {
	if len(*nodes) == 0 {
		return nil
	}
	root := (*nodes)[len(*nodes)-1]
	*nodes = (*nodes)[:len(*nodes)-1]
	if root != nil {
		root.Right = s.handleDeserialize(nodes)
		root.Left = s.handleDeserialize(nodes)
	}
	return root
}

func (s PostOrderTraverseSerialize) Deserialize(data string) *TreeNode {
	nodes := sepString(data)
	return s.handleDeserialize(&nodes)
}

func (s PostOrderTraverseSerialize) handleSerialize(root *TreeNode, b *strings.Builder) {
	if root == nil {
		fmt.Fprintf(b, Null)
		fmt.Fprintf(b, Sep)
		return
	}
	s.handleSerialize(root.Left, b)
	s.handleSerialize(root.Right, b)
	fmt.Fprintf(b, strconv.Itoa(root.Val))
	fmt.Fprintf(b, Sep)
}

func (s PostOrderTraverseSerialize) Serialize(root *TreeNode) string {
	b := strings.Builder{}
	s.handleSerialize(root, &b)
	str := b.String()
	if len(str) > 0 {
		str = str[:len(str)-1]
	}
	return str
}

type LevelTraverseSerialize int


func (s LevelTraverseSerialize) Deserialize(data string) *TreeNode{
	nodes := sepString(data)
	if len(nodes) == 0 {
		return nil
	}
	root := nodes[0]
	for i, j:= 0, 1; j < len(nodes); i++{
		if nodes[i] == nil {
			continue
		} else {
			nodes[i].Left = nodes[j]
			j++
			if j == len(nodes) {
				break
			}
			nodes[i].Right = nodes[j]
			j++
		}
	}
	return root

}
//func (s LevelTraverseSerialize)
func (s LevelTraverseSerialize) Serialize(root *TreeNode) string{
	b := strings.Builder{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		r := queue[0]
		queue = queue[1:]
		if r == nil {
			fmt.Fprintf(&b, Null)
			fmt.Fprintf(&b, Sep)
			continue
		}
		fmt.Fprintf(&b, strconv.Itoa(r.Val))
		fmt.Fprintf(&b, Sep)
		queue = append(queue, r.Left)
		queue = append(queue, r.Right)
	}
	str := b.String()
	if len(str) > 0 {
		str = str[:len(str) - 1]
	}
	return str
}
