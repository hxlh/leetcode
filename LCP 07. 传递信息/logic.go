package main

func numWays(n int, relation [][]int, k int) int {
	dp:=make([][]int,k+1)
	for i := 0; i < k+1; i++ {
		dp[i]=make([]int, n)
	}
	dp[0][0]=1
	for i := 0; i < k; i++ {
		for j := 0; j < len(relation); j++ {
			src:=relation[j][0]
			dst:=relation[j][1]
			dp[i+1][dst]+=dp[i][src]
		}
	}
	return dp[k][n-1]
}


func main(){

}