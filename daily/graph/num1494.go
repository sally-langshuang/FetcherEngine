package graph

import (
	"fmt"
)

type Queue struct {
	head, tail *Node
}

func (q *Queue) append(v *Node) {
	q.tail.head.tail = v
	v.tail = q.tail
}

type Node struct {
	id                    int
	leftRefs, rightDeepth int
	nexts                 []*Node
	head, tail            *Node
}

func initQueue(n int, relations [][]int) *Queue {
	head, tail := &Node{}, &Node{}
	nodes := make([]*Node, n+2)
	nodes[0], nodes[n+1] = head, tail
	for i := 1; i <= n; i++ {
		nodes[i] = &Node{id: i, head: nodes[i-1]}
		nodes[i-1].tail = nodes[i]
	}
	for i := n; i >= 1; i-- {
		nodes[i].tail = nodes[i+1]
		nodes[i+1].head = nodes[i]
	}
	q := &Queue{head: head, tail: tail}
	for x := range relations {
		i, j := nodes[relations[x][0]], nodes[relations[x][1]]
		if j.head != nil {
			j.head.tail = j.tail
			j.tail.head = j.head
		}
		j.leftRefs++
		if len(i.nexts) == 0 {
			i.rightDeepth++
		}
		i.nexts = append(i.nexts, j)
	}
	return q

}
func minNumberOfSemesters0(n int, relations [][]int, k int) int {
	q := initQueue(n, relations)
	fmt.Println(q)
	return 0
}

func getDeepth(i int, deepth []int, nexts [][]int) int {
	if deepth[i] != 0 {
		return deepth[i]
	}
	maxChildDeep := 0
	for j := range nexts[i] {
		if v := getDeepth(nexts[i][j], deepth, nexts); v > maxChildDeep {
			maxChildDeep = v
		}
	}
	return 1 + maxChildDeep
}
func sort(nums []int, less func(i, j int) bool) {

}
func minNumberOfSemesters(n int, relations [][]int, k int) int {
	deepth := make([]int, n)
	degree := make([]int, n)
	nexts := make([][]int, n)

	for i := range nexts {
		nexts[i] = []int{}
	}
	for _, x := range relations {
		degree[x[1]]++
		nexts[x[0]] = append(nexts[x[0]], x[1])
	}
	for i := range deepth {
		getDeepth(i, deepth, nexts)
	}

	less := func(i, j int) bool {
		if degree[i] > degree[j] {
			return false
		}
		return degree[i] < degree[j] || deepth[i] >= deepth[j]
	}
	queue := make([]int, n)
	for i := range queue {
		queue[i] = i
	}
	sort(queue, less)
	ans := 0
	for ; len(queue) > 0; ans++ {
		pops := []int{}
		for i := 0; i < k; i++ {
			num := queue[0]
			if degree[num] != 0 {
				break
			}
			queue = queue[1:]
			pops = append(pops, num)

		}
		for _, num := range pops {
			for _, x := range nexts[num] {
				degree[x]--
			}
		}
	}
	return ans
}
