package tree

type Node = TreeNode

func connectTwo(l, r *Node)  {
	if l == nil {
		return
	}
	connectTwo(l.Left, l.Right)
	if r != nil {
		l.Next = r
		connectTwo(l.Right, r.Left)
		connectTwo(r.Left, r.Right)
	}
}
func connect(root *Node) *Node {
	if root != nil && root.Right != nil {
		connectTwo(root.Left, root.Right)
	}
	return root
}
