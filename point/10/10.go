package main

import "fmt"

func main() {
	fmt.Println(fib(9))
}

func fib(n int) int {
	const mod int = 1e9 + 7
	if n < 2 {
		return n
	}
	p, q, r := 0, 1, 1
	for i := 2; i < n; i++ {
		p = q
		q = r
		r = (p + q) % mod
	}
	return r
}
