package lcp

import (
	"fmt"
	. "github.com/sally-langshuang/FetcherEngine/leet/tree"
)

func maxValueDp(root *TreeNode, k int) (int, []int) {
	if root == nil {
		return 0, []int{}
	}
	if k == 0 {
		return 0, []int{}
	}
	lkval, lDeep := maxValueDp(root.Left, k)
	rkval, rDeep := maxValueDp(root.Right, k)
	sumVal := lkval + rkval + root.Val
	rootDeep := mergeArr(lDeep, root.Val, rDeep)
	sumVal, rootDeep = keepKDeep(sumVal, rootDeep, k, len(lDeep))
	root.Print()
	fmt.Printf("==>root:%d %v", sumVal, rootDeep)
	return sumVal, rootDeep
}

func keepKDeep(sum int, arr []int, k int, rootIdx int) (int, []int) {
	// 0 <= i <= min(k, n)   max(i + 1, n-k-1)<= j
	n := len(arr)
	if k <= 0 {
		return 0, []int{}
	}
	if n <= k {
		if rootIdx+1 >= n-rootIdx {
			return sum, arr[:rootIdx+1]
		} else {
			return sum, arr[rootIdx:]
		}
	}
	if 2*k+1 >= n {
		arr = append(arr, make([]int, 2*k+2-n)...)
	}
	minI := k
	minIJ := minI + k + 1
	for i := k; i >= 0; i-- {
		j0 := i + k + 1
		for j := j0; j > i; j-- {
			if arr[j] < arr[j0] || (arr[j] == arr[j0] && j == rootIdx) {
				j0 = j
			}
		}
		if arr[i]+arr[j0] < arr[minI]+arr[minIJ] {
			minI, minIJ = i, j0
		}
	}
	minIJ = min(n, minIJ)
	sum -= arr[minI] + arr[minIJ]
	if minI == rootIdx || minIJ == rootIdx {
		return sum, []int{}
	} else if minI > rootIdx {
		if rootIdx+1 >= minI-rootIdx {
			return sum, arr[:rootIdx]
		} else {
			return sum, arr[rootIdx:minI]
		}
	} else if minIJ > rootIdx {
		if rootIdx-minI >= minIJ-rootIdx {
			return sum, arr[minI+1 : rootIdx]
		} else {
			return sum, arr[rootIdx:minIJ]
		}
	} else {
		if rootIdx-minIJ >= n-rootIdx {
			return sum, arr[minIJ+1 : rootIdx]
		} else {
			return sum, arr[rootIdx:]
		}
	}
}

func sort(a, b int) (int, int) {
	if a <= b {
		return a, b
	}
	return b, a
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func mergeArr(a []int, b int, c []int) []int {
	ans := append(a, b)
	for i := len(c) - 1; i >= 0; i-- {
		ans = append(ans, c[i])
	}
	return ans
}

func maxValue(root *TreeNode, k int) int {
	ans, _ := maxValueDp(root, k)
	return ans
}
