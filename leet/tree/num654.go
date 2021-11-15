package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxIdx := 0
	for i:= 0; i < len(nums); i++ {
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		}
	}
	root := TreeNode{
		Val: nums[maxIdx],
		Left: constructMaximumBinaryTree(nums[: maxIdx]),
		Right: constructMaximumBinaryTree(nums[maxIdx + 1:]),
	}
	return &root
}

