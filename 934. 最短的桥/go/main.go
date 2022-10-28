package main

import "fmt"

// 思路：深度优先遍历找第一个岛，找岛过程中遇到的海洋0加入队列,
// 然后将其不断向外延伸一圈，直到到达了另一座岛，延伸的圈数即为最短距离。
func shortestBridge(grid [][]int) int {
	q := make([][]int, 0)
	found := false
	for i := 0; i < len(grid); i++ {
		if found {
			break
		}
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				q = findFirstBridge(q, grid, i, j)
				found = true
				break
			}
		}
	}
	return bfs(q, grid)
}

// up right down left
var dirs []int=[]int{-1,0,1,0,-1}

func bfs(q [][]int, grid [][]int) int {
	level := 0
	for len(q) > 0 {
		level++
		n := len(q)
		for n > 0 {
			n--
			front := q[0]
			q = q[1:]
			for i := 0; i < 4; i++ {
				// 判断四个方向的下一个块
				x,y:=front[0]+dirs[i],front[1]+dirs[i+1]
				if x>=0 && x<len(grid) && y>=0 && y<len(grid[x]){
					if grid[x][y]==1{
						return level
					}
					if grid[x][y]==2{
						continue
					}
					q = append(q, []int{x,y})
					grid[x][y]=2
				}
			}	
		}
	}

	return 0
}

func findFirstBridge(q [][]int, grid [][]int, i, j int) [][]int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 2 {
		return q
	}
	if grid[i][j] == 0 {
		q = append(q, []int{i, j})
		return q
	}
	// 搜过的地方设为2，防止重复
	grid[i][j] = 2
	// up left right down
	q = findFirstBridge(q, grid, i-1, j)
	q = findFirstBridge(q, grid, i, j-1)
	q = findFirstBridge(q, grid, i, j+1)
	q = findFirstBridge(q, grid, i+1, j)

	return q
}

func main() {
	fmt.Printf("%v\n", shortestBridge([][]int{
		{0, 1},
		{1, 0},
	}))
}
