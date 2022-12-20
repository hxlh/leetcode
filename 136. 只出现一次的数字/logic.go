package main

func singleNumber(nums []int) int {
	total:=0
    for i := 0; i < len(nums); i++ {
		total^=nums[i]
	}
	return total
}

func main()  {
	
}