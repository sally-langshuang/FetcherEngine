package link_node

import (
	"fmt"
	"testing"
)

func TestReverseListt(t *testing.T)  {
	fmt.Println(reverseList(initListNode([]int{1, 2, 3,4,5})))
}
func TestExchangeNodes(t *testing.T)  {
	list := initListNode([]int{1, 2, 3, 4})
	changeTwoChild(list.Next, list.Next.Next)
	fmt.Println(list)
}
func TestNum24(t *testing.T)  {
 fmt.Println(reverseKGroup(initListNode([]int{1, 2,3,4,5,6}), 6)) //[4,3,2,1]
}
func TestNum23(t *testing.T)  {
	//[[-8,-7,-7,-5,1,1,3,4],[-2],[-10,-10,-7,0,1,3],[2]]
	l := initListNode([]int{-8,-7})
	r := initListNode([]int{-2})
	r2 := initListNode([]int{-10,-6,0})
	r3:= initListNode([]int{2})
	fmt.Println(mergeKLists2([]*ListNode{l, r, r2, r3}))
}
