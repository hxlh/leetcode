package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)
	return root
}

func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func main() {
	sortedArrayToBST([]int{-10,-3,0,5,9})
}
