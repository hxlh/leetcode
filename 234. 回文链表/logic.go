package main

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	tmpH := head
	count := 0
	for tmpH != nil {
		count++
		tmpH = tmpH.Next
	}

	tmpH = head
	nums := make([]int, 0)
	for i := 0; i < count/2; i++ {
		nums = append(nums, tmpH.Val)
		tmpH = tmpH.Next
	}
	if count%2!=0{
		tmpH=tmpH.Next
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != tmpH.Val {
			return false
		}
		tmpH = tmpH.Next
	}

	return true
}

func main() {
	
}
