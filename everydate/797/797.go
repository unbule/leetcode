package main

import "fmt"

func main() {
	graph := [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}
	fmt.Println(allPathsSourceTarget(graph))
}

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph) - 1
	ans := make([][]int, 0)

	var backtrace func(int, int, []int)
	backtrace = func(location, index int, path []int) {
		if n == location {
			res := make([]int, len(path))
			copy(res, path)
			ans = append(ans, res)
			return
		}
		for _, i := range graph[index] {
			path = append(path, i)
			backtrace(i, i, path)
			path = path[:len(path)-1]
		}
	}
	backtrace(0, 0, []int{0})
	return ans
}
