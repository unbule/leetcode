package main

import "fmt"

func main() {
	fmt.Println(tribonacci(4))
}

func tribonacci(n int) int {
	a, b, c := 0, 1, 1
	end := 1
	if n == 0 {
		return a
	}
	if n == 1 {
		return b
	}
	for j := 2; j <= n-1; j++ {
		end = a + b + c
		a = b
		b = c
		c = end
		fmt.Println(end)
	}
	return end
}
