package main

import "log"

func isValidSudoku(board [][]byte) bool {
	//--》
	for i := 0; i < 9; i++ {
		m := make(map[byte]byte)
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			_, ok := m[board[i][j]]
			if !ok {
				m[board[i][j]] = 1
			} else {
				return false
			}
		}
	}
	/*
		|
		|
	*/
	for i := 0; i < 9; i++ {
		m := make(map[byte]byte)
		for j := 0; j < 9; j++ {
			if board[j][i] == '.' {
				continue
			}
			_, ok := m[board[j][i]]
			if !ok {
				m[board[j][i]] = 1
			} else {
				return false
			}
		}
	}

	//田
	for i := 0; i < 9; i+=3 {
		nums:=make(map[byte]byte)
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if board[j][k+i] == '.' {
					continue
				}
				_,ok:=nums[board[j][k+i]]
				if !ok {
					nums[board[j][k+i]]=1
				}else{
					return false
				}
			}
		}
	}
	for i := 0; i < 9; i+=3 {
		nums:=make(map[byte]byte)
		for j := 3; j < 6; j++ {
			for k := 0; k < 3; k++ {
				if board[j][k+i] == '.' {
					continue
				}
				_,ok:=nums[board[j][k+i]]
				if !ok {
					nums[board[j][k+i]]=1
				}else{
					return false
				}
			}
		}
	}
	for i := 0; i < 9; i+=3 {
		nums:=make(map[byte]byte)
		for j := 6; j < 9; j++ {
			for k := 0; k < 3; k++ {
				if board[j][k+i] == '.' {
					continue
				}
				_,ok:=nums[board[j][k+i]]
				if !ok {
					nums[board[j][k+i]]=1
				}else{
					return false
				}
			}
		}
	}



	return true
}

func main() {
	board := [][]byte{
		{5, 3, '.', '.', 7, '.', '.', '.', '.'},
		{6, '.', '.', 1, 9, 5, '.', '.', '.'},
		{'.', 9, 8, '.', '.', '.', '.', 6, '.'},
		{8, '.', '.', '.', 6, '.', '.', '.', 3},
		{4, '.', '.', 8, '.', 3, '.', '.', 1},
		{7, '.', '.', '.', 2, '.', '.', '.', 6},
		{'.', 6, '.', '.', '.', '.', 2, 8, '.'},
		{'.', '.', '.', 4, 1, 9, '.', '.', 5},
		{'.', '.', '.', '.', 8, '.', '.', 7, 9},
	}
	log.Println(isValidSudoku(board))
}
