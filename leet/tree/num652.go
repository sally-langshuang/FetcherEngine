package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func equal(tree1, tree2 *TreeNode) (res bool) {
	if tree1 == nil && tree2 == nil {
		return true
	}
	if tree1.Val != tree2.Val {
		return false
	}
	return equal(tree1.Left, tree2.Left) && equal(tree1.Right, tree2.Right)
}

func in(tree *TreeNode, arr []*TreeNode) bool {
	if tree == nil || len(arr) == 0 {
		return false
	}
	for _, i := range arr {
		if equal(i, tree) {
			return true
		}
	}
	return false
}

func findDuplicateSubtrees(root *TreeNode) (res []*TreeNode) {
	res = make([]*TreeNode, 0)
	if root == nil {
		return
	}
	leftRes := findDuplicateSubtrees(root.Left)
	for _, i:= range leftRes{
		if !in(i, res) {
			res = append(res, i)
		}
	}
	rightRes := findDuplicateSubtrees(root.Right)
	for _, i:= range rightRes{
		if !in(i, res) {
			res = append(res, i)
		}
	}
	return res
}
