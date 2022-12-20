package main 

func containsDuplicate(nums []int) bool {
    m:=make(map[int]int,len(nums))
	for i := 0; i < len(nums); i++ {
	 	v,_:=m[nums[i]]
		if v==0 {
			m[nums[i]]=1
		}else{
			return true
		}
	}
	return false
}

func main(){

}