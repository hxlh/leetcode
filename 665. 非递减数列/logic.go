package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkPossibility([]int{4,2,1}))
}

func checkPossibility(nums []int) bool {

	if len(nums)==1 {
		return true
	}
	times := 1
	//若第一个就大，就说明处于7,5,1...递减这种情况
	if nums[0]>nums[1] {
		times=0
	}
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if times > 0 {
				times--
				if nums[i+1]>=nums[i-1] {
					nums[i]=nums[i+1]
				}else {
					nums[i+1]=nums[i]
				}
			} else {
				return false
			}
		}
	}

	return true
}
