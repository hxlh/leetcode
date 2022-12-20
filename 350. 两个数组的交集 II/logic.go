package main

import "log"

func intersect(nums1 []int, nums2 []int) []int {
	m:=make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		_,ok:=m[nums1[i]]
		if !ok {
			m[nums1[i]]=1
		}else {
			m[nums1[i]]+=1
		}
	}
	ret:=make([]int,0)
	for i := 0; i < len(nums2); i++ {
		v,_:=m[nums2[i]]
		if v>0 {
			m[nums2[i]]-=1
			ret = append(ret, nums2[i])
		}
	}
	
	return ret
}


func main() {
	log.Println(intersect([]int{4,9,5},[]int{9,4,9,8,4}))
}