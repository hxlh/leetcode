package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	h := head
	if h == nil {
		return nil
	}
	prev := head
	next := head.Next

	for next != nil {
		if next.Val == prev.Val {
			prev.Next = next.Next
		} else {
			prev = next
		}
		next = next.Next
	}
	return h
}

func main() {

}

func create_list(num []int) *ListNode {
	node := &ListNode{
		Val:  num[0],
		Next: nil,
	}
	n := node
	for i := 1; i < len(num); i++ {
		n.Next = &ListNode{Val: num[i], Next: nil}
		n = n.Next
	}
	return node
}

func display(node *ListNode) {
	n := node
	for n != nil {
		fmt.Printf("%v ", n.Val)
		n = n.Next
	}
}
