package main

import "log"

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums := make([]int, 0)
	index1 := 0
	index2 := 0
	for {
		if index2>=n && index1>=m{
			break
		}
		if index1>=m && index2<n {
			nums = append(nums, nums2[index2])
			index2++
			continue
		}
		if index2>=n && index1<m {
			nums = append(nums, nums1[index1])
			index1++
			continue
		}
		if nums1[index1]<=nums2[index2] {
			nums = append(nums, nums1[index1])
			index1++
		}else {
			nums = append(nums, nums2[index2])
			index2++
		}
	}
	copy(nums1,nums)
}


func main() {
	nums1 := []int{0,0,3,0,0,0,0,0,0}
	nums2 := []int{-1,1,1,1,2,3}
	merge2(nums1, 3, nums2, 6)
	log.Println(nums1)
}
func merge2(nums1 []int, m int, nums2 []int, n int) {
	mIndex:=m-1
	nIndex:=n-1
	//每一次插入aIndex都需要+1
	aIndex:=m-1
	if m==0{
		for i := 0; i < n; i++ {
			nums1[i]=nums2[i]
		}
		return
	}
	for nIndex>=0 {
		if nums1[aIndex]<=nums2[nIndex]{
			aIndex++
			nums1[aIndex]=nums2[nIndex]
			nIndex--
		}else if nums1[mIndex]>nums2[nIndex]{
			if mIndex==0 {
				copy(nums1[mIndex+1:],nums1[mIndex:])
				nums1[mIndex]=nums2[nIndex]
				nIndex--
				aIndex++
			}else{
				mIndex--
			}
			
		}else if nums1[mIndex]<=nums2[nIndex]{
			copy(nums1[mIndex+2:],nums1[mIndex+1:])
			nums1[mIndex+1]=nums2[nIndex]
			nIndex--
			aIndex++
		}
	}

}