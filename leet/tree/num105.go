package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//func buildTree(preorder []int, inorder []int) (root *TreeNode) {
//	if len(preorder) == 0 {
//		return
//	}
//	root = &TreeNode{Val: preorder[0]}
//	if len(preorder) == 1 {
//		return
//	}
//	var i int
//	for ; i < len(inorder);i++ {
//		if inorder[i] == preorder[0] {
//			break
//		}
//	}
//	root.Left = buildTree(preorder[1: 1+i], inorder[0: i])
//	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
//	return
//}
