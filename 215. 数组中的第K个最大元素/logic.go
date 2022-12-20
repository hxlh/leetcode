package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findKthLargest2([]int{2,1,3},2))
}

func  findKthLargest2(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

//快速选择：快排优化区间选择
func findKthLargest(nums []int, k int) int {
	l:=0
	r:=len(nums)-1
	target:=k-1
	mid:=0
	
	for l<=r{
		mid=quickSelect(nums,l,r)
		if mid==target {
			return nums[mid]
		}else if mid >target{
			r=mid-1
		}else {
			l=mid+1
		}
	}

	return nums[mid]
}

func quickSelect(nums[]int,l,r int) int {
	
	start:=l
	last:=r

	key:=nums[l]
	for start<last{
		for start<last && nums[last]<=key{
			last--
		}
		nums[start]=nums[last]

		for start<last && nums[start]>=key{
			start++
		}
		nums[last]=nums[start]
	}
	nums[start]=key

	return start
}

