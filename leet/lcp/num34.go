package lcp

import (
	. "github.com/sally-langshuang/FetcherEngine/leet/tree"
)

func max(nums []int) int {
	var ans int
	for i := range nums {
		if nums[i] > ans {
			ans = nums[i]
		}
	}
	return ans
}

// k为最多相连的个数 ans[i]表示为最多与root相连的i个数的val和
func dp(root *TreeNode, k int, f *map[*TreeNode][]int) []int {
	if (*f)[root] != nil {
		return (*f)[root]
	}
	ans := make([]int, k+1)
	if root == nil {
		return ans
	}
	ans[0] = max(dp(root.Left, k, f)) + max(dp(root.Right, k, f))
	for i := 1; i <= k; i++ {
		for l := 0; l < i; l++ {
			if v := dp(root.Left, k, f)[l] + root.Val + dp(root.Right, k, f)[i-1-l]; v > ans[i] {
				ans[i] = v
			}
		}
	}
	(*f)[root] = ans
	return ans
}

func maxValue(root *TreeNode, k int) int {
	f := make(map[*TreeNode][]int)
	var ans int
	for _, a := range dp(root, k, &f) {
		if a > ans {
			ans = a
		}
	}
	return ans
}
