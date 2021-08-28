package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	var first, second *ListNode = dummy, dummy

	// Advances first pointer so that the gap between first and second is n nodes apart
	// adjust for the dummy
	for i := 1; i <= n+1; i++ {
		first = first.Next
	}

	// move both maintaining the pointer gap
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// use the pointer that has fallen behind
	// delete the nth from the back
	second.Next = second.Next.Next
	return dummy.Next
}
