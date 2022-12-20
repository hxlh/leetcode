package main

import (
	"fmt"
)

func main() {
	fmt.Println(candy([]int{1,3,4,5,2}))
}

func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}

func candy(ratings []int) int {
	
	l := len(ratings)

	candies := make([]int, l)
	for i := 0; i < l; i++ {
		candies[i] = 1
	}

	for i := 0; i < l-1; i++ {
		if ratings[i]<ratings[i+1] {
			candies[i+1]=candies[i]+1
		}
	}

	for i := l-1; i >0 ; i-- {
		if ratings[i-1]>ratings[i] {
			candies[i-1]=max(candies[i-1],candies[i]+1)
		}
	}

	sum := 0
	for i := 0; i < l; i++ {
		sum += candies[i]
	}
	// fmt.Println(candies)
	return sum
}
