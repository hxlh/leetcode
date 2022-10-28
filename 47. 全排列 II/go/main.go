package main

func main() {
	permuteUnique([]int{0, 1, 0, 0, 9})
}

func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)
	ans = dfs(ans,0, nums)
	return ans
}

// 思路：已使用过的数字直接跳过
func dfs(ans [][]int,level int, nums []int) [][]int {
	if level == len(nums)-1 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		ans = append(ans, tmp)
		return ans
	}

	visited:=make(map[int]struct{})
	
	for i := level; i < len(nums); i++ {
		if _,ok:=visited[nums[i]];ok{
			continue
		}
		visited[nums[i]]=struct{}{}

		nums[i], nums[level] = nums[level], nums[i]
		ans = dfs(ans,level+1, nums)
		nums[i], nums[level] = nums[level], nums[i]
	}
	return ans
}
