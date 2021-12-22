package lcp

import (
	. "github.com/sally-langshuang/FetcherEngine/leet/tree"
)

// x为最多与root相连的个数 k为最多相连的个数
func dp(root *TreeNode, x, k int, f *map[*TreeNode]map[int]int) int {
	if root == nil || k == 0 {
		return 0
	}
	if r, ok := (*f)[root]; ok {
		if _, ok := r[x]; ok {
			return (*f)[root][x]
		}
	} else {
		(*f)[root] = make(map[int]int, 0)
	}
	ans := dp(root.Left, k, k, f) + dp(root.Right, k, k, f)
	for i := 0; i <= x-1; i++ {
		j := x - 1 - i
		if a := dp(root.Left, i, k, f) + root.Val + dp(root.Right, j, k, f); a > ans {
			ans = a
		}
	}
	(*f)[root][x] = ans
	return (*f)[root][x]
}

func maxValue(root *TreeNode, k int) int {
	f := map[*TreeNode]map[int]int{}
	ans := dp(root, k, k, &f)
	return ans
}
