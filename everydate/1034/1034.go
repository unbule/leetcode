package main

func main() {

}

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	m := len(grid)
	n := len(grid[0])
	ocolor := grid[row][col]
	type point struct{ x, y int }
	borders := []point{}
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	var dfs func(int, int)
	dfs = func(x, y int) {
		vis[x][y] = true
		isborder := false
		for _, dir := range dirs {
			nx := x + dir.x
			ny := y + dir.y
			// 判断是否为边界
			if !(nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == ocolor) {
				isborder = true
				// 判断是否已经搜索过
			} else if !vis[nx][ny] {
				vis[nx][ny] = true
				dfs(nx, ny)
			}
		}
		if isborder {
			borders = append(borders, point{x, y})
		}
	}
	dfs(row, col)
	for _, p := range borders {
		grid[p.x][p.y] = color
	}
	return grid
}
