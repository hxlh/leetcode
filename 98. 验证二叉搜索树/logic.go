package main

import (
	"math"
	"net/http"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}
	lastVal := math.MinInt64
	nodes := make([]*TreeNode, 0)
	cur := root
	for len(nodes) != 0 || cur != nil {
		for cur != nil {
			nodes = append(nodes, cur)
			cur = cur.Left
		}
		top := nodes[len(nodes)-1]
		nodes = nodes[0 : len(nodes)-1]
		if top.Val > lastVal {
			lastVal = top.Val
		} else {
			return false
		}
		cur = top.Right
	}
	return true
}

func main() {

}
