package link_node
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList2(head *ListNode) *ListNode {
	var newHead *ListNode
	for cur := head; cur != nil; {
		remain := cur.Next
		cur.Next = newHead
		newHead = cur
		cur = remain
	}
	return newHead
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}