package main

/*
 * Remove Nth Node From End of List
 * Input: head = [1,2,3,4,5], n = 2
 * Output: [1,2,3,5]
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	dummy := &ListNode{Next: head}
	prev := dummy
	for i := 0; i < length-n; i++ {
		prev = prev.Next
	}
	prev.Next = prev.Next.Next

	return dummy.Next
}
