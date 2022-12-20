package main

import (
	"math/rand"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {

	nums1 := make([]int, 20)
	ans := make([]int, 20)
	for i := 0; i < 20; i++ {
		//build nums
		for j := 0; j < 20; j++ {
			nums1[j] = rand.Intn(9999)
		}
		copy(ans, nums1)

		sort.Ints(nums1)
		QuickSort(ans, 0, len(ans))
		for k := 0; k < 20; k++ {
			if ans[k] != nums1[k] {
				t.Fatalf("failed\n")
			}
		}
	}
}

func TestMergeSort(t *testing.T) {

	nums1 := make([]int, 20)
	ans := make([]int, 20)
	tmp := make([]int, 20)
	for i := 0; i < 20; i++ {
		//build nums
		for j := 0; j < 20; j++ {
			nums1[j] = rand.Intn(9999)
		}
		copy(ans, nums1)

		sort.Ints(nums1)

		MergeSort(ans, 0, len(ans), tmp)
		for k := 0; k < 20; k++ {
			if ans[k] != nums1[k] {
				t.Fatalf("failed\n")
			}
		}
	}
}

func TestInsertionSort(t *testing.T) {

	nums1 := make([]int, 20)
	ans := make([]int, 20)

	for i := 0; i < 20; i++ {
		//build nums
		for j := 0; j < 20; j++ {
			nums1[j] = rand.Intn(9999)
		}
		copy(ans, nums1)

		sort.Ints(nums1)

		InsertionSort(ans)
		for k := 0; k < 20; k++ {
			if ans[k] != nums1[k] {
				t.Fatalf("failed\n")
			}
		}
	}
}

func TestBubbleSort(t *testing.T) {
	nums1 := make([]int, 20)
	ans := make([]int, 20)

	for i := 0; i < 20; i++ {
		//build nums
		for j := 0; j < 20; j++ {
			nums1[j] = rand.Intn(9999)
		}
		copy(ans, nums1)

		sort.Ints(nums1)

		BubbleSort(ans)
		for k := 0; k < 20; k++ {
			if ans[k] != nums1[k] {
				t.Fatalf("failed\n")
			}
		}
	}
}

func TestSelectionSort(t *testing.T) {
	nums1 := make([]int, 20)
	ans := make([]int, 20)

	for i := 0; i < 20; i++ {
		//build nums
		for j := 0; j < 20; j++ {
			nums1[j] = rand.Intn(9999)
		}
		copy(ans, nums1)

		sort.Ints(nums1)

		SelectionSort(ans)
		for k := 0; k < 20; k++ {
			if ans[k] != nums1[k] {
				t.Fatalf("failed\n")
			}
		}
	}
}

