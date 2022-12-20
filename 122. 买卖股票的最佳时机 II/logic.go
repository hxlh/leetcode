package main

import "fmt"

func maxProfit(prices []int) int {
	n := len(prices)
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		dp0, dp1 = max(dp0, dp1+prices[i]), max(dp1, dp0-prices[i])
	}
	return dp0
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(maxProfit2([]int{7, 1, 5, 3, 6, 4}))
}

//贪心
//只要今天比昨天股价高就卖出
func maxProfit2(prices []int) int {
	ans:=0
	for i := 1; i < len(prices); i++ {
		if diff:=prices[i]-prices[i-1];diff>0 {
			ans+=diff
		}
	
	}
	return ans
}
