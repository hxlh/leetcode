package main

import (
	"strconv"
	"strings"
)

func main() {
	binaryTreePaths(BuildTree([]int{1, 2, 3, null, 5}))
}

const null int = 0

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(nums []int) *TreeNode {
	root := &TreeNode{
		Val: nums[0],
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	i := 1
	for len(queue) > 0 && i < len(nums) {
		head := queue[0]
		queue = queue[1:]
		if nums[i] != 0 {
			head.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, head.Left)
		}
		if nums[i+1] != 0 {
			head.Right = &TreeNode{Val: nums[i+1]}
			queue = append(queue, head.Right)
		}
		i += 2
	}

	return root
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 思路: dfs(非递归)
func binaryTreePaths(root *TreeNode) []string {
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	path := make([]string, 0)
	tmpPath := make([]string, 0)
	visited := make(map[*TreeNode]bool)

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		if visited[cur] {
			tmpPath = tmpPath[:len(tmpPath)-1]
			stack = stack[:len(stack)-1]
			delete(visited, cur)
			continue
		}
		tmpPath = append(tmpPath, strconv.Itoa(cur.Val))
		visited[cur] = true
		if cur.Left == nil && cur.Right == nil {
			p := strings.Builder{}
			p.WriteString(tmpPath[0])
			for i := 1; i < len(tmpPath); i++ {
				p.WriteString("->" + tmpPath[i])
			}
			tmpPath = tmpPath[:len(tmpPath)-1]
			path = append(path, p.String())

			stack = stack[:len(stack)-1]
			delete(visited, cur)
			continue
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}

	return path
}
