package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 1 {
				maxArea = max(maxArea, dfs(grid, row, col))
			}
		}
	}

	return maxArea
}

func dfs(grid [][]int, row, col int) int {
	if grid[row][col] == 0 {
		return 0
	}
	area := 1
	// 防止重复搜索
	grid[row][col] = 0

	// up
	if row-1 >= 0 && grid[row-1][col] == 1 {
		area += dfs(grid, row-1, col)
	}
	// left
	if col-1 >= 0 && grid[row][col-1] == 1 {
		area += dfs(grid, row, col-1)
	}
	// right
	if col+1 < len(grid[row]) && grid[row][col+1] == 1 {
		area += dfs(grid, row, col+1)
	}
	// down
	if row+1 < len(grid) && grid[row+1][col] == 1 {
		area += dfs(grid, row+1, col)
	}
	return area
}

func main() {
	m:=[][]int{
		{1,1,0,0},
		{1,0,0,0},
		{1,1,0,0},
		{0,1,1,0},
	}
	fmt.Printf("%v\n",maxAreaOfIsland(m))
}
