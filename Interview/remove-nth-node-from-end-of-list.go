package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var dummy = &ListNode{}
	dummy.Next = head

	count := 0
	for head != nil {
		count += 1
		head = head.Next
	}
	head = dummy.Next
	pre := dummy
	count -= n
	for count != 0 {
		pre = head
		count -= 1
		head = head.Next
	}
	pre.Next = head.Next
	return dummy.Next
}
