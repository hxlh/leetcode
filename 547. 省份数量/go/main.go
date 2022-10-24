package main

import "fmt"


func findCircleNum(isConnected [][]int) int {
	if len(isConnected) == 0 || len(isConnected[0]) == 0 {
		return 0
	}
	maxCnt := 0
	visited := make([]bool, len(isConnected))
	for row := 0; row < len(isConnected); row++ {
		if !visited[row] {
			dfs(isConnected, row, 0, visited)
			maxCnt++
		}
	}

	return maxCnt
}

func dfs(isConnected [][]int, row, col int, visited []bool){
	visited[row] = true
	for i := col; i < len(isConnected[row]); i++ {
		if isConnected[row][i] == 1 && !visited[i] {
			dfs(isConnected, i, 0, visited)
		}
	}
}

func main() {
	fmt.Println(findCircleNum([][]int{
		{1,0,0,1},
		{0,1,1,0},
		{0,1,1,1},
		{1,0,1,1},
	}))
}
