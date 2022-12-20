package main

import (
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type TreeNodePcb struct{
	node *TreeNode
	level int
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	q := make([]*TreeNodePcb, 0)
	q = append(q, &TreeNodePcb{
		node: root,
		level: 0,
	})
	last:=&TreeNodePcb{node: root,level: 0}
	ret := make([][]int, 0)
	out := make([]int, 0)
	for len(q) > 0 {

		t := q[0]
		q = q[1:]
		if t.node.Left!=nil {
			q = append(q, &TreeNodePcb{node: t.node.Left,level: t.level+1})
		}
		if t.node.Right!=nil {
			q = append(q, &TreeNodePcb{node: t.node.Right,level: t.level+1})
		}
		if last.level!=t.level {
			ret = append(ret, out)
			out=make([]int, 0)
			last=t
		}
		out = append(out, t.node.Val)
		
	}
	ret = append(ret, out)

	return ret
}

func main() {
	// root := &TreeNode{
	// 	Val: 1,
	// 	Left: &TreeNode{
	// 		Val: 2,
	// 		Left: &TreeNode{
	// 			Val: 4,
	// 		},
	// 	},
	// 	Right: &TreeNode{
	// 		Val: 3,
	// 		Right: &TreeNode{
	// 			Val: 5,
	// 		},
	// 	},
	// }
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	log.Println(levelOrder(root))
}
