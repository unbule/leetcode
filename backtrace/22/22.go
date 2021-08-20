package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(2))
}

func generateParenthesis(n int) []string {
	res := []string{}

	var backtrace func(left, right int, path string)
	backtrace = func(left, right int, path string) {
		if 2*n == len(path) {
			res = append(res, path)
		}
		if left > 0 {
			backtrace(left-1, right, path+"(")
		}
		if left < right {
			backtrace(left, right-1, path+")")
		}
	}
	backtrace(n, n, "")
	return res
}
