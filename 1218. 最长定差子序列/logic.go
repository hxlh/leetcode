package main

import "log"

func longestSubsequence(arr []int, difference int) int {
	dp := make(map[int]int)
	ans := 0
	for i := 0; i < len(arr); i++ {
		dp[arr[i]]=dp[arr[i]-difference]+1


		if dp[arr[i]] > ans {
			ans = dp[arr[i]]
		}
	}
	return ans
}

func main() {
	log.Println(longestSubsequence([]int{1, 2, 3, 4}, 1))
}
