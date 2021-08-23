package main

func main() {

}

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	ans := 0
	var dfs func([][]byte, int, int)
	dfs = func(grid [][]byte, r, c int) {
		if r >= 0 && r < m && c >= 0 && c < n && grid[r][c] == '1' {
			grid[r][c] = '0'
			dfs(grid, r+1, c)
			dfs(grid, r-1, c)
			dfs(grid, r, c+1)
			dfs(grid, r, c-1)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(grid, i, j)
			}
		}
	}
	return ans
}
