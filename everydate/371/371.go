package main

import "fmt"

func main() {
	a := 11
	b := 4
	fmt.Println(getSum(a, b))
}

func getSum(a int, b int) int {
	for b != 0 {
		carry := uint(a&b) << 1
		a ^= b
		b = int(carry)
	}
	return a
}
