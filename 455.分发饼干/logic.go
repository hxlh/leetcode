package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	cnt:=0
	child:=0
	cookie:=0
	for child<len(g) && cookie<len(s){
		if g[child]<=s[cookie]{
			child++
			cnt++
		}
		cookie++
	}

	return cnt
}