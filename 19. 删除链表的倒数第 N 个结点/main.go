package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head==nil || head.Next==nil{
		return nil
	}
	h:=head
	pre:=h
	pos:=0
	for h!=nil{ 
		h=h.Next
		if pos==n && h==nil{
			pre.Next=pre.Next.Next
			return head
		}else if pos==n{
			pre=pre.Next
			pos--
		}
		pos++
	}
	if pos==n{
		head=head.Next
	}
	return head
}

func main() {
	head := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val: 2,
			Next: nil,
		},
	}
	removeNthFromEnd(head,1)
}
