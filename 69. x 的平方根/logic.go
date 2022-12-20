package main

import "fmt"

func main() {
	fmt.Println(mySqrt(1))
}

//二分法
func mySqrt(x int) int {
	if x==1 || x==0{
		return x
	}
	l,r:=0,x
	for l<=r{
		mid:=l+(r-l)/2
		sqrt:=x/mid
		if sqrt==mid{
			return mid
		}else if sqrt>mid{
			l=mid+1
		}else{
			r=mid-1
		}
	}

	return r
}
