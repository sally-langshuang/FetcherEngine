package link_node

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	return swapKPairs(head, 2)
}

func swapTwo(hp, tp *ListNode)  *ListNode{
	if hp == nil || tp == nil {
		return nil
	}
	h := hp.Next
	if hp.Next != tp {
		ht := hp.Next.Next
		hp.Next = tp.Next
		tp.Next = h
		h.Next = hp.Next.Next
		hp.Next.Next = ht
	} else {
		hp.Next = tp.Next
		tp.Next = hp.Next.Next
		hp.Next.Next = tp
	}
	return h
}
func swapKPairs(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	sentinel := &ListNode{Next: head}
	for hp := sentinel; hp.Next != nil; {
		i := 0
		tp := hp

		for ; i < k- 1 && tp.Next != nil; i++ {
			tp = tp.Next
		}
		if i < k- 1 || tp.Next == nil {
			break
		}

		hp = swapTwo(hp, tp)
	}
	return sentinel.Next
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	return swapKPairs(head, k)
}
