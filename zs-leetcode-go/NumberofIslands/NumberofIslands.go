package NumberofIslands

func search(grid [][]byte, i, j int, siteMap [][]int){
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	siteMap[i][j] = 1
	for _, dir := range dirs {
		nextX, nextY := i+dir[0], j+dir[1]
		if nextX >= 0 && nextX < len(grid) && nextY >= 0 && nextY < len(grid[0]) && grid[nextX][nextY] == '1' &&
			siteMap[nextX][nextY] == 0{
			search(grid, nextX, nextY, siteMap)
		}
	}
}

func numIslands(grid [][]byte) int {
	cnt := 0
	siteMap := make([][]int, len(grid))
	for index := 0; index < len(grid); index++ {
		siteMap[index] = make([]int, len(grid[0]))
	}
	if len(grid) == 0 {
		return cnt
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' && siteMap[i][j] == 0 {
				cnt += 1
				search(grid, i, j, siteMap)
			}
		}
	}
	return cnt
}