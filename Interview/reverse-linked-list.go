package main

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	post := head.Next
	head.Next = nil
	for post != nil {
		dummy.Next = post
		post = post.Next
		dummy.Next.Next = head
		head = dummy.Next
	}
	return dummy.Next
}
