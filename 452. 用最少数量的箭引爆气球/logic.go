package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMinArrowShots([][]int{{10,16},{2,8},{1,6},{7,12}}))
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

//找最多重叠区间
func findMinArrowShots(points [][]int) int {
	sort.Sort(mytype(points))
	fmt.Println(points)

	cnt:=0
	prev:=points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > prev {
			prev = points[i][1]
			cnt++
		}
	}
	cnt++

	return cnt
}
