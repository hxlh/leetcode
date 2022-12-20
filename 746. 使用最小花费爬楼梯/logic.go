package main

func min(a int,b int) int {
	if a<b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	mincost:=make([]int,len(cost)+1)
	for i := 2; i <= len(cost); i++ {
		mincost[i]=min(mincost[i-1]+cost[i-1],mincost[i-2]+cost[i-2])
	}
	return mincost[len(mincost)-1]
}

func main()  {
	
}