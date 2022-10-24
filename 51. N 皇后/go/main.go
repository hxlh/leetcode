package main

import "fmt"

func solveNQueens(n int) [][]string {
	board := make([][]byte, n)
	str := ""
	for i := 0; i < n; i++ {
		str += "."
	}
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		board[i] = []byte(str)
	}
	cols := make([]bool, n)
	// 2*n-1条斜边
	lslant := make([]bool, 2*n-1)
	rslant := make([]bool, 2*n-1)

	ans := make([][]string, 0)

	if n == 0 {
		return ans
	}

	ans = backtrace(ans, board, cols, lslant, rslant, 0)

	return ans
}

// row 当前行
func backtrace(ans [][]string, board [][]byte, cols []bool, lslant []bool, rslant []bool, row int) [][]string {
	if row >= len(board) {
		strArr := make([]string, 0)
		for i := 0; i < len(board); i++ {
			strArr = append(strArr, string(board[i]))
		}
		ans = append(ans, strArr)
		return ans
	}

	// 按列
	for i := 0; i < len(board); i++ {
		if cols[i] || lslant[((len(board)*2-2))/2+i-row] || rslant[(len(board)*2-2)-row-i] {
			continue
		}
		board[row][i] = 'Q'
		cols[i] = true
		lslant[((len(board)*2-2))/2+i-row] = true
		rslant[(len(board)*2-2)-row-i] = true

		ans = backtrace(ans, board, cols, lslant, rslant, row+1)

		board[row][i] = '.'
		cols[i] = false
		lslant[((len(board)*2-2))/2+i-row] = false
		rslant[(len(board)*2-2)-row-i] = false
	}

	return ans
}

func main() {
	ans := solveNQueens(4)
	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i])
		fmt.Println("----------------------")
	}
}
