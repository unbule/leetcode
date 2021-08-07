package main

import "fmt"

func main() {
	n := 5
	fmt.Println(countBits(n))
}

func countBits(n int) []int {
	heighbit := 0
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			heighbit = i
		}
		bits[i] = bits[i-heighbit] + 1
	}
	return bits
}
