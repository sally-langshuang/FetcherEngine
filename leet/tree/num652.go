package tree

import (
	"fmt"
	"strconv"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func f652(allOccur *map[string]int, ans *[]*TreeNode, root *TreeNode) string {
	if root == nil {
		return "#,"
	}
	rootStr := strconv.Itoa(root.Val) + "," + f652(allOccur, ans, root.Left) + f652(allOccur, ans, root.Right)
	if x := (*allOccur)[rootStr]; x <= 1 {
		(*allOccur)[rootStr] += 1
		if x == 1 {
			(*ans) = append(*ans, root)
		}
	}
	return rootStr
}

func findDuplicateSubtrees0(root *TreeNode) []*TreeNode {
	allOccur := make(map[string]int, 0)
	ans := make([]*TreeNode, 0)
	f652(&allOccur, &ans, root)
	return ans
}


func ff652(uids *map[string]string, counts *map[string]int, ans *[]*TreeNode, root *TreeNode, t *int) string{
	if root == nil {
		return "#"
	}
	str := fmt.Sprintf("%s,%s,%s",ff652(uids, counts, ans, root.Left, t), strconv.Itoa(root.Val), ff652(uids, counts, ans, root.Right, t))
	if _, ok := (*uids)[str]; !ok {
		(*uids)[str] = strconv.Itoa(*t)
		*t += 1
	}
	uid := (*uids)[str]
	c := (*counts)[uid]
	(*counts)[uid] = c + 1
	if (*counts)[uid] == 2 {
		*ans = append(*ans, root)
	}
	return uid
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	t := 0
	uids := make(map[string]string, 0)
	counts := make(map[string]int, 0)
	ans := make([]*TreeNode, 0)
	ff652(&uids, &counts, &ans, root, &t)
	return ans
}
