package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//[1,2,3,4,null,2,4,null,null,4]  [[2,4],[4]]
func isSame(a, b *TreeNode) bool{
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val && isSame(a.Left, b.Left) && isSame(a.Right, b.Right)
}
func inArr(t *TreeNode, ans []*TreeNode) bool {
	for x := range ans {
		if isSame(t, ans[x]) {
			return true
		}
	}
	return false
}
func f(last []*TreeNode, ans *[]*TreeNode) ([]*TreeNode){
	bord := len(last)
	for x := range last {
		if last[x] != nil{
			if last[x].Left != nil && !inArr(last[x].Left, *ans) {
				last = append(last, last[x].Left)
			}
			if last[x].Right != nil && !inArr(last[x].Right, *ans){
				last = append(last, last[x].Right)
			}
		}
	}
	for x := 0; x < len(last) - 1; x++ {
		for y :=  x + 1; y < len(last); y++{
			if isSame(last[x], last[y]) && !inArr(last[x], *ans) {
				*ans = append(*ans, last[x])
			}
		}
	}
	return last[bord:]
}
func findDuplicateSubtrees(root *TreeNode) (ans []*TreeNode) {
	ans = make([]*TreeNode, 0)
	for cur:= []*TreeNode{root}; len(cur) > 0; cur= f(cur, &ans){
	}
	return
}
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

func findDuplicateSubtrees0(root *TreeNode) (res []*TreeNode) {
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
