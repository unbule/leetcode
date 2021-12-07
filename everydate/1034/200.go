package main

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	ans := 0
	var dfs func([][]byte, int, int)
	dfs = func(g [][]byte, x, y int) {
		if x >= 0 && x < m && y >= 0 && y < n && g[x][y] == 1 {
			g[x][y] = 0
			dfs(g, x+1, y)
			dfs(g, x-1, y)
			dfs(g, x, y+1)
			dfs(g, x, y-1)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ans++
				dfs(grid, i, j)
			}
		}
	}
	return ans
}
