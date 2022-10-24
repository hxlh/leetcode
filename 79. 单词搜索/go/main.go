package main

import "fmt"

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	row := len(board)
	col := len(board[0])
	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, col)
	}
	find:=false
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			backtrace(i, j, board, word, visited, 0, &find)
		}
	}
	return find
}

func backtrace(i int, j int, board [][]byte, word string, visited [][]bool, pos int, find *bool) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return
	}
	if visited[i][j] || board[i][j] != word[pos] || *find {
		return
	}
	if pos == len(word)-1 {
		// found
		*find = true
		return
	}

	visited[i][j] = true
	// 递归搜索up left right down
	backtrace(i-1, j, board, word, visited, pos+1, find)
	backtrace(i, j-1, board, word, visited, pos+1, find)
	backtrace(i, j+1, board, word, visited, pos+1, find)
	backtrace(i+1, j, board, word, visited, pos+1, find)
	
	visited[i][j] = false
}

func main() {
	fmt.Println(exist([][]byte{
		{'A','B','C','E'},
		{'S','F','C','S'},
		{'A','D','E','E'},
	},"ABCCED"))
}
