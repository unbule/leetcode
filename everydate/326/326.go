package main

import "fmt"

func main() {
	fmt.Println(isPowerOfThree(27))
}

func isPowerOfThree(n int) bool {
	for n > 0 && n%3 == 0 {
		n = n / 3
	}
	return n == 1
}
