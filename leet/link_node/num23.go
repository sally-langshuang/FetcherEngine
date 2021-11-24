package link_node

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//时间复杂度：O(n)O(n)。
//空间复杂度：O(1)O(1)。
func mergeTwoList(l, r *ListNode) *ListNode {
	head := &ListNode{}
	tail := head
	for l != nil && r != nil {
		if l.Val < r.Val {
			tail.Next = l
			l = l.Next
			tail = tail.Next
		} else {
			tail.Next = r
			r = r.Next
			tail = tail.Next
		}
	}
	if l != nil {
		tail.Next = l
	} else {
		tail.Next = r
	}
	return head.Next
}

func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	left := mergeKLists(lists[:mid])
	right  := mergeKLists(lists[mid:])
	return mergeTwoList(left, right)
}

type priorityQueue struct {
	nodes []*ListNode
	higher func(n1, n2 *ListNode) bool
}

func initPriorityQueue(nodes []*ListNode) priorityQueue {
	q := priorityQueue{
		nodes: []*ListNode{nil},
		// 默认小堆
		higher: func(n1, n2 *ListNode) bool {
			return n1.Val <= n2.Val
		},
	}

	for i := range nodes {
		q.tailPush(nodes[i])
	}
	return q
}

func (q priorityQueue) String() string {
	b := strings.Builder{}
	fmt.Fprintf(&b, "queue->[")
	if !q.isEmpty() {
		for i:=1; i < len(q.nodes);i++ {
			fmt.Fprintf(&b, strconv.Itoa(q.nodes[i].Val))
			if i != len(q.nodes) - 1 {
				fmt.Fprintf(&b, ", ")
			}
		}
	}
	fmt.Fprintf(&b, "]")
	return b.String()
}

func (q *priorityQueue) isEmpty() bool {
	return len(q.nodes) == 1
}

func (q *priorityQueue) peak() *ListNode {
	if q.isEmpty() {
		return nil
	}
	return q.nodes[1]
}

func (q *priorityQueue) fall(idx int) {
	for i := idx; i<len(q.nodes); {
		j := i
		if 2 * i < len(q.nodes) && q.higher(q.nodes[2 * i], q.nodes[j]) {
			j = 2*i
		}
		if 2 * i + 1 < len(q.nodes) && q.higher(q.nodes[2 * i + 1], q.nodes[j]) {
			j = 2 * i + 1
		}
		if j == i {
			break
		} else {
			temp := q.nodes[i]
			q.nodes[i] = q.nodes[j]
			q.nodes[j] = temp
			i = j
		}
	}
}

func (q *priorityQueue) tailRaise()  {
	for i := len(q.nodes) - 1; i > 1 && q.higher(q.nodes[i], q.nodes[i / 2]); i /= 2 {
		t := q.nodes[i / 2]
		q.nodes[i / 2] = q.nodes[i]
		q.nodes[i] = t
	}
}

func (q *priorityQueue)	tailPush(n *ListNode){
	if n == nil {
		return
	}
	q.nodes = append(q.nodes, n)
	q.tailRaise()
}

func (q *priorityQueue)	pop() *ListNode{
	if q.isEmpty() {
		return nil
	}
	val := q.nodes[1]
	if len(q.nodes) > 2 {
		q.nodes[1] = q.nodes[len(q.nodes)-1]
	}
	q.nodes = q.nodes[:len(q.nodes) - 1]
	q.fall(1)
	return val
}

func mergeKLists(lists []*ListNode) *ListNode {
	q := initPriorityQueue(lists)
	head := &ListNode{}
	//fmt.Println(q)
	for tail := head;!q.isEmpty(); tail = tail.Next {
		if q.peak().Next == nil {
			tail.Next = q.pop()
		} else {
			n := q.peak()
			tail.Next = n
			q.nodes[1] = n.Next
			q.fall(1)

		}
		//fmt.Println(q)
	}
	return head.Next

}