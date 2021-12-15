package main

import "fmt"

func main() {
	richer := [][]int{{1, 0}, {2, 1}, {3, 1}, {3, 7}, {4, 3}, {5, 3}, {6, 3}}
	quiet := []int{3, 2, 5, 4, 6, 1, 7, 0}
	loudAndRich(richer, quiet)
}

func loudAndRich(richer [][]int, quiet []int) []int {
	n := len(quiet)
	cnt := make([][]int, n)
	for _, v := range richer {
		cnt[v[1]] = append(cnt[v[1]], v[0])
	}
	fmt.Println(cnt)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	var dfs func(int)
	dfs = func(m int) {
		if ans[m] != -1 {
			return
		}
		ans[m] = m
		for _, v := range cnt[m] {
			fmt.Println(m, v)
			dfs(v)
			if quiet[ans[v]] < quiet[ans[m]] {
				ans[m] = ans[v]
			}
		}
	}
	for i := 0; i < n; i++ {
		dfs(i)
	}
	return ans
}
