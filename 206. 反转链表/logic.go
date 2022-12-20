package main

import (
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	pre := head
	next := head.Next
	count := 0
	for next.Next != nil {
		nn := next.Next
		if count == 0 {
			pre.Next = nil
		}
		next.Next = pre
		pre = next
		next = nn
		count++
	}
	if next.Next == nil && count==0{
		pre.Next = nil
	}
	next.Next = pre

	return next
}

func main() {
	heads := []int{1, 2}
	head := &ListNode{}
	h := head
	for i := 0; i < len(heads); i++ {
		h.Val = heads[i]
		if i != len(heads)-1 {
			h.Next = &ListNode{}
		}

		h = h.Next
	}

	head = reverseList(head)

	for head != nil {
		log.Println(head.Val)
		head = head.Next
	}
}
