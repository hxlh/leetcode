package main

import "log"

func fib(n int) int {
	f := make([]int, n+1)
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[len(f)-1]
}

func main() {
	log.Println(fib(4))
}
