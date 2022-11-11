package main

import "sort"

func main() {
	combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
}

// 思路：与全排列Ⅱ类似，
//	去重关键是：
//（1）排序防止类似1->7和7->1之类的重复，
//（2）同一层的去重，即visited起的作用，防止类似[1,1,2,3],重复两次1的查找
func combinationSum2(candidates []int, target int) [][]int {
	path := make([]int, 0)
	ans := make([][]int, 0)
	sort.Ints(candidates)
	ans, _ = dfs(ans, path,candidates, 0, target)
	return ans
}

func dfs(ans [][]int, path []int,candidates []int, pos int, target int) ([][]int, []int) {
	if target== 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		ans = append(ans, tmp)
		
		return ans, path
	}
	
	visited:=make(map[int]bool, 0)
		continue
		}
		visited[candidates[i]]=true
		if candidates[i] > target {
			return ans, path
		}
		path = append(path, candidates[i])
		ans, path = dfs(ans, path,candidates, i+1, target-candidates[i])
	for i := pos; i < len(candidates); i++ {
		if visited[candidates[i]]{
			path = path[:len(path)-1]
	}
	
	return ans, path
}
