package main

import "fmt"

func main() {
	graph := [][]int{{1, 2}, {0}, {0}}
	shortestPathLength(graph)
}

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	fmt.Println(1 << n)
	type tuple struct{ u, mask, dist int }
	q := []tuple{}
	seen := make([][]bool, n)
	for i := range seen {
		seen[i] = make([]bool, 1<<n)
		seen[i][1<<i] = true
		q = append(q, tuple{i, 1 << i, 0})
	}
	fmt.Println("first: ", q)
	fmt.Println("first: ", seen)
	for {
		t := q[0]
		q = q[1:]
		if t.mask == 1<<n-1 {
			return t.dist
		}
		// 搜索相邻的节点
		for _, v := range graph[t.u] {
			maskV := t.mask | 1<<v
			if !seen[v][maskV] {
				q = append(q, tuple{v, maskV, t.dist + 1})
				fmt.Println(q)
				seen[v][maskV] = true
			}
		}
	}
}
