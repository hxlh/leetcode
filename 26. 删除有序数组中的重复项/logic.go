package main

import "log"

func removeDuplicates(nums []int) int {
	if len(nums)==0 {
		return 0
	}
	if len(nums)==1{
		return 1
	}
	pre := 1
	next := 1
	for {
		if next == len(nums) {
			break
		}
		if nums[next] != nums[next-1] {
			nums[pre] = nums[next]
			pre++
		}

		next++
	}
	return pre
}

func main() {
	nums := []int{1}
	n := removeDuplicates(nums)
	log.Println(nums[:n])
}