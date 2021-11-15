package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	//中序遍历 inorder = [9,3,15,20,7]
	//后序遍历 postorder = [9,15,7,20,3]
	if len(postorder) == 0 {
		return nil
	}
	root := TreeNode{Val: postorder[len(postorder) - 1]}
	i := 0
	for ; i<len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}
	root.Left = buildTree(inorder[:i], postorder[:i])
	root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder) - 1])
	return &root

}