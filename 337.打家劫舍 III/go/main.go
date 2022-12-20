package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
树形dp
采用后序遍历树
*/
func rob(root *TreeNode) int {
	res:=dfs(root)
	return max(res[0],res[1])
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}

// 返回值是dp数组,dp[0]为偷当前节点，dp[1]为不偷，所得到的最大收益
func dfs(node* TreeNode) []int{
	if node==nil{
		return []int{0,0}
	}
	
	left:=dfs(node.Left)
	right:=dfs(node.Right)
	// 偷
	val1:=node.Val+left[1]+right[1]
	// 不偷
	val2:=max(left[0],left[1])+max(right[0],right[1])

	return []int{val1,val2}
}

func main() {

}
