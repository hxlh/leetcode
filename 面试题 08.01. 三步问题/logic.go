package main

func waysToStep(n int) int {
	if n==1 {
		return 1
	}
	if n==2 {
		return 2

	}
	if n==3 {
		return 4
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 4
	for i := 4; i <= n; i++ {
		dp[i]=dp[i-1]+dp[i-2]+dp[i-3]
		dp[i]=dp[i]%1000000007
	}
	return dp[n]
}

func main() {

}
