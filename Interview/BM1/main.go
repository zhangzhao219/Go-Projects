package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 创建链表
func (head *ListNode) CreateList(numlist []int) {
	p := head
	for i := 0; i < len(numlist); i++ {
		p.Next = &ListNode{
			Val:  numlist[i],
			Next: nil,
		}
		p = p.Next
	}
}

// 带有头结点的链表翻转
func (head *ListNode) ReverseListWithHead() {
	dummy := head
	head = head.Next
	if head.Next == nil {
		return
	}
	var p *ListNode
	for head != nil {
		q := head.Next
		head.Next = p
		p = head
		head = q
	}
	dummy.Next = p
}

// 不带有头结点的链表翻转
func ReverseListWithOutHead(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var p *ListNode
	for head != nil {
		q := head.Next
		head.Next = p
		p = head
		head = q
	}
	return p
}

// 链表指定区间内翻转
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	cur := head
	for i := 1; i < m; i++ {
		pre = cur
		cur = cur.Next
	}
	for i := m; i < n; i++ {
		q := cur.Next
		cur.Next = q.Next
		q.Next = pre.Next
		pre.Next = q
	}
	return dummy.Next
}

// 链表K个结点一组翻转
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	count := 1
	for head != nil {
		if count == k {
			tail := head.Next
			front := pre.Next
			t := front
			p := front.Next
			for i := 0; i < k-1; i++ {
				q := p.Next
				p.Next = t
				t = p
				p = q
			}
			pre.Next = t
			front.Next = tail
			pre = front
			head = tail
			count = 1
		}
		if head == nil {
			break
		}
		head = head.Next
		count += 1
	}
	return dummy.Next
}

// 打印链表
func (head *ListNode) Print() {
	head = head.Next
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func main() {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	numlist := []int{1, 2, 3, 4, 5}
	head.CreateList(numlist)
	// head.ReverseListWithHead()
	// head = ReverseListWithOutHead(head)
	// head = reverseBetween(head, 2, 4)
	head = reverseKGroup(head, 1)
	head.Print()
}
