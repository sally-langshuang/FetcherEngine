package tree

import "math"

func Height(t *TreeNode) (h int) {
	if t == nil {
		return
	}
	return maxInt(Height(t.Left), Height(t.Right)) + 1
}

func RRRoll(t *TreeNode) *TreeNode {
	root := t.Right
	t.Right = root.Left
	root.Left = t
	return root
}
func LLRoll(t *TreeNode) *TreeNode {
	root := t.Left
	t.Left = root.Right
	root.Right = t
	return root
}
func RLRoll(t *TreeNode) *TreeNode {
	root := t.Right.Left
	t.Right.Left = root.Right
	root.Right = t.Right
	t.Right = root.Left
	root.Left = t
	return root
}

func LRRoll(t *TreeNode) *TreeNode {
	root := t.Left.Right
	t.Left.Right = root.Left
	root.Left =  t.Left
	t.Left = root.Right
	root.Right = t
	return root
}

func Balance(t *TreeNode) *TreeNode {
	// postorder
	if t == nil {
		return t
	}
	t.Left = Balance(t.Left)
	t.Right = Balance(t.Right)
	for math.Abs(float64(Height(t.Left) - Height(t.Right))) > 1{
		if Height(t.Left) > Height(t.Right) {
			if Height(t.Left.Left) > Height(t.Left.Right) {
				t = LLRoll(t)
			}else {
				t = LRRoll(t)
			}
		}else{
			if Height(t.Right.Left) > Height(t.Right.Right) {
				t = RLRoll(t)
			}else {
				t = RRRoll(t)
			}
		}
	}
	return t
}

func InsertNode(val int, t *TreeNode) *TreeNode {
	if t == nil {
		return &TreeNode{Val: val}
	}
	if t.Val >= val {
		t.Left = InsertNode(val, t.Left)
	} else {
		t.Right = InsertNode(val, t.Right)
	}
	return t
}
func InsertNodeAndBalance(val int, t *TreeNode) *TreeNode {
	t = InsertNode(val, t)
	t = Balance(t)
	return t
}
func (t *TreeNode) InsertNode(val int) *TreeNode {
	x := InsertNode(val, t)
	return x
}
func (t *TreeNode) InsertNodes(vals ...int) *TreeNode {
	x := t
	for _, v:= range vals {
		x = InsertNode(v, x)
	}
	return x
}
func (t *TreeNode) InsertNodeAndBalance(val int) *TreeNode {
	x := InsertNodeAndBalance(val, t)
	return x
}

func (t *TreeNode) InsertNodesAndBalance(vals ...int) *TreeNode {
	x := t
	for _, v:= range vals {
		x = InsertNodeAndBalance(v, x)
	}
	return x
}
func InsertNodesAndBalance(vals ...int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	x := &TreeNode{Val: vals[0]}
	for _, v:= range vals {
		x = InsertNodeAndBalance(v, x)
	}
	return x
}
