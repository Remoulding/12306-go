package service

import "fmt"

// SelectAdjacentSeats 查找相邻座位
// 参数：numSeats 需要选择的座位数，seatLayout 座位布局矩阵（0表示可用）
// 返回值：选择的座位位置数组（行列从1开始计数），找不到时返回nil
func SelectAdjacentSeats(numSeats int, seatLayout [][]int) [][]int {
	rows := len(seatLayout)
	if rows == 0 {
		return nil
	}
	cols := len(seatLayout[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if seatLayout[i][j] == 0 {
				// 找到连续空位
				consecutive := 0
				startCol := j
				for k := j; k < cols; k++ {
					if seatLayout[i][k] == 0 {
						consecutive++
						if consecutive == numSeats {
							return convertToActualSeat(i, startCol, numSeats, true)
						}
					} else {
						consecutive = 0
						startCol = k + 1
					}
				}
			}
		}
	}
	return nil
}

// SelectNonAdjacentSeats 查找非相邻座位
func SelectNonAdjacentSeats(numSeats int, seatLayout [][]int) [][]int {
	var selected [][]int
	rows := len(seatLayout)
	if rows == 0 {
		return nil
	}
	cols := len(seatLayout[0])

outer:
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if seatLayout[i][j] == 0 {
				selected = append(selected, []int{i + 1, j + 1})
				if len(selected) == numSeats {
					break outer
				}
			}
		}
	}

	if len(selected) < numSeats {
		return nil
	}
	return selected
}

// convertToActualSeat 转换座位坐标（从0-based到1-based）
func convertToActualSeat(row, startCol, numSeats int, horizontal bool) [][]int {
	result := make([][]int, numSeats)
	for i := 0; i < numSeats; i++ {
		if horizontal {
			result[i] = []int{row + 1, startCol + i + 1}
		} else {
			// 纵向排列（当前实现仅处理横向）
			result[i] = []int{row + i + 1, startCol + 1}
		}
	}
	return result
}

// 示例用法
func ExampleUsage() {
	// 测试相邻座位选择
	seatLayout := [][]int{
		{1, 1, 1, 1},
		{1, 1, 1, 0},
		{1, 1, 1, 0},
		{0, 0, 0, 0},
	}

	if selected := SelectAdjacentSeats(2, seatLayout); selected != nil {
		fmt.Println("找到相邻座位：")
		for _, s := range selected {
			fmt.Printf("第%d排 第%d列\n", s[0], s[1])
		}
	} else {
		fmt.Println("没有足够相邻座位")
	}

	// 测试非相邻座位选择
	seatLayout2 := [][]int{
		{1, 0, 1, 1},
		{1, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 0, 0, 0},
	}

	if selected := SelectNonAdjacentSeats(3, seatLayout2); selected != nil {
		fmt.Println("\n找到非相邻座位：")
		for _, s := range selected {
			fmt.Printf("第%d排 第%d列\n", s[0], s[1])
		}
	} else {
		fmt.Println("没有足够空位")
	}
}

/*
示例输出：
找到相邻座位：
第4排 第1列
第4排 第2列

找到非相邻座位：
第1排 第2列
第2排 第3列
第2排 第4列
*/
