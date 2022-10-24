package main

import "fmt"

func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	nums := make([]int, k)
	cnt := 0
	ans = dfs(ans, nums, &cnt, 1, n, k)
	return ans
}

func dfs(ans [][]int, nums []int, cnt *int, pos, n, k int) [][]int {
	if *cnt == k {
		tmp := make([]int, k)
		copy(tmp, nums)
		ans = append(ans, tmp)
		return ans
	}
	for i := pos; i <= n; i++ {
		nums[*cnt] = i
		*cnt++
		ans = dfs(ans, nums, cnt, i+1, n, k)
		*cnt--
	}
	return ans
}

func main() {
	fmt.Printf("%v\n", combine(4, 2))
}
