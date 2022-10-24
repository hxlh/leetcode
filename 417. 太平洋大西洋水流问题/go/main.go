package main

func pacificAtlantic(heights [][]int) [][]int {
	pacific_reach := make([][]bool, len(heights))
	atlantic_reach := make([][]bool, len(heights))
	for i := 0; i < len(pacific_reach); i++ {
		pacific_reach[i] = make([]bool, len(heights[0]))
		atlantic_reach[i] = make([]bool, len(heights[0]))
	}
	// 上边界
	for i := 0; i < len(heights[0]); i++ {
		dfs(heights, pacific_reach, 0, i)
	}
	// 左边界
	for i := 1; i < len(heights); i++ {
		dfs(heights, pacific_reach, i, 0)
	}

	// 右边界
	for i := 0; i < len(heights); i++ {
		dfs(heights, atlantic_reach, i, len(heights[0])-1)
	}

	// 下边界
	for i := 0; i < len(heights[0])-1; i++ {
		dfs(heights, atlantic_reach, len(heights)-1, i)
	}

	// fmt.Println("pacific_reach")
	// for i := 0; i < len(pacific_reach); i++ {
	// 	fmt.Printf("%v ", pacific_reach[i])
	// 	fmt.Println("")
	// }

	// fmt.Println("atlantic_reach")
	// for i := 0; i < len(atlantic_reach); i++ {
	// 	fmt.Printf("%v ", atlantic_reach[i])
	// 	fmt.Println("")
	// }

	res := make([][]int, 0)
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[0]); j++ {
			if pacific_reach[i][j] && atlantic_reach[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	// fmt.Println("res")
	// for i := 0; i < len(res); i++ {
	// 	fmt.Printf("%v ", res[i])
	// 	fmt.Println("")
	// }

	return res
}

func dfs(heights [][]int, ocean [][]bool, x, y int) {
	if ocean[x][y] {
		return
	}
	ocean[x][y] = true
	// 遍历四个方向 up left right down
	if x-1 >= 0 && heights[x-1][y] >= heights[x][y] {
		dfs(heights, ocean, x-1, y)
	}
	if y-1 >= 0 && heights[x][y-1] >= heights[x][y] {
		dfs(heights, ocean, x, y-1)
	}
	if y+1 < len(ocean[0]) && heights[x][y+1] >= heights[x][y] {
		dfs(heights, ocean, x, y+1)
	}
	if x+1 < len(ocean) && heights[x+1][y] >= heights[x][y] {
		dfs(heights, ocean, x+1, y)
	}
}

func main() {
	pacificAtlantic([][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	})
}
