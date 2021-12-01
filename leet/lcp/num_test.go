package lcp

import (
	"fmt"
	"github.com/sally-langshuang/FetcherEngine/leet/tree"
	"testing"
)

func TestNum34(t *testing.T)  {
	//root = [5,2,3,4], k = 2
	//+-----+
	//|   5 *|
	//|  ┌┴┐|
	//|  2 3*|
	//|┌─┘  |
	//|4*   |
	//+-----+
	//输出：12
	//===========
	//root = [4,1,3,9,null,null,2], k = 2
	//+-------+
	//|   4*   |
	//|  ┌┴┐  |
	//|  1 3*  |
	//|┌─┘ └─┐|
	//|9*     2|
	//+-------+
	//输出：16
	datas := []struct{
		tree *tree.TreeNode
		k, ans int
	}{
		{tree.InitTree("[5,2,3,4]"), 2, 12},
		{tree.InitTree("[4,1,3,9,null,null,2]"), 2, 16},
	}
	for _, x := range datas{
		fmt.Println(maxValue(x.tree, x.k) == x.ans)
	}
}
