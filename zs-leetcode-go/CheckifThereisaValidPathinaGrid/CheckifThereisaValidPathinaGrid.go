package CheckifThereisaValidPathinaGrid

// 输入条件保证了
// 每种情况下，可走的方向
var dirs = map[int][][]int{
	1: [][]int{{0, -1}, {0, 1}},
	2: [][]int{{-1, 0}, {1, 0}},
	3: [][]int{{0, -1}, {1, 0}},
	4: [][]int{{0, 1}, {1, 0}},
	5: [][]int{{0, -1}, {-1, 0}},
	6: [][]int{{-1, 0}, {0, 1}},
}

type pair struct {
	X, Y int
}

func hasValidPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	if grid[0][0] == 5 {
		return false
	}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	queue := []pair{{0, 0}}
	for len(queue) > 0 {
		frontPair := queue[0]
		if frontPair.X == m-1 && frontPair.Y == n-1 {
			return true
		}
		queue = queue[1:]
		frontType := grid[frontPair.X][frontPair.Y]
		for _, dir := range dirs[frontType] {
			nx, ny := frontPair.X+dir[0], frontPair.Y+dir[1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n || visited[nx][ny] {
				continue
			}
			nextDirs := dirs[grid[nx][ny]]
			if (dir[0] == -nextDirs[0][0] && dir[1] == -nextDirs[0][1]) || (dir[0] == -nextDirs[1][0] && dir[1] == -nextDirs[1][1]) {
				queue = append(queue, pair{nx, ny})
				visited[nx][ny] = true
			}
		}
	}
	return false
}
