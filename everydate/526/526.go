package main

import "fmt"

func main() {
	n := 2
	fmt.Println(countArrangement(n))
}

func countArrangement(n int) int {
	ans := 0
	used := make([]bool, n+1)
	match := make([][]int, n+1)
	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			if i%j == 0 || j%i == 0 {
				match[i] = append(match[i], j)
			}
		}
	}

	var backtrace func(int)

	backtrace = func(i int) {
		if i > n {
			ans++
			return
		}
		for _, x := range match[i] {
			if !used[x] {
				used[x] = true
				backtrace(i + 1)
				used[x] = false
			}
		}
	}
	backtrace(1)
	return ans
}
