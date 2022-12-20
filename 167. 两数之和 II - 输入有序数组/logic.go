package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
}

func twoSum(numbers []int, target int) []int {
	ans:=make([]int,2)
	start:=0
	end:=len(numbers)-1

	for ; start!=end; {
		if numbers[start]+numbers[end]>target {
			end--
		}else if numbers[start]+numbers[end]<target{
			start++
		}else {
			break
		}
	}

	ans[0]=start+1
	ans[1]=end+1
	return ans
}