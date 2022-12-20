package main

import "log"

func rotate(nums []int, k int) {
	k %= len(nums)
    first := 0
	last := len(nums) - 1
    if len(nums)==1 {
        return
    }

	for first != last  && first<last{
		tmp := nums[first]
		nums[first] = nums[last]
        nums[last]=tmp
		first++
		last--
	}

    first=0
    last=k-1
    for first!=last && first<last{
        tmp := nums[first]
		nums[first] = nums[last]
        nums[last]=tmp
		first++
		last--
    }

    first=k
    last=len(nums)-1
    for first!=last && first<last{
        tmp := nums[first]
		nums[first] = nums[last]
        nums[last]=tmp
		first++
		last--
    }

}

func main() {
	var buf = []int{1, 2, 3, 4, 5, 6, 7}
	rotate(buf, 3)
	log.Println(buf)
}
