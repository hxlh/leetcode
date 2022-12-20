package main

import (
	"log"
)

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	rotate(matrix)
	log.Println(matrix)
}

func rotate(matrix [][]int) {
	//上下翻转
	for i := 0; i < len(matrix); i++ {
		pre := 0
		next := len(matrix[i]) - 1
		for {
			tmp := matrix[next][i]
			matrix[next][i] = matrix[pre][i]
			matrix[pre][i] = tmp
			pre++
			next--
			if pre == next || pre > next {
				break
			}
		}
	}
	//对角翻转
	for i := 0; i < len(matrix); i++ {
		for j := i+1; j < len(matrix[i]); j++ {
			tmp:=matrix[i][j]
			matrix[i][j]=matrix[j][i]
			matrix[j][i]=tmp
		}
	}

}
