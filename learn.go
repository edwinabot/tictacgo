package main

import (
	"fmt"
)

var rays [][]int = [][]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	{0, 4, 8},
	{6, 4, 2},
}

func isTicTacToe(grid []int) bool {
	for i := 0; i < len(rays); i++ {
		var x, y, z = rays[i][0], rays[i][1], rays[i][2]
		if (grid[x] == grid[y] && grid[y] == grid[z]) && (grid[x] != 0) {
			return true
		}
	}
	return false
}

func main() {
	// Create a tic-tac-toe board.
	var grid []int = []int{-1, 0, 0, 1, -1, 1, 0, 0, -1}

	var isIt bool = isTicTacToe(grid)

	fmt.Printf("Is tictactgo? %t\n", isIt)
}
