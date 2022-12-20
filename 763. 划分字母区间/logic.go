package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}

type mytype [][]int

func (m mytype) Len() int {
	return len(m)
}
func (m mytype) Less(i, j int) bool {
	return m[i][0] < m[j][0]
}
func (m mytype) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}

//转换为区间问题,求最多不重叠区间
func partitionLabels(s string) []int {
	//统计各字母最后一次出现的位置
	m := make(map[byte][]int)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = append(m[s[i]], i, i)
		} else {
			m[s[i]][1] = i
		}
	}

	arr := make([][]int, 0)
	for _, v := range m {
		arr = append(arr, v)
	}
	sort.Sort(mytype(arr))

	ret := make([]int, 0)
	prev := arr[0][1]
	start := arr[0][0]
	for i := 1; i < len(arr); i++ {
		//重叠
		if arr[i][0] > prev {
			ret = append(ret, prev-start+1)
			start = arr[i][0]
		}
		prev = max(prev, arr[i][1])
	}
	ret=append(ret,prev-start+1)

	// fmt.Println(arr)
	return ret
}
