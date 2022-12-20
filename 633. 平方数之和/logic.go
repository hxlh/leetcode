package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(judgeSquareSum(88888888))
}

func judgeSquareSum(c int) bool {
	var start int64 = 0
	var end int64 = int64(math.Sqrt(float64(c)))
	for start <= end {
		total := start*start + end*end
		if total > int64(c) {
			end--
		} else if total < int64(c) {
			start++
		} else {
			return true
		}
	}

	return false
}
