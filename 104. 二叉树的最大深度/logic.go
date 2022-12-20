package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct{
	Node *TreeNode
	Dp int
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nodes := make([]Node, 0)
	deepth := 0
	nodes = append(nodes, Node{Node: root,Dp: 1})

	for {
		if len(nodes) == 0 {
			break
		}
		r := nodes[0]
		nodes = nodes[1:]
		if r.Node.Left != nil {
			nodes = append(nodes,Node{r.Node.Left,r.Dp+1})
		}
		if r.Node.Right != nil {
			nodes = append(nodes, Node{r.Node.Right,r.Dp+1})
		}
		if  r.Node.Left==nil &&  r.Node.Right==nil {
			if deepth<r.Dp {
				deepth=r.Dp
			}
		}
	}
	return deepth
}

func main() {

}
