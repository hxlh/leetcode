package main

import (
	"fmt"
)

func main() {
	fmt.Println(singleNonDuplicate([]int{0,1,1}))
}

//mid为奇数时,左边有[0,mid-1]为奇数，如果nums[mid]==nums[mid-1],则[0,mid-2]为偶数,则target在[mid+1,len(nums)-1]
//mid为偶数时,左边有[0,mid-1]为偶数,如果nums[mid]==nums[mid-1],则[0,mid-2]为奇数，则target在[0,mid-2]中
func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if mid%2 == 0 {
			//偶
			if nums[mid] == nums[mid-1] {
				r = mid
			} else if nums[mid] == nums[mid+1] {
				l = mid +1
			} else {
				return nums[mid]
			}

		} else {
			if nums[mid] == nums[mid-1] {
				l = mid+1
			} else if nums[mid] == nums[mid+1] {
				r = mid-1
			} else {
				return nums[mid]
			}
		}
	}
	if l == r{
		return nums[l]
	}

	return nums[l-1]
}
