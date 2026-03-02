package leetcode

import "container/list"

func OrangesRotting(grid [][]int) int {
	freshOranges := 0 // 用于记录新鲜橘子数
	m, n := len(grid), len(grid[0])
	queue := list.New()
	// 遍历统计新鲜橘子和腐烂橘子
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch grid[i][j] {
			case 1:
				freshOranges++ //好橘子加1
			case 2:
				queue.PushBack([2]int{i, j}) // 拦橘子入队
			}
		}
	}
	if freshOranges == 0 {
		return 0 // 没有新鲜橘子，则直接退出
	}
	if queue.Len() == 0 {
		return -1 // 没有拦橘子则返回
	}
	//
	dis := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	minutes := 0
	// 广度优先遍历直到为空
	for queue.Len() > 0 {
		levelSize := queue.Len()
		// 遍历整个层
		for i := 0; i < levelSize; i++ {
			pos := queue.Front().Value.([2]int)
			x, y := pos[0], pos[1]
			//出队
			queue.Remove(queue.Front())
			// 判断四个方向有什么东西
			for _, d := range dis {
				newX, newY := x+d[0], y+d[1]
				if newX < m && newY < n && newX >= 0 && newY >= 0 && grid[newX][newY] == 1 {
					grid[newX][newY] = 2
					// 新鲜橘子--
					freshOranges--
					// 拦橘子入队
					queue.PushBack([2]int{newX, newY})
				}
			}
		}
		if queue.Len() > 0 {
			minutes++
		}
	}
	if freshOranges != 0 {
		return -1
	}
	return minutes
}
