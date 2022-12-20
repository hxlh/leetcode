package main

func leastMinutes(n int) int {
	sp:=1
	ans:=1
	for sp<n{
		sp=sp*2
		ans++
	}
	return ans
}

func main()  {
	
}