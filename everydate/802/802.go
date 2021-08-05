package main

import "fmt"

func main() {
	graph := [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}
	fmt.Println(eventualSafeNodes(graph))
}

func eventualSafeNodes(graph [][]int) (ans []int) {
	n := len(graph)
	color := make([]int, n)
	var safe func(int) bool
	safe = func(i int) bool {
		if color[i] > 0 {
			return color[i] == 2
		}
		color[i] = 1
		for _, g := range graph[i] {
			if !safe(g) {
				return false
			}
		}
		color[i] = 2
		return true
	}

	for i := 0; i < n; i++ {
		if safe(i) {
			ans = append(ans, i)
		}
	}
	return
}
