package main

import "log"

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]] = 1
	}
	ret := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		v, _ := m[nums2[i]]
		if v == 1 {
			ret = append(ret, nums2[i])
			m[nums2[i]] = 2
		}
	}

	return ret
}


func main() {
	log.Println(intersect([]int{1,2,2,1},[]int{2,2}))
}