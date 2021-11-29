package link_node

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	return reverseKGroup(head, 2)
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

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	sentinel := &ListNode{Next: head}
	for p := sentinel; p != nil;{
		arrP := make([]*ListNode, k)

		i:=0
		for ; i < k && p.Next != nil;  {
			arrP[i] = p
			i++
			p = p.Next
		}
		if i != k {
			break
		}

		for j, iVal := k - 1, arrP[0]; j >= (k+1) / 2; j--{
			changeTwoChild(iVal, arrP[j])
			fmt.Println(sentinel)
			if j == k - 1 {
				if k == 2 {
					p = arrP[j]
				} else {
					p = arrP[j].Next
				}

			}
			iVal = iVal.Next
		}
	}
	return sentinel.Next
}
