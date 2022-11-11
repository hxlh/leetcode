package main

func main() {

}

// 思路：遍历将空格加入到一个队列中方便递归回溯，并将遇到的数字加入已出现的队列
func solveSudoku(board [][]byte) {
	var row, column [9][9]bool
	var block [3][3][9]bool
	space := make([][2]int, 0)
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				space = append(space, [2]int{i, j})
				continue
			}
			num := board[i][j] - '1'
			row[i][num] = true
			column[j][num] = true
			block[i/3][j/3][num] = true
		}
	}
	var dfs func(pos int) bool
	dfs = func(pos int) bool {
		if pos == len(space) {
			return true
		}
		x, y := space[pos][0], space[pos][1]

		for i := byte(0); i < 9; i++ {
			if !row[x][i] && !column[y][i] && !block[x/3][y/3][i] {
				row[x][i] = true
				column[y][i] = true
				block[x/3][y/3][i] = true
				// 不需要回溯board，因为下一次会覆盖
				board[x][y] = i + '1'
				if dfs(pos + 1) {
					return true
				}
				// 回溯
				row[x][i] = false
				column[y][i] = false
				block[x/3][y/3][i] = false
			}
		}
		return false
	}

	dfs(0)
}
