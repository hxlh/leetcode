package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}
//快慢指针
func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head

	if head == nil || head.Next == nil {
		return nil
	}
	if head.Next.Next == head {
		return head
	}

	find := false
	for {
		if fast.Next==nil {
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
		if fast == nil {
			break
		} else if slow == fast {
			find = true
			break
		}
	}
	if !find {
		return nil
	}
	fast = head
	for {
		slow = slow.Next
		fast = fast.Next
		if slow == fast {
			break
		}
	}

	return fast
}
