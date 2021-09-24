package main

import "fmt"

func main() {
	fmt.Println(combine1(3, 2))
}

func combine(n int, k int) [][]int {
	seen := make([]bool, n+1)
	ans := make([][]int, 0)
	res := make([]int, 0)
	var backtrace func([]bool, int, []int)
	backtrace = func(seen []bool, index int, res []int) {
		if len(res) == k {
			temp := make([]int, len(res))
			copy(temp, res)
			ans = append(ans, temp)
			return
		}
		for i := 1; i <= n; i++ {
			if !seen[i] {
				seen[i] = true
				res = append(res, i)
				backtrace(seen, i+1, res)
				res = res[:len(res)-1]
				seen[i] = false
			}
		}
	}
	backtrace(seen, 1, res)
	return ans
}

func combine1(n int, k int) [][]int {
	res := [][]int{}
	var backtrace func(start int, path []int)
	backtrace = func(start int, path []int) {

		ans := make([]int, len(path))
		copy(ans, path)
		res = append(res, ans)

		for i := start; i <= n; i++ {
			path = append(path, i)
			backtrace(i+1, path)
			path = path[:len(path)-1]
		}
	}
	backtrace(1, []int{})
	return res
}
