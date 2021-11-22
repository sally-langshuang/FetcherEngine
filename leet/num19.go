package leet

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	sentinel := &ListNode{Next: head}
	p, distance := sentinel, 1
	for i := head; i.Next != nil; i = i.Next {
		if distance < n {
			distance++
		}else {
			p = p.Next
		}
	}
	if distance == n && p.Next != nil {
		p.Next = p.Next.Next
	}
	return sentinel.Next
}
