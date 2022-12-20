package main

import "log"

func moveZeroes(nums []int) {
	pre := 0
	next := 0
	// count := 0
	for {
		if nums[next] != 0 {
			tmp:=nums[pre]
			nums[pre] = nums[next]
			nums[next]=tmp
			pre++
			// count++
		}

		next++
		if next >= len(nums) {
			break
		}
	}
	// for i := 0; i < len(nums)-count; i++ {
	// 	nums[len(nums)-1-i] = 0
	// }
}

func main() {
	nums:=[]int{ 0,1,0,3,12}
	moveZeroes(nums)
	log.Println(nums)
}