package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func isSymmetric(root *TreeNode) bool {
//     return check(root, root)
// }
func isSymmetric(root *TreeNode) bool {
	q:=make([]*TreeNode,0)
	q = append(q, root)
	q = append(q, root)

	for len(q)>0{
		l:=q[0]
		q=q[1:]
		r:=q[0]
		q=q[1:]
		if l == nil && r == nil {
			continue
		}
		if (l == nil || r == nil) || (l.Val!=r.Val) {
			return false
		}
		q = append(q, l.Left)
		q = append(q, r.Right)

		q = append(q, l.Right)
		q = append(q, r.Left)
	}

	return true
}

// func check(p, q *TreeNode) bool {
//     if p == nil && q == nil {
//         return true
//     }
//     if p == nil || q == nil {
//         return false
//     }
//     return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
// }

func main() {

}
