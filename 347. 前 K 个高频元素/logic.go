package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(topKFrequent([]int{4, 1, -1, 2, -1, 2, 3}, 2))
}

type intint [][]int

func (a intint) Len() int           { return len(a) }
func (a intint) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a intint) Less(i, j int) bool { return a[i][1] < a[j][1] }

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	tmp := make([][]int,len(m))
	for i := 0; i < len(m); i++ {
		tmp[i]=make([]int, 2)
	}

	i := 0
	for key, v := range m {
		tmp[i][0] = key
		tmp[i][1] = v
		i++
	}

	sort.Sort(intint(tmp))

	ans := make([]int, 0)
	index := len(tmp) - 1
	for kk := k; kk >0; kk-- {
		ans = append(ans, tmp[index][0])
		index--
	}

	return ans
}
