package link_node

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}


func (l *ListNode) String() (s string) {
	for temp := l; temp != nil; temp = temp.Next{
		s += fmt.Sprintf("%d", temp.Val)
		if temp.Next != nil {
			s += " -> "
		}
	}
	s = fmt.Sprintf("ListNode[%s]", s)
	return
}

func (l *ListNode) Equal(l2 *ListNode) (bool, error) {
	if l == nil || l2 == nil {
		return false, fmt.Errorf("should not nil")
	}
	t1, t2 := l, l2;
	for ; t1 != nil && t2 != nil; t1, t2 = t1.Next, t2.Next{
		if t1.Val != t2.Val {
			fmt.Println()
			return false, fmt.Errorf("element not equal")
		}
	}
	if t1 != nil || t2 != nil {
		return false, fmt.Errorf("len not equal")
	}
	return true, nil
}

func initListNode(nums []int)  (res *ListNode){
	head := &ListNode{}
	tail := head
	for i := range nums {
		tail.Next = &ListNode{Val: nums[i]}
		tail = tail.Next
	}
	return head.Next
}
