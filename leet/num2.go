package leet

import "fmt"

type ListNode struct {
	Val int
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
	last := res
	for i:= 0; i < len(nums); i = i+ 1{
		temp := &ListNode{}
		if last == nil {
			res = temp
		} else {
			last.Next = temp
		}
		temp.Val = nums[i]
		last = temp
	}
	return res
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	var tail, temp *ListNode
	var carry int
	for ; l1 != nil || l2 != nil; {
		if tail == nil {
			head = &ListNode{}
			temp = head
		} else {
			tail.Next = &ListNode{}
			temp = tail.Next
		}
		if carry != 0 {
			temp.Val += carry
		}
		if l1 != nil {
			temp.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			temp.Val += l2.Val
			l2 = l2.Next
		}
		temp.Val, carry = temp.Val % 10, temp.Val / 10
		tail = temp
	}
	if carry > 0 && tail != nil {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}