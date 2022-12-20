package main

import (
	"log"
)

type Line struct {
	Bp Point
	Dp Point
}
type Point struct {
	X int
	Y int
}

func isSelfCrossing(distance []int) bool {
    for i := 3; i < len(distance); i++ {
        // 第 1 类路径交叉的情况
        if distance[i] >= distance[i-2] && distance[i-1] <= distance[i-3] {
            return true
        }

        // 第 2 类路径交叉的情况
        if i == 4 && distance[3] == distance[1] &&
            distance[4] >= distance[2]-distance[0] {
            return true
        }

        // 第 3 类路径交叉的情况
        if i >= 5 && distance[i-3]-distance[i-5] <= distance[i-1] &&
            distance[i-1] <= distance[i-3] &&
            distance[i] >= distance[i-2]-distance[i-4] &&
            distance[i-2] > distance[i-4] {
            return true
        }
    }
    return false
}

func main() {
	log.Println(isSelfCrossing([]int{1,1,2,1,1}))       //true
	log.Println(isSelfCrossing([]int{2, 1, 1, 2}))       //true
	log.Println(isSelfCrossing([]int{1, 1, 2, 2, 1, 1})) //true
	log.Println(isSelfCrossing([]int{3, 3, 3, 2, 1,1}))    //false
}
