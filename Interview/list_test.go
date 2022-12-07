package main

import (
	"testing"
)

func createlist(valData []int) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for _, v := range valData {
		node := &ListNode{Val: v}
		p.Next = node
		p = node
	}
	return dummy.Next
}

func printList(head *ListNode) []int {
	var resultList []int
	for head != nil {
		resultList = append(resultList, head.Val)
		head = head.Next
	}
	return resultList
}

func TestRemoveNthFromEnd(t *testing.T) {
	head := createlist([]int{1})
	head = removeNthFromEnd(head, 1)
	result := printList(head)
	t.Log(result)
}

func TestReverseList(t *testing.T) {
	head := createlist([]int{1, 2, 3, 4, 5, 6})
	head = reverseList(head)
	result := printList(head)
	t.Log(result)
}
