package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(eraseOverlapIntervals([][]int{{1,2},{2,3},{3,4},{1,3}}))
}
type mytype [][]int  

func (m mytype) Len() int {
	return len(m)
}

func (m mytype) Less(i, j int) bool {
	return m[i][1] < m[j][1]
}

func (m mytype) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Sort(mytype(intervals))
	fmt.Println(intervals)
	cnt:=0

	prev:=intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0]<prev{
			//移除该区间
			cnt++
		}else {
			prev=intervals[i][1]
		}
	}
	return cnt
}
