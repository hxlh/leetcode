package main

import "fmt"

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	var ansPtr *[][]int = &ans
	backtrace(nums, 0, &ansPtr)
	return ans
}

func swap(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}

func copyFrom[T any](src []T) []T {
	dst:=make([]T,len(src))
	copy(dst,src)
	return dst
}

func backtrace(nums []int, depth int, ans **[][]int) {
	if depth == len(nums)-1 {
		tmp:=make([]int,len(nums))
		copy(tmp,nums)
		**ans = append(**ans,tmp)
		return
	}

	for i := depth; i < len(nums); i++ {
		swap(nums, i, depth)
		backtrace(nums, depth+1, ans)
		swap(nums, i, depth)
	}
}

func main() {
	fmt.Printf("%v", permute([]int{1, 2, 3}))
}
