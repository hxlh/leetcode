package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}))
}

type mytype [][]int

func (m mytype) Len() int {
	return len(m)
}
func (m mytype) Less(i, j int) bool {
	if m[i][0] == m[j][0] {
		return m[i][1] < m[j][1]
	}
	return m[i][0] > m[j][0]
}
func (m mytype) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func reconstructQueue(people [][]int) [][]int {

	sort.Sort(mytype(people))

	ans := make([][]int, len(people))
	for i := 0; i < len(people); i++ {
		copy(ans[people[i][1]+1:], ans[people[i][1]:])
		ans[people[i][1]] = people[i]
	}
	fmt.Println(ans)
	return ans
}
