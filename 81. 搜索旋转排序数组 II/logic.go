package main

import "fmt"

func main() {
	fmt.Println(search([]int{13,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}, 13))
}

/**
	二分搜索：
	1、每次二分后会出现一边有序一边序，
	2、若target在有序区间则在有序区间搜索
	3、若target不在有序区间则在无序区间搜索
	4、若在无序区间，则下一次又会再次出现情况 1
*/
func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return true
		}
		if nums[l]==nums[mid] && nums[r]==nums[mid] {
			l++
			r--
		}else if nums[l] <= nums[mid] {
			//左边递增
			if nums[l]<=target && target<nums[mid] {
				//target在有序递增区间
				r = mid - 1
			}else {
				//无序
				l = mid + 1
			}
		}else {
			//右边递增
			if nums[mid]<target && target<=nums[r] {
				//target在有序递增区间
				l=mid+1
			}else {
				//无序
				r = mid -1
			}
		}
	}

	return false
}
