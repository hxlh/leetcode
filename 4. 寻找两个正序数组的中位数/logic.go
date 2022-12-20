package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3, 4}, []int{5,6,7,8,9,10}))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func check(k int, nums1 []int, nums2 []int) int {
	index1 := 0
	index2 := 0

	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		
		if index2== len(nums2) {
			return nums1[index1+k-1]
		}
		
		if k==1{
			return min(nums1[index1], nums2[index2])
		}

		half := k / 2
		newIndex1 := min(index1+half, len(nums1))-1
		newIndex2 := min(index2+half, len(nums2))-1
		n1 := nums1[newIndex1]
		n2 := nums2[newIndex2]
		if n1 <= n2 {
			k -= (newIndex1 - index1+1)
			index1 = newIndex1+1
		} else {
			k -= (newIndex2 - index2+1)
			index2 = newIndex2+1
		}

		
	}
}

//主要思路：要找到第 k (k>1) 小的元素,每次淘汰k/2
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	k := (len(nums1) + len(nums2)) / 2
	if (len(nums1) + len(nums2)) % 2 == 0 {
		//偶数
		midIndex1, midIndex2 := k,k+1
		return (float64(check(midIndex1, nums1, nums2)) + float64(check(midIndex2, nums1, nums2))) / 2.0
	}else {
		return float64(check(k+1, nums1, nums2))
	}
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
    totalLength := len(nums1) + len(nums2)
    if totalLength%2 == 1 {
        midIndex := totalLength/2
        return float64(getKthElement(nums1, nums2, midIndex + 1))
    } else {
        midIndex1, midIndex2 := totalLength/2 - 1, totalLength/2
        return float64(getKthElement(nums1, nums2, midIndex1 + 1) + getKthElement(nums1, nums2, midIndex2 + 1)) / 2.0
    }
    return 0
}

func getKthElement(nums1, nums2 []int, k int) int {
    index1, index2 := 0, 0
    for {
        if index1 == len(nums1) {
            return nums2[index2 + k - 1]
        }
        if index2 == len(nums2) {
            return nums1[index1 + k - 1]
        }
        if k == 1 {
            return min(nums1[index1], nums2[index2])
        }
        half := k/2
        newIndex1 := min(index1 + half, len(nums1)) - 1
        newIndex2 := min(index2 + half, len(nums2)) - 1
        pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
        if pivot1 <= pivot2 {
            k -= (newIndex1 - index1 + 1)
            index1 = newIndex1 + 1
        } else {
            k -= (newIndex2 - index2 + 1)
            index2 = newIndex2 + 1
        }
    }
    return 0
}

