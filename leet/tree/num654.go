package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dp654(nums []int, l, h int) (res *TreeNode) {
	if l == h {
		return nil
	}
	maxNumIdx := l
	for i:= maxNumIdx + 1 ; i < h; i++{
		if nums[i] > nums[maxNumIdx] {
			maxNumIdx = i
		}
	}
	root := TreeNode{Val: nums[maxNumIdx]}
	root.Left = dp654(nums, 0, maxNumIdx)
	root.Right = dp654(nums, maxNumIdx + 1, h)
	return &root
}

func constructMaximumBinaryTree(nums []int) (res *TreeNode) {
	return dp654(nums, 0, len(nums))
}