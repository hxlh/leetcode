package main

import "fmt"

func main() {
	fmt.Println(findMinHeightTrees(6, [][]int{
		{3, 0},
		{3, 1},
		{3, 2},
		{3, 4},
		{5, 4},
	}))
}

// 思路：找树的最长路径的中间点,设起点为x,先找离x最远的点y，再找离y最远的点z，若两者path长度不同，则选长度较长者取中间节点
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	// 先构建映射关系
	e := make(map[int][]int, 0)
	for _, edge := range edges {
		e[edge[0]] = append(e[edge[0]], edge[1])
		e[edge[1]] = append(e[edge[1]], edge[0])
	}

	// dfs
	visited := make(map[int]struct{}, 0)
	lastNum := 0
	path := make([]int, 0,n)
	var maxPath []int

	var dfs func(num int)
	dfs = func(num int) {

		visited[num] = struct{}{}
		path = append(path, num)
		nextCnt := 0
		for _, val := range e[num] {
			if _, ok := visited[val]; ok {
				continue
			}
			nextCnt++
			dfs(val)
		}

		if nextCnt == 0 && len(path) > len(maxPath) {
			lastNum = num
			tmp := make([]int, len(path))
			copy(tmp, path)
			maxPath = tmp
		}

		path = path[:len(path)-1]
	}

	// first
	dfs(0)
	maxPath1 := maxPath
	maxPath = nil
	// reset
	for k := range visited {
		delete(visited, k)
	}

	// second
	dfs(lastNum)
	maxPath2 := maxPath

	if len(maxPath1) > len(maxPath2) {
		maxPath = maxPath1
	} else {
		maxPath = maxPath2
	}

	if len(maxPath)%2 == 0 {
		return []int{maxPath[len(maxPath)/2-1], maxPath[len(maxPath)/2]}
	}
	return []int{maxPath[len(maxPath)/2]}
}
