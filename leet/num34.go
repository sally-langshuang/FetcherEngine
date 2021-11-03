package leet


// Definition for a binary tree node.
type Tree struct {
    Val int
    Left *Tree
    Right *Tree
}

func max(vals ...int) (res int) {
	for i := 0; i < len(vals); i++ {
		if vals[i] > res {
			res = vals[i]
		}
	}
	return
}


func dp(root *TreeNode, k, t int) (res int) {
	for j:= 0; j < t; j++ {
		val := root.Val + maxValue(root.Left, t) + maxValue(root.Right, k - t)
		if val > res {
			res = val
		}
	}
	return
}


func maxValue(root *TreeNode, k int) (res int){
	if k == 0 || root == nil {
		return
	}

	res = dp(root, k, 0)
	for i:= 1; i <= k; i++ {
		val := dp(root, k, i) + root.Val
		if val > res {
			res = val
		}
	}
	return

}