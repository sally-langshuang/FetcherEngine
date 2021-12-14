package graph

import "fmt"

type Queue struct {
	head, tail *Node
}

func (q *Queue) append(v *Node)  {
	q.tail.head.tail = v
	v.tail = q.tail
}

type Node struct {
	id int
	leftRefs, rightDeepth int
	nexts []*Node
	head, tail *Node
}

func initQueue(n int, relations [][]int)  *Queue {
	head, tail := &Node{}, &Node{}
	nodes := make([]*Node, n + 2)
	nodes[0], nodes[n + 1] = head, tail
	for i:=1; i <= n; i++ {
		nodes[i] = &Node{id: i, head: nodes[i - 1]}
		nodes[i - 1].tail = nodes[i]
	}
	for i:=n; i >= 1; i-- {
		nodes[i].tail = nodes[i + 1]
		nodes[i + 1].head = nodes[i]
	}
	q := &Queue{head: head, tail: tail}
	for x := range relations{
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
func minNumberOfSemesters(n int, relations [][]int, k int) int {
	q := initQueue(n, relations)
	fmt.Println(q)
	return 0
}