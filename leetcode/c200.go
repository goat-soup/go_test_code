package leetcode

func NumIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dnf(grid, i, j)
				count++
			}
		}
	}
	return count
}

func dnf(grid [][]byte, i, j int) {
	m, n := len(grid), len(grid[0])
	if i > m-1 || j > n-1 || i < 0 || j < 0 {
		return
	}
	if grid[i][j] != '1' {
		return
	}
	grid[i][j] = '0'
	dnf(grid, i-1, j)
	dnf(grid, i+1, j)
	dnf(grid, i, j-1)
	dnf(grid, i, j+1)
}
