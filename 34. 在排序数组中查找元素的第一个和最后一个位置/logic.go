package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{1,1}, 1))
}

/**
二分搜索：
查找第1个target：寻找第1个大于等于target 的下标
查找第2个target：寻找第1个大于target 的下标
*/
func searchRange(nums []int, target int) []int {
	if len(nums)==0{
		return []int{-1, -1}
	}

	upper := findUpper(nums, target) - 1
	lower := findLower(nums, target)

	if nums[lower]!=target {
		return []int{-1, -1}
	}

	return []int{lower, upper}
}

func findLower(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		ans := nums[mid]

		if ans >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func findUpper(nums []int, target int) int {
	l, r := 0, len(nums)

	for l < r {
		mid := l + (r-l)/2
		ans := nums[mid]

		if ans > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
