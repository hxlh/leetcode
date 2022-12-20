package main

import "log"

func plusOne(digits []int) []int {
	if digits[0] == 0 {
		return []int{1}
	}
	
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 > 9{
			digits[i] = 0
		} else {
			digits[i] += 1
			break
		}
	}
	ret:=make([]int,0)
	if digits[0]==0 {
		ret = append(ret, 1)
		digits[0]=0
		ret = append(ret, digits...)
	}else {
		ret = append(ret, digits...)
	}
	
	return ret
}

func main() {
	log.Println(plusOne([]int{2}))
}
