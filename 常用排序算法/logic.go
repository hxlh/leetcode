package main

func main(){

}

//选择一个数（一般是第一个）分组排序，再各分组排序
func QuickSort(nums []int,l,r int)  {
	if l+1>=r {
		return
	}

	first:=l
	last:=r-1
	key:=nums[first]

	for first<last{
		for first<last && nums[last]>=key {
			last--
		}
		nums[first]=nums[last]

		for first<last && nums[first]<=key{
			first++
		}
		nums[last]=nums[first]
	}
	nums[first]=key
	//left
	QuickSort(nums,l,first)
	//right
	QuickSort(nums,first+1,r)
}

//类似快排分组但不排序，分割后合并
func MergeSort(nums[]int,l,r int,tmp[]int){
	if l+1>=r {
		return
	}
	mid:=l+(r-l)/2
	MergeSort(nums,l,mid,tmp)
	MergeSort(nums,mid,r,tmp)
	p,index:=l,l
	q:=mid
	for p<mid || q<r{
		if q>=r || (p<mid && nums[p]<=nums[q]) {
			tmp[index]=nums[p]
			index++
			p++
		}else {
			tmp[index]=nums[q]
			index++
			q++
		}
	}

	for i := l; i < r; i++ {
		nums[i] = tmp[i]
	}
}

//每次插入一项，该项在排好序的数组中选择一个合适位置
func InsertionSort(nums []int)  {
	for i := 0; i < len(nums); i++ {
		for j := i; j>0 && nums[j]<nums[j-1]; j-- {
			nums[j],nums[j-1]=nums[j-1],nums[j]
		}
	}
}

//每次将最大（小）排到最后
func BubbleSort(nums[]int)  {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j]<nums[j-1] {
				nums[j],nums[j-1]=nums[j-1],nums[j]
			}
		}
	}
}

//每次从i~n之间选择最小（大）的一项跟i交换
func SelectionSort(nums []int)  {
	for i := 0; i < len(nums); i++ {
		min:=i
		
		for j := i+1; j < len(nums); j++ {
			if nums[j]<nums[min] {
				min=j
			}
		}
		nums[i],nums[min]=nums[min],nums[i]
	}
}