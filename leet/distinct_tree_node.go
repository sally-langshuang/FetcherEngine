package leet

import "fmt"

// https://leetcode-cn.com/problems/all-nodes-distance-k-in-binary-tree/solution/er-cha-shu-zhong-suo-you-ju-chi-wei-k-de-qbla/
const null = -1

func childIdx(p int) (int, int) {
	return p*2 + 1, p*2 + 2 // todo非完全二叉
}

func parentIdx(c int) int {
	return (c - 1) / 2
}

type TreeNode struct {
	Val                 int
	Left, Right, Parent *TreeNode
}

func distanceK(root, target *TreeNode, k int) (ans []int) {
	// 从 root 出发 DFS，记录每个结点的父结点
	parents := map[int]*TreeNode{}
	var findParents func(*TreeNode)
	findParents = func(node *TreeNode) {
		if node.Left != nil {
			parents[node.Left.Val] = node
			findParents(node.Left)
		}
		if node.Right != nil {
			parents[node.Right.Val] = node
			findParents(node.Right)
		}
	}
	findParents(root)

	// 从 target 出发 DFS，寻找所有深度为 k 的结点
	var findAns func(*TreeNode, *TreeNode, int)
	findAns = func(node, from *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == k { // 将所有深度为 k 的结点的值计入结果
			ans = append(ans, node.Val)
			return
		}
		if node.Left != from {
			findAns(node.Left, node, depth+1)
		}
		if node.Right != from {
			findAns(node.Right, node, depth+1)
		}
		if parents[node.Val] != from {
			findAns(parents[node.Val], node, depth+1)
		}
	}
	findAns(target, nil, 0)
	return
}

func DistinctK(root []int, target, k int) []int {
	fmt.Println(root)

	var thisLevel, nextLevel = []*TreeNode{{Val: 0}}, []*TreeNode{}
	var targetNode *TreeNode
	//层序遍历
	for l := 0; len(thisLevel) > 0; {
		fmt.Printf("==%d==\n", l)
		for i, _ := range thisLevel {
			idx := thisLevel[i]
			fmt.Printf("%d ", root[idx.Val])
			if root[idx.Val] == target {
				targetNode = idx
			}
			if root[idx.Val] == null {
				continue
			}
			left := idx.Val + (i)*2 + len(thisLevel) - i
			if left < len(root) {
				l := &TreeNode{Val: left}
				if root[left] != null {
					l.Parent = idx
					idx.Left = l
				}
				nextLevel = append(nextLevel, l)
			}
			right := left + 1
			if right < len(root) {
				r := &TreeNode{Val: right}
				if root[right] != null {
					r.Parent = idx
					idx.Right = r
				}
				nextLevel = append(nextLevel, r)
			}

		}
		thisLevel, nextLevel = nextLevel, []*TreeNode{}
		l++
	}
	var f func( *TreeNode, *TreeNode, int) []int
	f = func (t *TreeNode, from *TreeNode, d int) []int{
		if k < 0  || t == nil{
			return []int{}
		}
		if d==0{
			return []int{root[t.Val]}
		}
		res := []int{}
		if t.Left != from {
			res = append(res, f(t.Left, t, d-1)...)
		}
		if t.Right != from {
			res = append(res, f(t.Right, t, d-1)...)
		}
		if t.Parent != from {
			res = append(res, f(t.Parent, t, d-1)...)
		}

		return res
	}
	r:= f(targetNode,nil, k)

	return r
}
