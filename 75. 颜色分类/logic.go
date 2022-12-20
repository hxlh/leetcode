package main

func main() {
	sortColors([]int{2, 0, 2, 1, 1, 0})
}

func sortColors(nums []int) {
	quickSort(nums, 0, len(nums))
}

func quickSort(nums []int, l, r int) {
	if l+1 >= r {
		return
	}
	start := l
	last := r - 1
	key := nums[start]

	for start < last {
		for start < last && nums[last] >= key {
			last--
		}
		nums[start] = nums[last]
		for start < last && nums[start] <= key {
			start++
		}
		nums[last] = nums[start]
	}
	nums[start] = key
	quickSort(nums, l, start)
	quickSort(nums, start+1, r)
}
