package main

import "log"

func singleNumber(nums []int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, ok := m[nums[i]]
		if ok {
			m[nums[i]]=1
			continue
		} else {
			m[nums[i]]=0
		}
	}
	ret:=make([]int,0)
	for k,v:= range m{
		if v==0 {
			ret = append(ret, k)
		}
	}
	return ret
}

func main() {
	log.Println(singleNumber([]int{0,1}))
}
