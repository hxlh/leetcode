package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{0}, 1))
}

func canPlaceFlowers(flowerbed []int, n int) bool {

	for i := 0; i < len(flowerbed); i++ {

		if (i-1 < 0 && i+1 < len(flowerbed)) && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			n--
		} else if (i+1 >= len(flowerbed) && i-1 >= 0) && flowerbed[i] == 0 && flowerbed[i-1] == 0 {
			flowerbed[i] = 1
			n--
		} else if (i+1 <len(flowerbed) && i-1 >= 0) && flowerbed[i] == 0 && flowerbed[i+1] == 0 && flowerbed[i-1] == 0 {
			flowerbed[i] = 1
			n--
		}else if len(flowerbed)==1 && flowerbed[0]==0{
			flowerbed[i] = 1
			n--
		}

	}
	if n <= 0 {
		return true
	}

	return false
}
