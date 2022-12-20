package main

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	ans := &ListNode{}
	tmpAns := ans
	for {
		if (l1 != nil && l2 == nil) {
			tmpAns.Val = l1.Val
			l1 = l1.Next
		} else if (l2 != nil && l1 == nil)  {
			tmpAns.Val = l2.Val
			l2 = l2.Next
		}else if (l1.Val < l2.Val){
			tmpAns.Val = l1.Val
			l1 = l1.Next
		}else if (l2.Val <= l1.Val){
			tmpAns.Val = l2.Val
			l2 = l2.Next
		}

		
		if l1 == nil && l2 == nil {
			break
		}
		tmpAns.Next = &ListNode{}
		tmpAns = tmpAns.Next
	}

	return ans
}

func main() {
	heads := []int{}
	head := &ListNode{}
	h := head
	for i := 0; i < len(heads); i++ {
		h.Val = heads[i]
		if i != len(heads)-1 {
			h.Next = &ListNode{}
		}

		h = h.Next
	}

	heads2 := []int{0}
	head2 := &ListNode{}
	h2 := head2
	for i := 0; i < len(heads2); i++ {
		h2.Val = heads2[i]
		if i != len(heads2)-1 {
			h2.Next = &ListNode{}
		}

		h2 = h2.Next
	}

	head=nil
	ret := mergeTwoLists(head, head2)

	for ret != nil {
		log.Println(ret.Val)
		ret = ret.Next
	}
}
