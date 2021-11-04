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
	}
	return
}
const maxK = 10
var dpRes = make(map[*Tree][maxK]int, 0)

func setRes(tree *Tree, t, val int)  {
	res := dpRes[tree]
	res[t] = val
}

func getRes(tree *Tree, idx int) int {
	res := dpRes[tree]
	return res[idx]
}

func maxValue(root *Tree, k int) (res int){
	if k == 0 || root == nil {
		return
	}

	return

}