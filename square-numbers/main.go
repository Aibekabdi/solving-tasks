package main

import "fmt"

func main() {
	arr := [3][3]int{{-1, -1, -1}, {-1, -1, -1}, {-1, -1, -1}}
	if backtrack(&arr, 0, 0) {
		for i := 0; i < 3; i++ {
			fmt.Println(arr[i])
		}
	} else {
		fmt.Println("impossible")
	}
}
func backtrack(grid *[3][3]int, row, col int) bool {
	if row == 2 && col == 3 {
		return check(grid)
	}
	if col == 3 {
		row += 1
		col = 0
	}
	if grid[row][col] > -1 {
		return backtrack(grid, row, col+1)
	}
	for num := 3; num <= 11; num++ {
		if isSafe(grid, num) {
			grid[row][col] = num
			if backtrack(grid, row, col+1) {
				return true
			}
		}
		grid[row][col] = -1
	}
	return false
}

func isSafe(grid *[3][3]int, num int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i][j] == num {
				return false
			}
		}
	}
	return true
}

func check(grid *[3][3]int) bool {
	for i := 0; i < len(grid); i++ {
		rowProduct := grid[i][0] * grid[i][1] * grid[i][2]
		colProduct := grid[0][i] * grid[1][i] * grid[2][i]
		if rowProduct != colProduct {
			return false
		}
	}
	return true
}
