package main

import "fmt"

func main() {
	solve([][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	})
}

const debug=true

func Debug(s string,v ...interface{})  {
	if debug{
		fmt.Printf(s,v...)
	}
}

func solve(board [][]byte) {
	m, n := len(board), len(board[0])

	// 从四边开始扫
	edge_queue := make([][]int, 0)
	for i := 0; i < n; i++ {
		if board[0][i] == 'O' {
			edge_queue = append(edge_queue, []int{0, i})
		}
	}
	for i := 0; i < n; i++ {
		if board[m-1][i] == 'O' {
			edge_queue = append(edge_queue, []int{m - 1, i})
		}
	}
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			edge_queue = append(edge_queue, []int{i, 0})
		}
	}
	for i := 0; i < m; i++ {
		if board[i][n-1] == 'O' {
			edge_queue = append(edge_queue, []int{i, n - 1})
		}
	}

	// 从edge_queue bfs
	for len(edge_queue) > 0 {
		cnt := len(edge_queue)
		for cnt > 0 {
			cnt--

			pos := edge_queue[0]
			edge_queue = edge_queue[1:]

			x, y := pos[0], pos[1]
			board[x][y] = 'E'

			// up down left right
			if x-1 >= 0 && board[x-1][y] == 'O' {
				edge_queue = append(edge_queue, []int{x - 1, y})
			}
			if x+1 < len(board) && board[x+1][y] == 'O' {
				edge_queue = append(edge_queue, []int{x + 1, y})
			}
			if y-1 >= 0 && board[x][y-1] == 'O' {
				edge_queue = append(edge_queue, []int{x, y - 1})
			}
			if y+1 < len(board[x]) && board[x][y+1] == 'O' {
				edge_queue = append(edge_queue, []int{x, y + 1})
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j]=='O'{
				board[i][j]='X'
			}else if board[i][j]=='E'{
				board[i][j]='O'
			}
		}
	}
	
	if debug{
		for i := 0; i < len(board); i++ {
			fmt.Println(board[i])
		}
	}
}
