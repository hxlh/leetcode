package main

import "fmt"

func main() {  
	fmt.Println(findMin([]int{4,5,6,1,2,3}))
}


func findMin(nums []int) int {
	if len(nums)==1{
		return nums[0]
	}	
	l, r := 0, len(nums)-1
	min:=nums[0]
	for l < r {
    	mid := l + (r-l)/2
		if nums[mid] == nums[l] && nums[mid] == nums[r]{
			if min>nums[mid]{
				min=nums[mid]
			}
			l++
			r--  
		}else if nums[mid] >= nums[l] {
			if min>nums[l]{
				min=nums[l]
			}
			l=mid+1
		} else if nums[mid] <= nums[r] {
			if min>nums[mid]{
				min=nums[mid]
			}
			r=mid-1
		}
	}

	if l==r && min>nums[l]{
		min=nums[l]
	}

	return min
}